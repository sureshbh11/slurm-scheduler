// (C) Copyright 2020 Hewlett Packard Enterprise Development LP

package constants

const (
	//ContainerPlatform name in authz resources
	ContainerPlatform = "Container Platform MLOps"

	//ContainerPlatformPath path in authz resources
	ContainerPlatformPath = "/"

	//RootAppliance name in authz resource
	RootAppliance = "Sites"

	//RootAppliancePath path in resource
	RootAppliancePath = "appliances"

	//RootProject name in authz resource
	RootProject = "Projects"

	//RootProjectPath path in resource
	RootProjectPath = "projects"

	//AllAppliances name in resource
	AllAppliances = "All Sites"

	//OpentracingServiceName is the name of the Jaeger service
	OpentracingServiceName = "mlops-api-svc"

	//ApplianceNotFound is custom Error Message if Appliance not found
	ApplianceNotFound = "ApplianceNotFound"
)
