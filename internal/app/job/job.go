// (c) Copyright 2020 Hewlett Packard Enterprise Development LP

package job

import (
	"context"
	"math"
)

const (
	// TenantTypeDocker defines the type of tenant type
	TenantTypeDocker = "docker"
	// Internal defines the type of RDA endpoint
	Internal = "internal"
)

// JobAPI is an interface which implements Job API
type JobAPI interface {
	List(ctx context.Context) ([]Job, error)
	Get(ctx context.Context, id int64) (Job, error)
}

// Job is a struct with Job Info
type Job struct {
	id    int64
	name string
	status string
}

// New is a initializer for struct Job
func New(id int64, name, status string) Job {
	return Job{
		id:    id,
		name: name,
		status: status,
	}
}

// List is a function which returns list of jobs from Slurm
func (s *Job) List(ctx context.Context) ([]Job, error) {
	jobs := []Job {
		Job{
			1, "my job", "r",
		},
		Job{
			2, "my job2", "q",
		},
	}
	return jobs, nil
}

// Get is a function which returns job from Slurm
func (s *Job) Get(ctx context.Context, id int64) (Job, error) {
	job := Job {
			1, "my job", "r",
		}

	return job, nil
}

// Convert the value from MB(int64) to GB(float64)
func convertMBToGB(value int64) float64 {
	result := math.Round((float64(value)/1024.0)*100) / 100
	return result
}
