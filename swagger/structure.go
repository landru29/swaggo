package swagger

// http://swagger.io/specification

//Swagger is the swagger structure
type Swagger struct {
    Consumes            []string                  `json:"consumes"`
    Produces            []string                  `json:"produces"`
    Schemes             []string                  `json:"schemes"`
    Swagger             string                    `json:"swagger"`
    Host                string                    `json:"host"`
    BasePath            string                    `json:"basePath"`
    Info                InfoStruct                `json:"info"`
    Paths               PathsStruct               `json:"paths"`
    Definitions         DefinitionsStruct         `json:"definitions"`
    Parameters          ParametersStruct          `json:"parameters"`
    Responses           ResponsesStruct           `json:"responses"`
    Tags                []TagStruct               `json:"tags"`
    ExternalDocs        ExternalDocsStruct        `json:"externalDocs"`
    SecurityDefinitions SecurityDefinitionsStruct `json:"securityDefinitions"`
    Security            []SecurityStruct          `json:"security"`
}

//InfoStruct is the swagger info structure
type InfoStruct struct {
    Description    string        `json:"description"`
    Title          string        `json:"title"`
    Version        string        `json:"version"`
    TermsOfService string        `json:"termsOfService"`
    License        LicenseStruct `json:"licence"`
    Contact        ContactStruct `json:"contact"`
}

// DefinitionsStruct is the swagger definition structure
type DefinitionsStruct struct {
}

// ParametersStruct is the swagger parameter structure
type ParametersStruct struct {
}

// ResponsesStruct is the swagger response structure
type ResponsesStruct struct {
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
    Name string `json:"Name"`
    URL  string `json:"url"`
}

//ContactStruct is the swagger contact structure
type ContactStruct struct {
    Email string `json:"email"`
}

// ParameterStruct is the swagger parameter structure
type ParameterStruct struct {
}

// OperationStruct is the swagger route definition
type OperationStruct struct {
    Tags         []string           `json:"tags"`
    Summary      []string           `json:"summary"`
    Description  string             `json:"description"`
    ExternalDocs ExternalDocsStruct `json:"externalDocs"`
    OperationID  string             `json:"operationId"`
    Consumes     []string           `json:"consumes"`
    Produces     []string           `json:"produces"`
    Parameters   []ParameterStruct  `json:"parameters"`
    Responses    ResponsesStruct    `json:"responses"`
    Schemes      []string           `json:"schemes"`
    Deprecated   bool               `json:"deprecated"`
    Security     []SecurityStruct   `json:"security"`
}

// PathItemStruct is the swagger routes list (get => ..., post => ...)
type PathItemStruct map[string]OperationStruct

// PathsStruct is the swagget path (path => Route)
type PathsStruct map[string]PathItemStruct
