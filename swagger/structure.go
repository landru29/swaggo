package swagger

// http://swagger.io/specification

//Swagger is the swagger structure
type Swagger struct {
	Consumes            []string                  `json:"consumes,omitempty"`
	Produces            []string                  `json:"produces,omitempty"`
	Schemes             []string                  `json:"schemes,omitempty"`
	Swagger             string                    `json:"swagger,omitempty"`
	Host                string                    `json:"host,omitempty"`
	BasePath            string                    `json:"basePath,omitempty"`
	Info                InfoStruct                `json:"info,omitempty"`
	Paths               PathsStruct               `json:"paths,omitempty"`
	Definitions         DefinitionsStruct         `json:"definitions,omitempty"`
	Parameters          []ParameterStruct         `json:"parameters,omitempty"`
	Responses           ResponsesStruct           `json:"responses,omitempty"`
	Tags                []TagStruct               `json:"tags,omitempty"`
	ExternalDocs        ExternalDocsStruct        `json:"externalDocs,omitempty"`
	SecurityDefinitions SecurityDefinitionsStruct `json:"securityDefinitions,omitempty"`
	Security            []SecurityStruct          `json:"security,omitempty"`
}

//InfoStruct is the swagger info structure
type InfoStruct struct {
	Description    string        `json:"description,omitempty"`
	Title          string        `json:"title,omitempty"`
	Version        string        `json:"version,omitempty"`
	TermsOfService string        `json:"termsOfService,omitempty"`
	License        LicenseStruct `json:"licence,omitempty"`
	Contact        ContactStruct `json:"contact,omitempty"`
}

// DefinitionsStruct is the swagger definition structure
type DefinitionsStruct struct {
}

// ResponseStruct is the swagger response structure
type ResponseStruct struct {
	Description string         `json:"description,omitempty"`
	Schema      SchemaStruct   `json:"schema,omitempty"`
	Headers     HeadersStruct  `json:"headers,omitempty"`
	Examples    ExamplesStruct `json:"examples,omitempty"`
}

// SchemaStruct is the swagger schema structure
type SchemaStruct struct {
}

// HeaderStruct is the swagger header structure
type HeaderStruct struct {
	Description      string      `json:"description,omitempty"`
	Type             string      `json:"type,omitempty"`
	Format           string      `json:"format,omitempty"`
	Items            ItemsStruct `json:"items,omitempty"`
	CollectionFormat string      `json:"collectionFormat,omitempty"`
	//Default * `json:"default,omitempty"`
	Maximum          float32 `json:"maximum,omitempty"`
	ExclusiveMaximum bool    `json:"exclusiveMaximum,omitempty"`
	Minimum          float32 `json:"minimum,omitempty"`
	ExclusiveMinimum bool    `json:"exclusiveMinimum,omitempty"`
	MaxLength        int     `json:"maxLength,omitempty"`
	MinLength        int     `json:"minLength,omitempty"`
	Pattern          string  `json:"pattern,omitempty"`
	MaxItems         int     `json:"maxItems,omitempty"`
	MinItems         int     `json:"minItems,omitempty"`
	UniqueItems      bool    `json:"uniqueItems,omitempty"`
	//Enum [] `json:"enum,omitempty"`
	MultipleOf float32 `json:"multipleOf,omitempty"`
}

// ExampleStruct is the swagger example structure
type ExampleStruct struct {
}

// ItemStruct is the swagger item structure
type ItemStruct struct {
}

// ExternalDocsStruct is the swagger external doc structure
type ExternalDocsStruct struct {
}

// SecurityDefinitionsStruct is the swagger security definition structure
type SecurityDefinitionsStruct struct {
}

// SecurityStruct is the swagger security structure
type SecurityStruct struct {
}

// TagStruct is the swagger tag structure
type TagStruct struct {
}

//LicenseStruct is the swagger licence structure
type LicenseStruct struct {
	Name string `json:"Name,omitempty"`
	URL  string `json:"url,omitempty"`
}

//ContactStruct is the swagger contact structure
type ContactStruct struct {
	Email string `json:"email,omitempty"`
}

// ParameterStruct is the swagger parameter structure
type ParameterStruct struct {
	Name        string `json:"name,omitempty"`
	In          string `json:"in,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
}

// OperationStruct is the swagger route definition
type OperationStruct struct {
	Tags         []string           `json:"tags,omitempty"`
	Summary      string             `json:"summary,omitempty"`
	Description  string             `json:"description,omitempty"`
	ExternalDocs ExternalDocsStruct `json:"externalDocs,omitempty"`
	OperationID  string             `json:"operationId,omitempty"`
	Consumes     []string           `json:"consumes,omitempty"`
	Produces     []string           `json:"produces,omitempty"`
	Parameters   []ParameterStruct  `json:"parameters,omitempty"`
	Responses    ResponsesStruct    `json:"responses,omitempty"`
	Schemes      []string           `json:"schemes,omitempty"`
	Deprecated   bool               `json:"deprecated,omitempty"`
	Security     []SecurityStruct   `json:"security,omitempty"`
}

// PathItemStruct is the swagger routes list (get => ..., post => ...)
type PathItemStruct map[string]OperationStruct

// PathsStruct is the swagger path (path => Route)
type PathsStruct map[string]PathItemStruct

// HeadersStruct is the mapper for headers
type HeadersStruct map[string]HeaderStruct

// ResponsesStruct is the mapper for responses
type ResponsesStruct map[string]ResponseStruct

// ExamplesStruct is the mapper for examples
type ExamplesStruct map[string]ExampleStruct

// ItemsStruct is the mapper for items
type ItemsStruct map[string]ItemStruct
