package swagger

import (
    "strings"

    "github.com/spf13/viper"
)

// NewSwagger create a new swagger object
func NewSwagger() Swagger {
    return Swagger{
        Schemes:  strings.Split(viper.GetString("api_scheme"), ","),
        Swagger:  "2.0",
        Host:     viper.GetString("api_host"),
        BasePath: viper.GetString("api_basepath"),
    }
}
