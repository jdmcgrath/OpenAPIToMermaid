package mermaid

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// GenerateTopDownDiagram generates a Mermaid TD diagram from an OpenAPI3 document.
func GenerateTopDownDiagram(doc *openapi3.T) string {
	var builder strings.Builder

	builder.WriteString("graph TD\n")
	paths := doc.Paths.Map()
	for path, pathItem := range paths {
		for method, operation := range pathItem.Operations() {
			// Create a sanitized node ID for Mermaid
			nodeID := fmt.Sprintf("%s_%s", sanitize(path), method)
			nodeLabel := fmt.Sprintf("%s %s", strings.ToUpper(method), path)

			// Define the node for the method and path
			builder.WriteString(fmt.Sprintf("    %s[\"%s\"]\n", nodeID, nodeLabel))

			// If there's a request body, add a node for it
			if operation.RequestBody != nil {
				requestBodyID := fmt.Sprintf("%s_requestBody", nodeID)
				builder.WriteString(fmt.Sprintf("    %s((\"Request Body\"))\n", requestBodyID))
				builder.WriteString(fmt.Sprintf("    %s --> %s\n", nodeID, requestBodyID))
			}
			responses := operation.Responses.Map()
			// Add nodes for each response
			for status := range responses {
				responseID := fmt.Sprintf("%s_response_%s", nodeID, status)
				builder.WriteString(fmt.Sprintf("    %s((\"%s Response\"))\n", responseID, status))
				builder.WriteString(fmt.Sprintf("    %s --> %s\n", nodeID, responseID))
			}
		}
	}

	return builder.String()
}
