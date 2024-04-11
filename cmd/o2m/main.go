package main

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"io"
	"os"

	"github.com/jdmcgrath/OpenAPIToMermaid/mermaid"
)

func main() {
	code := run(os.Stdout, os.Args)
	if code != 0 {
		os.Exit(code)
	}
}

const usageText = `Usage: openapi-to-mermaid <Diagram To Generate> <OpenAPI JSON file path>`

func run(w io.Writer, args []string) (code int) {
	if len(args) < 3 {
		fmt.Fprintln(w, usageText)
		return 0
	}
	loader := openapi3.NewLoader()
	spec, err := loader.LoadFromFile(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading OpenAPI specification: %s\n", err)
		os.Exit(1)
	}
	switch args[1] {
	case "er-diagram":
		fmt.Println(mermaid.GenerateERDiagram(spec))
		return 1
	case "td-diagram":
		fmt.Println(mermaid.GenerateTopDownDiagram(spec))
		return 1
	}
	fmt.Fprintln(w, usageText)
	return 0
}
