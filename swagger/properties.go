package swagger

import (
	"fmt"
	"regexp"

	"github.com/landru29/swaggo/descriptor"
)

// GetDefinitions search new Definitions
func GetDefinitions(swag *Swagger, fileAnalyze *descriptor.FileAnalyze, verbose bool) {
	for _, block := range fileAnalyze.BlockComments {
		name, definition := oneProperty(block, verbose)
		if len(name) > 0 {
			swag.Definitions[name] = definition
		}
	}
}

// TransformStructureName ...
func TransformStructureName(name string) string {
	re := regexp.MustCompile(`[^\w\d]`)
	return re.ReplaceAllString(name, `_`)
}

func createProperty(lineComments []string) DefinitionStruct {
	prop := DefinitionStruct{}

	propertyType := lineComments[0]
	switch propertyType {
	case "string":
		prop.Type = "string"
	case "number":
		prop.Type = "number"
		prop.Format = lineComments[1]
	case "integer":
		prop.Type = "integer"
		prop.Format = lineComments[1]
	case "boolean":
		prop.Type = "boolean"
	case "array":
		prop.Type = "array"
		item := createProperty(lineComments[1:])
		prop.Items = &item
	default:
		prop.Ref = "#/definitions/" + TransformStructureName(propertyType)
	}
	return prop
}

func oneProperty(comments []string, verbose bool) (string, DefinitionStruct) {
	structureComments, ok := descriptor.GetField(comments, "structure")
	if !ok || len(structureComments) == 0 {
		return "", DefinitionStruct{}
	}
	propertiesComments := descriptor.GetFields(comments, "property")
	if len(propertiesComments) == 0 {
		return "", DefinitionStruct{}
	}
	structureName := TransformStructureName(structureComments[0])
	structureType := "object"
	if len(structureComments) > 1 {
		structureType = structureComments[1]
	}

	result := DefinitionStruct{
		Type: structureType,
	}

	if verbose {
		fmt.Printf("# PROPERTY [%s] as %s\n", structureName, structureType)
	}

	if structureType == "object" {
		properties := make(map[string]DefinitionStruct)

		for _, lineComments := range propertiesComments {
			if len(lineComments) > 1 {
				propertyName := lineComments[0]
				if verbose {
					fmt.Printf("   - %s\n", propertyName)
				}
				properties[propertyName] = createProperty(lineComments[1:])
			}
		}
		result.Properties = properties
	}

	if structureType == "array" && len(propertiesComments) > 0 {
		items := createProperty(propertiesComments[0])
		result.Items = &items
		if verbose {
			fmt.Printf("   - %v\n", propertiesComments[0])
		}
	}

	return structureName, result

}
