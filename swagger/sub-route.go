package swagger

import "github.com/landru29/swaggo/parser"

// @SubApi Users [/users]
// @SubApi Allows you access to different features of the users, login, get status etc [/users]

// SubRoute search for new routes
func SubRoute(fileAnalyze *parser.FileAnalyze, swag *Swagger) {
    for _, block := range fileAnalyze.BlockComments {
        oneSubRoute(block, swag)
    }
}

// GetTag find a tag by resource
func GetTag(swag *Swagger, resource string) (tag TagStruct, ok bool) {
    ok = false
    for _, tagIt := range swag.Tags {
        if tagIt.Resource == resource {
            tag = tagIt
            ok = true
            return
        }
    }
    return
}

func oneSubRoute(comments []string, swag *Swagger) {
    subAPI := parser.GetFields(comments, "SubApi")
    tag := TagStruct{}
    for _, lineComments := range subAPI {
        if res, str, _, ok := parser.DescID(lineComments); ok {
            if len(tag.Name) == 0 {
                tag.Name = str
                tag.Resource = res
            } else if tag.Resource == res {
                tag.Description = str
            }
        }
    }
    if len(tag.Resource) > 0 {
        swag.Tags = append(swag.Tags, tag)
    }
}
