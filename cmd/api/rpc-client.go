//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"google.golang.org/grpc"

	stub "github.com/hpe-hcss/hpcaas-job-scheduler/internal/pkg/domain"

	"github.com/hpe-hcss/loglib/pkg/log"
)

func main() {
	// create the scmClient using aws credentials and scmCreds
	ctx := context.Background()

	log.Info(ctx, "Starting HPC Slurm Scheduler Service")

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Error(ctx,"did not connect: %s", err)
	}
	defer conn.Close()

	c := stub.NewTxtMessageServiceClient(conn)

	response, err := c.SayHello(context.Background(), &stub.TxtMessage{Txt: "Hello From Client!"})
	if err != nil {
		log.Error(ctx, "Error when calling SayHello: %s", err)
	}
	log.Info(ctx, "Response from server: %s", response.Txt)

}
