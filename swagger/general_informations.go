package swagger

import (
	"fmt"
	"strings"

	"github.com/landru29/swaggo/descriptor"
)

// GeneralInformations retrieves general informations
func (swag *Swagger) GeneralInformations(fileAnalyze *descriptor.FileAnalyze, verbose bool) {
	if APITitle, ok := descriptor.GetField(fileAnalyze.FileComments, "APITitle"); ok {
		if verbose {
			fmt.Printf("# GENERAL INFORMATION [%s]\n", fileAnalyze.Filename)
		}
		swag.Info.Title = strings.Join(APITitle, " ")

		if APIVersion, ok := descriptor.GetField(fileAnalyze.FileComments, "APIVersion"); ok {
			swag.Info.Version = strings.Join(APIVersion, "")
		}
		if APIDescription, ok := descriptor.GetField(fileAnalyze.FileComments, "APIDescription"); ok {
			swag.Info.Description = strings.Join(APIDescription, " ")
		}
		if contact, ok := descriptor.GetField(fileAnalyze.FileComments, "Contact"); ok {
			swag.Info.Contact.Email = strings.Join(contact, ",")
		}
		if termOfServiceURL, ok := descriptor.GetField(fileAnalyze.FileComments, "TermsOfServiceUrl"); ok {
			swag.Info.TermsOfService = termOfServiceURL[0]
		}
		if license, ok := descriptor.GetField(fileAnalyze.FileComments, "License"); ok {
			swag.Info.License.Name = strings.Join(license, " ")
		}
		if licenseURL, ok := descriptor.GetField(fileAnalyze.FileComments, "LicenseUrl"); ok {
			swag.Info.License.URL = licenseURL[0]
		}
		produces := descriptor.GetFields(fileAnalyze.FileComments, "APIProduces")
		if len(produces) > 0 {
			swag.Produces = []string{}
			for _, produce := range produces {
				swag.Produces = append(swag.Produces, strings.Join(produce, " "))
			}
		}
		consumes := descriptor.GetFields(fileAnalyze.FileComments, "APIConsumes")
		if len(consumes) > 0 {
			swag.Consumes = []string{}
			for _, consume := range consumes {
				swag.Consumes = append(swag.Consumes, strings.Join(consume, " "))
			}
		}
		if verbose {
			fmt.Printf("    * Title: %s\n", swag.Info.Title)
			fmt.Printf("    * Version: %s\n", swag.Info.Version)
			fmt.Printf("    * Description: %s\n", swag.Info.Description)
			fmt.Printf("    * Contact: %s\n", swag.Info.Contact.Email)
			fmt.Printf("    * TermsOfService: %s\n", swag.Info.TermsOfService)
			fmt.Printf("    * Licence: %s\n", swag.Info.License.Name)
			fmt.Printf("    * Licence URL: %s\n", swag.Info.License.URL)
			fmt.Printf("    * Produces: %v\n", swag.Produces)
			fmt.Printf("    * consumes: %v\n", swag.Consumes)
		}
	}

	swag.Paths = make(PathsStruct)
}
