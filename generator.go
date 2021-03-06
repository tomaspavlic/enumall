package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const generatorTemplate = `// Code generated by "enumall -type={{.TypeName}}"; DO NOT EDIT.

package {{.PackageName}}

var All{{.TypeName}} = []{{.TypeName}}{
{{range .Values}}	{{.}},
{{end}}}`

// generator holds information for generating all enum variable.
type generator struct {
	PackageName string
	TypeName    string
	Values      []string
}

// generate creates the output file containing values of given type name.
func (g *generator) generate() {
	if len(g.Values) > 0 {
		baseName := fmt.Sprintf("%s_all.go", strings.ToLower(g.TypeName))
		wd, _ := os.Getwd()
		outputName := filepath.Join(wd, baseName)
		f, err := os.Create(outputName)
		handleError(err)
		// parsing well known template should not fail
		t, _ := tmpl.Parse(generatorTemplate)

		// write to the output file
		err = t.Execute(f, g)
		handleError(err)
	}
}
