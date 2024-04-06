package mermaid

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// GenerateERDiagram generates a Mermaid ER diagram from an OpenAPISpec
func GenerateERDiagram(spec *openapi3.T) string {
	var builder strings.Builder
	builder.WriteString("erDiagram\n")

	// Collect and sort schema names
	var schemaNames []string
	if spec.Components != nil && spec.Components.Schemas != nil {
		for name := range spec.Components.Schemas {
			schemaNames = append(schemaNames, name)
		}
	} else {
		log.Fatalf("Components or Schemas is nil")
	}
	sort.Strings(schemaNames)

	// Generate entity definitions
	for _, name := range schemaNames {
		schema := spec.Components.Schemas[name]
		builder.WriteString(fmt.Sprintf("    %s {\n", name))
		for propName, prop := range schema.Value.Properties {
			propType := getPropertyType(prop)
			builder.WriteString(fmt.Sprintf("        %s %s\n", propType, propName))
		}
		builder.WriteString("    }\n")
	}

	// Generate relationships
	for _, name := range schemaNames {
		schema := spec.Components.Schemas[name]
		for propName, prop := range schema.Value.Properties {
			if prop.Ref != "" {
				refName := extractRefName(prop.Ref)
				relationship := getRelationship(schema.Value.Required, propName, prop)
				builder.WriteString(fmt.Sprintf("    %s %s %s : \"%s\"\n", name, relationship, refName, propName))
			} else if prop.Value.Type == "array" && prop.Value.Items != nil && prop.Value.Items.Ref != "" {
				refName := extractRefName(prop.Value.Items.Ref)
				relationship := getRelationship(schema.Value.Required, propName, prop)
				builder.WriteString(fmt.Sprintf("    %s %s %s : \"%s[]\"\n", name, relationship, refName, propName))
			}
		}
	}

	return builder.String()
}
