package gcloud

import (
	"log"
	"os"
)

type Project string
type Zone string
type Region string
type ContainerCluster string

func CurrentProject() Project {
	return Project(getRequired("CLOUDSDK_CORE_PROJECT"))
}
func CurrentZone() Zone {
	return Zone(getRequired("CLOUDSDK_COMPUTE_ZONE"))
}
func CurrentRegion() Region {
	return Region(getRequired("CLOUDSDK_COMPUTE_REGION"))
}
func CurrentCluster() ContainerCluster {
	return ContainerCluster(getRequired("CLOUDSDK_CONTAINER_CLUSTER"))
}

func getRequired(envKey string) string {
	value, found := os.LookupEnv(envKey)
	if !found {
		log.Panicf("Environment variable %s is required but was not set", envKey)
	}
	return value
}

func init() {
	os.Setenv("CLOUDSDK_CORE_DISABLE_PROMPTS", "True")
}
