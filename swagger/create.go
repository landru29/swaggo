package swagger

import "fmt"

// NewSwagger create a new swagger object
func NewSwagger(host string, basePath string, schemes []string, verbose bool) (swag *Swagger) {
	swag = new(Swagger)
	swag.Schemes = schemes
	swag.Swagger = "2.0"
	swag.Host = host
	swag.BasePath = basePath
	if verbose {
		fmt.Printf("# CREATING A NEW SWAGGER STRUCTURE\n")
		fmt.Printf("    * Schemes: %v\n", swag.Schemes)
		fmt.Printf("    * Swagger version: %s\n", swag.Swagger)
		fmt.Printf("    * Host: %s\n", swag.Host)
		fmt.Printf("    * Base path: %s\n", swag.BasePath)
	}
	return
}
