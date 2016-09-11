package swagger

// NewSwagger create a new swagger object
func NewSwagger(host string, basePath string, schemes []string) Swagger {
	return Swagger{
		Schemes:  schemes,
		Swagger:  "2.0",
		Host:     host,
		BasePath: basePath,
	}
}
