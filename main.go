package main

import (
	"fmt"
	"strings"

	"github.com/landru29/swaggo/swagger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/* Main Command to parse
   command line */
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

/**
 * The Main application really starts here
 */
func mainApp() (err error) {
	err = swagger.ProcessProject(
		viper.GetString("project_folder"),
		viper.GetString("api_host"),
		viper.GetString("api_basepath"),
		strings.Split(viper.GetString("api_scheme"), ","))
	return
}

// et un autre
func main() {
	mainCommand.Execute()
}

func init() {
	flags := mainCommand.Flags()
	flags.String("api-host", "localhost:3000", "API host")
	viper.BindPFlag("api_host", flags.Lookup("api-host"))

	flags.String("api-scheme", "http", "http|https separated by a comma")
	viper.BindPFlag("api_scheme", flags.Lookup("api-scheme"))

	flags.String("api-basepath", "/", "API basepath")
	viper.BindPFlag("api_basepath", flags.Lookup("api-basepath"))

	flags.String("output", "./swagger.json", "Output filename")
	viper.BindPFlag("output", flags.Lookup("output"))

	flags.String("project-folder", ".", "Folder to scan")
	viper.BindPFlag("project_folder", flags.Lookup("project-folder"))
}
