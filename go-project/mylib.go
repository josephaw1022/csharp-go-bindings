package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

// Release represents a Helm release object
type Release struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Revision   int    `json:"revision"`
	Updated    string `json:"updated"`
	Status     string `json:"status"`
	Chart      string `json:"chart"`
	AppVersion string `json:"app_version"`
}

//export HelmList
func HelmList() *C.char {
	settings := cli.New()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("Error initializing Helm configuration: %+v", err)
		return C.CString(`[]`)
	}

	client := action.NewList(actionConfig)
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		log.Printf("Error listing Helm releases: %+v", err)
		return C.CString(`[]`)
	}

	// Convert results to custom Release objects
	releases := []Release{}
	for _, rel := range results {
		releases = append(releases, Release{
			Name:       rel.Name,
			Namespace:  rel.Namespace,
			Revision:   rel.Version,
			Updated:    rel.Info.LastDeployed.String(),
			Status:     rel.Info.Status.String(),
			Chart:      rel.Chart.Metadata.Name,
			AppVersion: rel.Chart.Metadata.AppVersion,
		})
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(releases)
	if err != nil {
		log.Printf("Error marshaling JSON: %+v", err)
		return C.CString(`[]`)
	}

	return C.CString(string(jsonData))
}

func main() {

}
