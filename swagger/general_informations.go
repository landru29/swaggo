package swagger

import (
	"strings"

	"github.com/landru29/swaggo/parser"
)

// GeneralInformations retrieves general informations
func GeneralInformations(fileAnalyze *parser.FileAnalyze, swag *Swagger) {
	if fileAnalyze.Package == "main" {
		if APIVersion, ok := parser.GetField(fileAnalyze.FileComments, "APIVersion"); ok {
			swag.Info.Version = strings.Join(APIVersion, "")
		}
		if APITitle, ok := parser.GetField(fileAnalyze.FileComments, "APITitle"); ok {
			swag.Info.Title = strings.Join(APITitle, " ")
		}
		if APIDescription, ok := parser.GetField(fileAnalyze.FileComments, "APIDescription"); ok {
			swag.Info.Description = strings.Join(APIDescription, " ")
		}
		if contact, ok := parser.GetField(fileAnalyze.FileComments, "Contact"); ok {
			swag.Info.Contact.Email = strings.Join(contact, ",")
		}
		if termOfServiceURL, ok := parser.GetField(fileAnalyze.FileComments, "TermsOfServiceUrl"); ok {
			swag.Info.TermsOfService = termOfServiceURL[0]
		}
		if license, ok := parser.GetField(fileAnalyze.FileComments, "License"); ok {
			swag.Info.License.Name = strings.Join(license, " ")
		}
		if licenseURL, ok := parser.GetField(fileAnalyze.FileComments, "LicenseUrl"); ok {
			swag.Info.License.URL = licenseURL[0]
		}
		produces := parser.GetFields(fileAnalyze.FileComments, "APIProduces")
		if len(produces) > 0 {
			swag.Produces = []string{}
			for _, produce := range produces {
				swag.Produces = append(swag.Produces, strings.Join(produce, " "))
			}
		}
		consumes := parser.GetFields(fileAnalyze.FileComments, "APIConsumes")
		if len(consumes) > 0 {
			swag.Consumes = []string{}
			for _, consume := range consumes {
				swag.Consumes = append(swag.Consumes, strings.Join(consume, " "))
			}
		}
		swag.Paths = make(PathsStruct)
	}
}
