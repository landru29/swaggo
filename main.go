package main

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
    "strings"

    "github.com/landru29/swaggo/parser"
    "github.com/landru29/swaggo/swagger"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

/* Gros comments
sur deux lignes */

var mainCommand = &cobra.Command{
    Use:   "api-go",
    Short: "API by noopy",
    Long:  "Full API by noopy",
    Run: func(cmd *cobra.Command, args []string) {
        viper.SetEnvPrefix("swaggo")
        viper.AutomaticEnv()
        // Application statup here
        err := mainApp()
        if err != nil {
            fmt.Println(err)
        }
    },
}

// ceci est un test
func getFileList(searchDir string) (fileList []string, err error) {
    fileList = []string{}
    var goFileRegExp = regexp.MustCompile(`\.go$`)
    var notHiddentDirRegExp = regexp.MustCompile(`\/\.\w+|^\.\w+`)
    err = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        if !f.IsDir() && goFileRegExp.MatchString(f.Name()) && !notHiddentDirRegExp.MatchString(path) {
            fileList = append(fileList, path)
        }
        return nil
    })
    return
}

func mainApp() (err error) {
    swagger := swagger.NewSwagger()
    filenames, err := getFileList(".")
    if err == nil {
        for _, filename := range filenames {
            fileAnalyze, _ := parser.ParseComments(filename)
            if fileAnalyze.Package == "main" {
                if APIVersion, ok := parser.GetField(fileAnalyze.Comments, "APIVersion"); ok {
                    swagger.Info.Version = strings.Join(APIVersion, "")
                }
                if APITitle, ok := parser.GetField(fileAnalyze.Comments, "APITitle"); ok {
                    swagger.Info.Title = strings.Join(APITitle, " ")
                }
                if APIDescription, ok := parser.GetField(fileAnalyze.Comments, "APIDescription"); ok {
                    swagger.Info.Description = strings.Join(APIDescription, " ")
                }
                if contact, ok := parser.GetField(fileAnalyze.Comments, "Contact"); ok {
                    swagger.Info.Contact.Email = strings.Join(contact, ",")
                }
                if termOfServiceURL, ok := parser.GetField(fileAnalyze.Comments, "TermsOfServiceUrl"); ok {
                    swagger.Info.TermsOfService = termOfServiceURL[0]
                }
                if license, ok := parser.GetField(fileAnalyze.Comments, "License"); ok {
                    swagger.Info.License.Name = strings.Join(license, " ")
                }
                if licenseURL, ok := parser.GetField(fileAnalyze.Comments, "LicenseUrl"); ok {
                    swagger.Info.License.URL = licenseURL[0]
                }
                produces := parser.GetFields(fileAnalyze.Comments, "APIProduces")
                if len(produces) > 0 {
                    swagger.Produces = []string{}
                    for _, produce := range produces {
                        swagger.Produces = append(swagger.Produces, strings.Join(produce, " "))
                    }
                }
                consumes := parser.GetFields(fileAnalyze.Comments, "APIConsumes")
                if len(produces) > 0 {
                    swagger.Consumes = []string{}
                    for _, consume := range consumes {
                        swagger.Consumes = append(swagger.Consumes, strings.Join(consume, " "))
                    }
                }
            }

            fmt.Printf("############ %s ############ %s #\n", filename, fileAnalyze.Package)
            for _, comment := range fileAnalyze.Comments {
                fmt.Println(comment)
            }
        }
    }

    err = swagger.Save()
    return
}

// et un autre
func main() {
    mainCommand.Execute()
}

func init() {
    flags := mainCommand.Flags()
    flags.String("api-host", "localhost", "API host")
    viper.BindPFlag("api_host", flags.Lookup("api-host"))

    flags.String("api-scheme", "http", "http|https separated by a comma")
    viper.BindPFlag("api_scheme", flags.Lookup("api-scheme"))

    flags.String("api-basepath", "/", "API basepath")
    viper.BindPFlag("api_basepath", flags.Lookup("api-basepath"))

    flags.String("output", "./swagger.json", "Output filename")
    viper.BindPFlag("output", flags.Lookup("output"))
}
