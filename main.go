package main

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"

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
    filenames, err := getFileList(".")
    if err == nil {
        for _, filename := range filenames {
            fmt.Printf("############ %s ############\n", filename)
            comments, _ := parser.ParseComments(filename)
            for _, comments := range comments {
                fmt.Println(comments)
            }
        }
    }
    swagger := swagger.NewSwagger()
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
