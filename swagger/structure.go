package swagger

// http://swagger.io/specification

//Swagger is the swagger structure
type Swagger struct {
	Consumes            []string                    `json:"consumes,omitempty"`
	Produces            []string                    `json:"produces,omitempty"`
	Schemes             []string                    `json:"schemes,omitempty"`
	Swagger             string                      `json:"swagger,omitempty"`
	Host                string                      `json:"host,omitempty"`
	BasePath            string                      `json:"basePath,omitempty"`
	Info                InfoStruct                  `json:"info,omitempty"`
	Paths               map[string]PathItemStruct   `json:"paths,omitempty"`
	Definitions         map[string]DefinitionStruct `json:"definitions,omitempty"`
	Parameters          []ParameterStruct           `json:"parameters,omitempty"`
	Responses           map[string]ResponseStruct   `json:"responses,omitempty"`
	Tags                []*TagStruct                `json:"tags,omitempty"`
	AllSubRoutes        []*TagStruct                `json:"-"`
	ExternalDocs        ExternalDocsStruct          `json:"externalDocs,omitempty"`
	SecurityDefinitions SecurityDefinitionsStruct   `json:"securityDefinitions,omitempty"`
	Security            []map[string][]string       `json:"security,omitempty"`
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

// DefinitionStruct is the swagger definition structure
type DefinitionStruct struct {
	Type        string                      `json:"type,omitempty"`
	Enum        []string                    `json:"enum,omitempty"`
	Format      string                      `json:"format,omitempty"`
	Description string                      `json:"description,omitempty"`
	Items       *DefinitionStruct           `json:"items,omitempty"`
	Properties  map[string]DefinitionStruct `json:"properties,omitempty"`
	Ref         string                      `json:"$ref,omitempty"`
	XML         XMLStruct                   `json:"xml,omitempty"`
}

// ResponseStruct is the swagger response structure
type ResponseStruct struct {
	Description string                   `json:"description,omitempty"`
	Schema      SchemaStruct             `json:"schema,omitempty"`
	Headers     map[string]HeaderStruct  `json:"headers,omitempty"`
	Examples    map[string]ExampleStruct `json:"examples,omitempty"`
}

// XMLStruct is the swagger xml structure
type XMLStruct struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Prefix    string `json:"prefix,omitempty"`
	Attribute bool   `json:"attribute,omitempty"`
	Wrapped   bool   `json:"wrapped,omitempty"`
}

// SchemaStruct is the swagger schema structure
type SchemaStruct struct {
	Ref              string             `json:"$ref,omitempty"`
	Title            string             `json:"title,omitempty"`
	Description      string             `json:"description,omitempty"`
	Default          interface{}        `json:"default,omitempty"`
	MultipleOf       interface{}        `json:"multipleOf,omitempty"`
	Maximum          float64            `json:"maximum,omitempty"`
	ExclusiveMaximum bool               `json:"exclusiveMaximum,omitempty"`
	Minimum          float64            `json:"minimum,omitempty"`
	ExclusiveMinimum bool               `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                `json:"maxLength,omitempty"`
	MinLength        int                `json:"minLength,omitempty"`
	Pattern          interface{}        `json:"pattern,omitempty"`
	MinItems         int                `json:"minItems,omitempty"`
	MaxItems         int                `json:"maxItems,omitempty"`
	UniqueItems      interface{}        `json:"uniqueItems,omitempty"`
	MinProperties    interface{}        `json:"minProperties,omitempty"`
	MaxProperties    interface{}        `json:"maxProperties,omitempty"`
	Properties       interface{}        `json:"properties,omitempty"`
	Required         bool               `json:"required,omitempty"`
	Enum             []interface{}      `json:"enum,omitempty"`
	Type             string             `json:"type,omitempty"`
	Discriminator    string             `json:"discriminator,omitempty"`
	ReadOnly         bool               `json:"readOnly,omitempty"`
	XML              XMLStruct          `json:"xml,omitempty"`
	ExternalDocs     ExternalDocsStruct `json:"externalDocs,omitempty"`
	Example          interface{}        `json:"example,omitempty"`
}

// HeaderStruct is the swagger header structure
type HeaderStruct struct {
	Description      string                `json:"description,omitempty"`
	Type             string                `json:"type,omitempty"`
	Format           string                `json:"format,omitempty"`
	Items            map[string]ItemStruct `json:"items,omitempty"`
	CollectionFormat string                `json:"collectionFormat,omitempty"`
	Default          interface{}           `json:"default,omitempty"`
	Maximum          float32               `json:"maximum,omitempty"`
	ExclusiveMaximum bool                  `json:"exclusiveMaximum,omitempty"`
	Minimum          float32               `json:"minimum,omitempty"`
	ExclusiveMinimum bool                  `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                   `json:"maxLength,omitempty"`
	MinLength        int                   `json:"minLength,omitempty"`
	Pattern          string                `json:"pattern,omitempty"`
	MaxItems         int                   `json:"maxItems,omitempty"`
	MinItems         int                   `json:"minItems,omitempty"`
	UniqueItems      bool                  `json:"uniqueItems,omitempty"`
	Enum             []interface{}         `json:"enum,omitempty"`
	MultipleOf       float32               `json:"multipleOf,omitempty"`
}

// ExampleStruct is the swagger example structure
type ExampleStruct struct {
}

// ItemStruct is the swagger item structure
type ItemStruct struct {
}

// ExternalDocsStruct is the swagger external doc structure
type ExternalDocsStruct struct {
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}

// SecurityDefinitionsStruct is the swagger security definition structure
type SecurityDefinitionsStruct struct {
}

// TagStruct is the swagger tag structure
type TagStruct struct {
	Name           string             `json:"Name,omitempty"`
	Description    string             `json:"description,omitempty"`
	ExternalDocs   ExternalDocsStruct `json:"externalDocs,omitempty"`
	Resource       string             `json:"-"`
	Router         string             `json:"-"`
	Parent         *TagStruct         `json:"-"`
	ParentResource string             `json:"-"`
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
	Name        string        `json:"name,omitempty"`
	In          string        `json:"in,omitempty"` //Possible values are "query", "header", "path", "formData" or "body".
	Description string        `json:"description,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Type        string        `json:"type,omitempty"`   //The value MUST be one of "string", "number", "integer", "boolean", "array" or "file"
	Schema      *SchemaStruct `json:"schema,omitempty"` //If in is "body"
}

// OperationStruct is the swagger route definition
type OperationStruct struct {
	Tags         []string                  `json:"tags,omitempty"`
	Summary      string                    `json:"summary,omitempty"`
	Description  string                    `json:"description,omitempty"`
	ExternalDocs ExternalDocsStruct        `json:"externalDocs,omitempty"`
	OperationID  string                    `json:"operationId,omitempty"`
	Consumes     []string                  `json:"consumes,omitempty"`
	Produces     []string                  `json:"produces,omitempty"`
	Parameters   []ParameterStruct         `json:"parameters,omitempty"`
	Responses    map[string]ResponseStruct `json:"responses,omitempty"`
	Schemes      []string                  `json:"schemes,omitempty"`
	Deprecated   bool                      `json:"deprecated,omitempty"`
	Security     []map[string][]string     `json:"security,omitempty"`
}

// PathItemStruct is the swagger routes list (get => ..., post => ...)
type PathItemStruct map[string]OperationStruct
