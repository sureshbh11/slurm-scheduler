//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/gin-gonic/gin"

	abRegistrationClient "github.com/hpe-hcss/iam-lib/pkg/ab-registration-client"
	idClient "github.com/hpe-hcss/iam-lib/pkg/identity-client"
	spacesclient "github.com/hpe-hcss/iam-lib/pkg/spaces-client"
	accessToken "github.com/hpe-hcss/iam-lib/pkg/token"
	"github.com/hpe-hcss/loglib/pkg/log"
	logmiddleware "github.com/hpe-hcss/loglib/pkg/middleware/gin"
	"github.com/hpe-hcss/scm-lib/pkg/scmclient"
	"github.com/hpe-hcss/tracing/pkg/trace"

	authzbroker "github.com/hpe-hcss/hpcaas-job-scheduler/internal/app/authz-broker"
	abregistrar "github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/iam/ab-registrar"
)

var (
	healthzEndpoint = "/healthz"
	skipPath        = []string{healthzEndpoint}
)

func main() {
	// create the scmClient using aws credentials and scmCreds
	serviceConfig, err := authzbroker.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("Could not load service options: %v", err.Error()))
	}
	scmCreds := aws.Config{
		Credentials: credentials.NewStaticCredentials(
			serviceConfig.SCMAWSAccessID,
			serviceConfig.SCMAWSSecretKey,
			""),
		Region: &serviceConfig.SCMAWSRegion,
	}
	scmConfig, err := scmclient.NewConfig(serviceConfig.SCMConfigPath)
	if err != nil {
		panic(fmt.Sprintf("unable to create scm client config: %v", err))
	}

	scmClient, err := scmclient.NewWithAwsCreds(scmConfig, scmCreds)
	if err != nil {
		panic(fmt.Sprintf("could not create an scm client instance: %v", err))
	}

	authBrokerConfig, err := authzbroker.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("Could not load service options: %v", err.Error()))
	}

	ctx := context.Background()

	log.Info(ctx, "Starting HPC Slurm Scheduler Service")

	httpClient := &http.Client{Timeout: time.Second * 30}
	httpClient.Transport = &trace.Transport{}

	spacesClient := spacesclient.New(serviceConfig.IamURL)

	abClient := abregistrar.NewRegistrar(serviceConfig.IamURL, abRegistrationClient.New(
		serviceConfig.IamURL,
		serviceConfig.IamURL,
		httpClient), httpClient)

	authBroker := authzbroker.Broker(
		authBrokerConfig, // authbroker config
		scmClient,        // this is currently triggered by loadConfig
		abClient,         // what is a abClient
		spacesClient,
		scmConfig.Client.ServiceSpace,
	)

	router := gin.New()
	configureRouter(router)

	healthz := router.Group(healthzEndpoint)
	healthz.GET("",
		func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

	// This must be called by OSB while tenant/customer on-boarding
	scheduler := router.Group("/scheduler/v1")

	configureAuthMiddleware(scheduler, serviceConfig)

	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient = &http.Client{Timeout: time.Second * 30, Transport: tr}

	scheduler.POST("/bootstrap", authBroker.BootStrapAB)
	scheduler.POST("/deregister", authBroker.DeRegisterAB)

	resources := scheduler.Group("/resources")
	resources.GET("/", authBroker.Resources)

	jobSchedulerClient := JobSchedulerClient{httpClient}

	job := scheduler.Group("/scheduler")
	job.GET("/listJobs", jobSchedulerClient.ListJobs)
	job.GET("/getJob?id=:id", jobSchedulerClient.GetJob)

	log.Infof(ctx, "Listening on %d (status)", serviceConfig.Port)

	err = router.Run(":" + strconv.Itoa(serviceConfig.Port))
	if err != nil {
		panic(fmt.Sprintf("could not start the HPCaas-ab microservice: %v", err))
	}

	log.Info(ctx, "HPCaas Authorization Broker microservice gracefully shutting down")
}

func configureRouter(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.HandleMethodNotAllowed = true
	config := logmiddleware.AuditConfig{
		SkipPath: skipPath,
	}
	router.Use(logmiddleware.AuditMiddleware(config))
}

func configureAuthMiddleware(v1 *gin.RouterGroup, config authzbroker.ServiceOptions) {
	if config.AuthType == "no_auth" {
		return
	}
	if config.IamURL == "" {
		err := "Error while fetching IAMURL. IAMURL variable not set"
		log.Error(context.Background(), err)
		panic(err)
	}
	iamClient := idClient.New(config.IamURL)
	authMW := accessToken.ExtractToken(iamClient)
	v1.Use(authMW)
}
