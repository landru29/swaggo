package swagger

import "github.com/landru29/swaggo/descriptor"

// @SubApi Users [/users]
// @SubApi Allows you access to different features of the users, login, get status etc [/users]

// @SubApi [/admin]
// @Resource /user

type subRoute struct {
	Name        string
	Description string
	Resource    string
	parent      *subRoute
	ParentStr   string
}

// GetSubRoute search for new routes
// TODO: build a structure of resources
// TODO: Build the Tags
// TODO: tool to search resource
// TODO: add ID on resources (to diff 2 sub resource with same name)
func GetSubRoute(fileAnalyze *descriptor.FileAnalyze, swag *Swagger) {
	subroutes := []subRoute{}
	for _, block := range fileAnalyze.BlockComments {
		oneSubRoute(block, &subroutes)
	}
	for _, sub := range subroutes {
		tag := TagStruct{
			Name:        sub.Name,
			Description: sub.Description,
			Resource:    sub.Resource,
		}
		swag.Tags = append(swag.Tags, tag)
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

func oneSubRoute(comments []string, subs *[]subRoute) {
	tag := subRoute{}
	subAPI := descriptor.GetFields(comments, "SubApi")
	if parentRes, ok := descriptor.GetField(comments, "Resource"); ok {
		tag.ParentStr = parentRes[0]
	}
	for _, lineComments := range subAPI {
		if res, str, _, ok := descriptor.DescID(lineComments); ok {
			if len(tag.Resource) == 0 {
				if len(str) > 0 {
					tag.Name = str
				}
				tag.Resource = res
			} else if tag.Resource == res {
				tag.Description = str
			}
		}
	}
	if len(tag.Resource) > 0 {
		*subs = append(*subs, tag)
	}
}
