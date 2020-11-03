//(C) Copyright 2020 Hewlett Packard Enterprise Development LP

package main

import (
	"context"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"time"

	"fmt"

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

	// create a new HTTP client
	client := &http.Client {
		Timeout: 2 * time.Minute,
	}

	// send an GET request from `client`
	res, err := client.Get( "https://localhost:9000" )
	if err != nil {
		log.Error(ctx, "Error when calling SayHello: %s", err)
	} else {
		fmt.Println( "Success: status-code", res.StatusCode );
	}

	// read all response body
	data, _ := ioutil.ReadAll( res.Body )

	// close response body
	res.Body.Close()
	log.Info(ctx, "Response from server: %s", data)

	// print `data` as a string
	fmt.Printf( "%s\n", data )

}
