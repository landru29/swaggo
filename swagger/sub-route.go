package swagger

import (
	"errors"
	"fmt"

	"github.com/landru29/swaggo/descriptor"
)

// @SubApi Users [/users]
// @SubApi Allows you access to different features of the users, login, get status etc [/users]

// @SubApi [bob]
// @Resource /user
// @Router /admin

// GetSubRoute search for new routes
// TODO: tool to search resource
func GetSubRoute(swag *Swagger, fileAnalyze *descriptor.FileAnalyze, verbose bool) {
	for _, block := range fileAnalyze.BlockComments {
		oneSubRoute(swag, block, verbose)
	}
}

// CompileSubRoutes build a structure with sub-routes
func CompileSubRoutes(swag *Swagger, verbose bool) {
	if verbose {
		fmt.Printf("# COMPILE SUB-ROUTES\n")
	}
	for _, sub := range swag.AllSubRoutes {
		if verbose {
			if len(sub.ParentResource) > 0 {
				fmt.Printf("    * [%s] has Parent [%s]\n", sub.Resource, sub.ParentResource)
			} else {
				fmt.Printf("    * [%s] has not parent\n", sub.Resource)
			}
		}
		if len(sub.Name) > 0 {
			swag.Tags = append(swag.Tags, sub)
			if verbose {
				fmt.Printf("        + Named Tag [%s] appended to dictionnary\n", sub.Name)
			}
		}
		if parentTag, ok := swag.GetTag(sub.ParentResource); ok {
			if verbose {
				fmt.Printf("        + Linking Res Parent [%s] <=> Current: [%s]\n", parentTag.Resource, sub.ParentResource)
			}
			sub.Parent = parentTag
		}
	}
}

// GetTag find a tag by resource
func (swag *Swagger) GetTag(resource string) (tag *TagStruct, ok bool) {
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
func (swag *Swagger) GetNamedTag(resource string) (tag *TagStruct, ok bool) {
	tag, ok = swag.GetTag(resource)
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

func oneSubRoute(swag *Swagger, comments []string, verbose bool) {
	tag := new(TagStruct)
	subAPI := descriptor.GetFields(comments, "SubApi")
	if len(subAPI) == 0 {
		return
	}
	if verbose {
		fmt.Printf("# SUB-ROUTE")
	}
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
	if verbose {
		fmt.Printf(" [%s]\n", tag.Resource)
		fmt.Printf("    * Name: %s\n", tag.Name)
		fmt.Printf("    * Description: %s\n", tag.Description)
		fmt.Printf("    * URI: %s\n", tag.Router)
		fmt.Printf("    * Child of: %s\n", tag.ParentResource)
	}
	if len(tag.Resource) > 0 {
		swag.AllSubRoutes = append(swag.AllSubRoutes, tag)
	}
}
