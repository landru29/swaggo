package swagger

import (
	"errors"

	"github.com/landru29/swaggo/descriptor"
)

// @SubApi Users [/users]
// @SubApi Allows you access to different features of the users, login, get status etc [/users]

// @SubApi [/admin]
// @Resource /user

// GetSubRoute search for new routes
// TODO: tool to search resource
func GetSubRoute(fileAnalyze *descriptor.FileAnalyze, swag *Swagger) {
	//fmt.Printf("#### Scanning %s\n", fileAnalyze.Filename)
	for _, block := range fileAnalyze.BlockComments {
		oneSubRoute(block, swag)
	}
}

// CompileSubRoutes build a structure with sub-routes
func (swag *Swagger) CompileSubRoutes() {
	for _, sub := range swag.AllSubRoutes {
		//fmt.Printf("Parent: %s - Current: %s\n", sub.ParentResource, sub.Resource)
		if len(sub.Name) > 0 {
			swag.Tags = append(swag.Tags, sub)
		}
		if parentTag, ok := GetTag(swag, sub.ParentResource); ok {
			//fmt.Printf("* Found Res %s\n", parentTag.Resource)
			sub.Parent = parentTag
		}
	}
}

// GetTag find a tag by resource
func GetTag(swag *Swagger, resource string) (tag *TagStruct, ok bool) {
	ok = false
	if len(resource) == 0 {
		return
	}
	for _, tagIt := range swag.AllSubRoutes {
		if tagIt.Resource == resource {
			tag = tagIt
			ok = true
			return
		}
	}
	return
}

// GetNamedTag find the first named tag by resource
func GetNamedTag(swag *Swagger, resource string) (tag *TagStruct, ok bool) {
	tag, ok = GetTag(swag, resource)
	if ok {
		for i := 0; i < 50; i++ {
			if len(tag.Name) > 0 {
				return
			}
			if tag != nil {
				tag = tag.Parent
			}
		}
	}
	return
}

// GetPath get the path of a subroute
func (tag *TagStruct) GetPath() (path string, err error) {
	path = ""
	current := tag
	for i := 0; i < 50; i++ {
		if current == nil {
			return
		}
		path = current.Router + path
		current = current.Parent
	}
	err = errors.New("Infinite loop")
	return
}

func oneSubRoute(comments []string, swag *Swagger) {
	tag := new(TagStruct)
	subAPI := descriptor.GetFields(comments, "SubApi")
	if parentRes, ok := descriptor.GetField(comments, "Resource"); ok {
		tag.ParentResource = parentRes[0]
	}
	for _, lineComments := range subAPI {
		if res, str, _, ok := descriptor.DescID(lineComments); ok {
			if len(tag.Resource) == 0 {
				if len(str) > 0 {
					tag.Name = str
				}
				tag.Resource = res
				tag.Router = res
			} else if tag.Resource == res {
				tag.Description = str
			}
		}
	}
	if router, ok := descriptor.GetField(comments, "Router"); ok {
		tag.Router = router[0]
	}
	if len(tag.Resource) > 0 {
		swag.AllSubRoutes = append(swag.AllSubRoutes, tag)
	}
}
