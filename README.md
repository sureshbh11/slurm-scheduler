[![CircleCI](https://circleci.com/gh/hpe-hcss/hpcaas-.svg?style=svg&circle-token=57443e33f23ffee050a58c36eb159664e512ba63)](https://app.circleci.com/settings/project/github/hpe-hcss/hpcaas-broker)

# hpcaas-job-scheduler
Scheduler Service for Slurm in HPCaaS

## Setup
* set up your proto files repo using https://github.com/hpe-hcss/go-service-template-lib as a template
* update versions in Dockerfile and go.mod
* replace hpcaas- in Dockerfile with the name of your service
* replace hpcaas- in the files in etc/ with the name of your service
* set the entrypoint in the Dockerfile for your project
* update CODEOWNERS to refer to your team
* update the release criteria for your project
* update the go files as necessary for your project
* run `go mod tidy` to update your go modules files
* update the alerts.yaml file in helm with alerts that make sense for your service
* Optionally, you can add grafana dashboards to monitor your service. Put dashboard JSON file(s) in the `./helm/hpcaas-/dashboards/` directory. Update dashboards-configmap.yaml to include any dashboard files.
* Update the CircleCI badge
