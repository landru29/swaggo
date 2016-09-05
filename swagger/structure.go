package swagger

//Swagger is the swagger structure
type Swagger struct {
    Consumes []string `json:"consumes"`
    Produces []string `json:"produces"`
    Scheme   []string `json:"scheme"`
    Swagger  string   `json:"swagger"`
    Host     string   `json:"host"`
    BasePath string   `json:"basePath"`
    Info     Info     `json:"info"`
}

//Info is the swagger info structure
type Info struct {
    Description    string  `json:"description"`
    Title          string  `json:"title"`
    Version        string  `json:"version"`
    TermsOfService string  `json:"termsOfService"`
    License        License `json:"licence"`
    Contact        Contact `json:"contact"`
}

//License is the swagger licence structure
type License struct {
    Name string `json:"Name"`
    URL  string `json:"url"`
}

//Contact is the swagger contact structure
type Contact struct {
    Email string `json:"email"`
}
