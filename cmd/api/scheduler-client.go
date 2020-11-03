package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type JobSchedulerClientAPI interface {
	ListJobs(*gin.Context) error
	GetJob(*gin.Context, string) error
}

type JobSchedulerClient struct {
	client *http.Client
}

type Job struct {
	id int
	name string
	status string
}

func processResponse(res *http.Response, err error) {

	// check for response error
	if err != nil {

		// get `url.Error` struct pointer from `err` interface
		urlErr := err.( *url.Error )

		// check if error occurred due to timeout
		if urlErr.Timeout() {
			fmt.Println( "Error occurred due to a timeout." );
		}

		// log error and exit
		log.Fatal( "Error:", err )
	} else {
		fmt.Println( "Success: status-code", res.StatusCode );
	}

	// read all response body
	data, _ := ioutil.ReadAll( res.Body )

	// close response body
	res.Body.Close()

	// print `data` as a string
	fmt.Printf( "%s\n", data )
}

func(s *JobSchedulerClient) ListJobs(c *gin.Context) {
	// send an GET request from `client`
	res, err := s.client.Get("https://localhost:9002/listJobs")
	processResponse(res, err)
}

func(s *JobSchedulerClient) GetJob(c *gin.Context) {
	// send an GET request from `client`
	id := c.Params.ByName("id")
	res, err := s.client.Get(fmt.Sprintf( "https://localhost:9002/job?id=%s", id))
	processResponse(res, err)

}
