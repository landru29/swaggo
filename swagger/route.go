package swagger

import (
    "strings"

    "github.com/landru29/swaggo/parser"
)

// @Title Get Users Information
// @Description Get Users Information
// @Accept application/json
// @Param userId path int true "User ID"
// @Success 200 {object} string "Success"
// @Failure 401 {object} string "Access denied"
// @Failure 404 {object} string "Not Found"
// @Resource /users
// @Router /:userId.json [get]

// Route search for new routes
func Route(fileAnalyze *parser.FileAnalyze, swag *Swagger) {
    for _, block := range fileAnalyze.BlockComments {
        oneRoute(block, swag)
    }
}

func oneRoute(comments []string, swag *Swagger) {
    if router, ok := parser.GetField(comments, "Router"); ok {
        if len(router) < 2 {
            return
        }
        method, _, elts, hasRouter := parser.DescID(router)
        if !hasRouter {
            return
        }
        path := elts[0]

        var operation OperationStruct
        operation = OperationStruct{}

        if resource, ok := parser.GetField(comments, "Resource"); ok {
            if len(resource) > 0 {
                path = resource[0] + path
                if tag, ok := GetTag(swag, resource[0]); ok {
                    operation.Tags = append(operation.Tags, tag.Name)
                }
            }
        }

        if _, ok := swag.Paths[path]; !ok {
            swag.Paths[path] = PathItemStruct{}
        }

        if _, ok := swag.Paths[path][method]; !ok {

            if description, ok := parser.GetField(comments, "Description"); ok {
                operation.Description = strings.Join(description, " ")
            }

            if title, ok := parser.GetField(comments, "Title"); ok {
                operation.Summary = strings.Join(title, " ")
            }

            produces := parser.GetFields(comments, "Produces")
            if len(produces) > 0 {
                operation.Produces = []string{}
                for _, produce := range produces {
                    operation.Produces = append(operation.Produces, strings.Join(produce, " "))
                }
            }
            consumes := parser.GetFields(comments, "Accept")
            if len(consumes) > 0 {
                operation.Consumes = []string{}
                for _, consume := range consumes {
                    operation.Consumes = append(operation.Consumes, strings.Join(consume, " "))
                }
            }
            params := parser.GetFields(comments, "Param")
            if len(params) > 0 {
                operation.Parameters = []ParameterStruct{}
                for _, param := range params {
                    if len(param) > 3 {
                        p := ParameterStruct{
                            Name:        param[0],
                            In:          param[1],
                            Description: strings.Trim(param[len(param)-1], `"`),
                            Required:    param[len(param)-2] == "true",
                        }
                        operation.Parameters = append(operation.Parameters, p)
                    }

                }
            }

            swag.Paths[path][method] = operation
        }
    }
}
