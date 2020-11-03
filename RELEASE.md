# Release criteria
This is a set of criteria to determine if this project is ready to release.

## Process
1. apply hpcaas-job-scheduler to a namespace
  * see [hpcaas-dev-deploy](https://github.com/hpe-hcss/hpcaas-dev-deploy)
1. verify that the Version grpc command works
  * This can be done using `kubectl port-forward`
  * example: `kubectl port-forward deployment/hpcaas-job-scheduler 8080`
  * [evans](https://github.com/ktr0731/evans) can be ran in interactive mode using `evans --host 127.0.0.1 --port 8080 -r`
1. verify that metrics are available and working as expected
