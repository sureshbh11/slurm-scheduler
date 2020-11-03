// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package view

import (
	"github.com/hpe-hcss/hpcaas-job-scheduler/internal/app/job"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	hpeErrors "github.com/hpe-hcss/errors/pkg/errors"
)

const (
	errGenericBadRequest = "ErrBadRequest"
)

// HTTPView is the HTTP API viewport into this service
type HTTPView struct {
	job        job.JobAPI
}

// NewHTTP is the factory method for this object
func NewHTTP(j job.JobAPI) HTTPView {
	log.Infof("Creating a new HTTP Interface for JobAPI")
	return HTTPView{
		job: j,
	}
}

// ListJobs lists the job
func (h *HTTPView) ListJobs(c *gin.Context) {
	ctx := c.Request.Context()

	jobs, err := h.job.List(ctx)
	if err != nil {
		log.Errorf("Could not process List Job request: %+v", err)
		hpeErrors.SetResponseIfError(c, err)
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *HTTPView) GetJob(c *gin.Context, id int64) {
	ctx := c.Request.Context()

	job, err := h.job.Get(ctx, id)
	if err != nil {
		log.Errorf("Could not process List Job request: %+v", err)
		hpeErrors.SetResponseIfError(c, err)
		return
	}

	c.JSON(http.StatusOK, job)
}