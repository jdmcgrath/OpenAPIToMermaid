package mermaid

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// sanitize prepares a string for use as a Mermaid node ID.
func sanitize(input string) string {
	return strings.NewReplacer("/", "_", "{", "", "}", "").Replace(input)
}

func getPropertyType(prop *openapi3.SchemaRef) string {
	if prop.Ref != "" {
		return "object"
	} else if prop.Value.Type == "array" {
		return "array"
	}
	return prop.Value.Type
}

func extractRefName(ref string) string {
	// Assuming the ref format is "#/components/schemas/SchemaName"
	parts := strings.Split(ref, "/")
	return parts[len(parts)-1] // The last part is the schema name
}

func getRelationship(requiredFields []string, propName string, prop *openapi3.SchemaRef) string {
	// Determine cardinality based on whether the property is an array and if it's required
	isRequired := stringInSlice(propName, requiredFields)
	if prop.Value.Type == "array" {
		if isRequired {
			return "}o --o{" // One or more
		}
		return "}o --o|" // Zero or more
	}
	if isRequired {
		return "|| -- ||" // Exactly one
	}
	return "|o -- o|" // Zero or one
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
