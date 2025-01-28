package main

import "C"

import (
	"log"
	"os"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

//export HelmList
func HelmList() {
	settings := cli.New()

	actionConfig := new(action.Configuration)
	// You can pass an empty string instead of settings.Namespace() to list all namespaces
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("Error initializing Helm configuration: %+v", err)
		return
	}

	client := action.NewList(actionConfig)
	client.Deployed = true // Only list deployed releases
	results, err := client.Run()
	if err != nil {
		log.Printf("Error listing Helm releases: %+v", err)
		return
	}

	for _, rel := range results {
		log.Printf("Release: %+v", rel.Name)
	}
}

func main() {
	// Required for `c-shared` build mode
}

