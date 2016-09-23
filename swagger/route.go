package swagger

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/landru29/swaggo/descriptor"
)

// @Title Get Users Information
// @Description Get Users Information
// @Accept application/json
// @Param userId path integer true "User ID"
// @Success 200 {object} string "Success"
// @Failure 401 {object} string "Access denied"
// @Failure 404 {object} string "Not Found"
// @Resource /users
// @Router /:userId.json [get]

// Route search for new routes
func (swag *Swagger) Route(fileAnalyze *descriptor.FileAnalyze, verbose bool) {
	for _, block := range fileAnalyze.BlockComments {
		swag.oneRoute(block, verbose)
	}
}

func transformParam(b []byte) []byte {
	s := string(b)
	s = s[1:len(s)]
	t := "{" + s + "}"
	return []byte(t)
}

func replaceParams(in string) (out string) {
	r := regexp.MustCompile(`:([^\/]*)`)
	inb := []byte(in)
	out = string(r.ReplaceAllFunc(inb, transformParam))
	return
}

func (swag *Swagger) oneRoute(comments []string, verbose bool) {
	if router, ok := descriptor.GetField(comments, "Router"); ok {

		if len(router) < 2 {
			return
		}
		method, _, elts, hasRouter := descriptor.DescID(router)
		if !hasRouter {
			return
		}

		path := elts[0]

		var operation OperationStruct
		operation = OperationStruct{}

		if resource, ok := descriptor.GetField(comments, "Resource"); ok {
			if len(resource) > 0 {
				if tag, ok := swag.GetTag(resource[0]); ok {
					if subRoutePath, err := tag.GetPath(); err == nil {
						path = subRoutePath + path
					} else {
						path = subRoutePath + path + "[infiniteLoop]"
					}
					if namedTag, ok := swag.GetNamedTag(resource[0]); ok {
						operation.Tags = append(operation.Tags, namedTag.Name)
					}
				}
			}
		}

		path = replaceParams(path)

		if verbose {
			fmt.Printf("# ROUTE [%s] %s\n", strings.ToUpper(method), path)
		}

		if _, ok := swag.Paths[path]; !ok {
			swag.Paths[path] = PathItemStruct{}
		}

		if _, ok := swag.Paths[path][method]; !ok {
			if _, ok := descriptor.GetField(comments, "Deprecated"); ok {
				operation.Deprecated = true
			}

			if description, ok := descriptor.GetField(comments, "Description"); ok {
				operation.Description = strings.Join(description, " ")
			}

			if title, ok := descriptor.GetField(comments, "Title"); ok {
				operation.Summary = strings.Join(title, " ")
			}

			produces := descriptor.GetFields(comments, "Produces")
			if len(produces) > 0 {
				operation.Produces = []string{}
				for _, produce := range produces {
					operation.Produces = append(operation.Produces, strings.Join(produce, " "))
				}
			}
			consumes := descriptor.GetFields(comments, "Accept")
			if len(consumes) > 0 {
				operation.Consumes = []string{}
				for _, consume := range consumes {
					operation.Consumes = append(operation.Consumes, strings.Join(consume, " "))
				}
			}
			params := descriptor.GetFields(comments, "Param")
			if len(params) > 0 {
				operation.Parameters = []ParameterStruct{}
				for _, param := range params {
					if len(param) > 3 {
						p := ParameterStruct{
							Name:        param[0],
							In:          param[1],
							Type:        param[2],
							Description: strings.Trim(param[len(param)-1], `"`),
							Required:    param[len(param)-2] == "true",
						}
						operation.Parameters = append(operation.Parameters, p)
					}

				}
			}

			if verbose {
				fmt.Printf("    * Title: %s\n", operation.Summary)
				fmt.Printf("    * Description: %s\n", operation.Description)
				fmt.Printf("    * Produces: %v\n", operation.Produces)
				fmt.Printf("    * consumes: %v\n", operation.Consumes)
				for _, param := range operation.Parameters {
					fmt.Printf("    * Parameter: [%s] in (%s) type {%s}\n", param.Name, param.In, param.Type)
				}
			}

			swag.Paths[path][method] = operation
		}
	}
}
