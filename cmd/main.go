package main

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/jdmcgrath/OpenAPIToMermaid/mermaid"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: openapi-to-mermaid <OpenAPI JSON file path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	loader := openapi3.NewLoader()
	spec, err := loader.LoadFromFile(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading OpenAPI specification: %s\n", err)
		os.Exit(1)
	}

	erDiagram := mermaid.GenerateERDiagram(spec)
	fmt.Println(erDiagram)
}
