package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

var codeTpl = template.Must(template.New("codeTpl").Parse(`/**
 * This file has been automatically generated.
 * Please use 'go generate' to update this file.
 */

#include "ts3_functions.h"

typedef struct TS3Functions TS3Functions;

{{ range $ }}
// {{.ReturnType}} {{.Name}}({{.Parameters}})
{{.ReturnType}} {{.Name}}(const struct TS3Functions funcs, {{.Parameters}}) {
	{{$delim:=""}}
	return funcs.{{.Name}}({{.Parameters.Names}});
}
{{ end }}
`))

var (
	rxParam    = regexp.MustCompile(`^\s*(?P<type>.+?)(?P<name>[A-Za-z0-9_]+)\s*$`)
	rxFunction = regexp.MustCompile(`\b(?P<returnType>(?:unsigned\s+)?(int|bool|char|byte|void|short|uint64))\s*\(\*(?P<name>[A-Za-z0-9_]+)\)\s*(?P<parameters>\((.+?[A-Za-z0-9_]+(,\s*|\)))*)\;`)
)

type Function struct {
	Name       string
	ReturnType string
	Parameters Parameters
}

type Parameters []*Parameter

func (p Parameters) String() string {
	a := []string{}
	for _, current := range p {
		a = append(a, current.String())
	}
	return strings.Join(a, ", ")
}

func (p Parameters) Names() string {
	a := []string{}
	for _, current := range p {
		a = append(a, current.Name)
	}
	return strings.Join(a, ", ")
}

type Parameter struct {
	Name string
	Type string
}

func (p Parameter) String() string {
	return fmt.Sprintf("%s %s", p.Type, p.Name)
}

func must(err error) {
	if err == nil {
		return
	}

	fmt.Println(err.Error())
	os.Exit(1)
}

func main() {
	ts3functionsBytes, err := ioutil.ReadFile(filepath.Join("pluginsdk", "include", "ts3_functions.h"))
	must(err)

	namedCaptureGroupIndices := map[string]int{}
	for i, name := range rxFunction.SubexpNames() {
		if i != 0 && name != "" {
			namedCaptureGroupIndices[name] = i
		}
	}

	functions := []*Function{}

	for _, submatches := range rxFunction.FindAllSubmatch(ts3functionsBytes, -1) {
		name := string(submatches[namedCaptureGroupIndices["name"]])
		returnType := string(submatches[namedCaptureGroupIndices["returnType"]])
		parametersStr := submatches[namedCaptureGroupIndices["parameters"]]
		parameters := strings.FieldsFunc(string(parametersStr[1:len(parametersStr)-1]), func(c rune) bool {
			return c == ','
		})

		f := new(Function)
		f.Name = name
		f.ReturnType = returnType
		f.Parameters = Parameters{}
		for _, p := range parameters {
			paramFields := rxParam.FindStringSubmatch(p)
			if len(paramFields) < 3 {
				panic("Param regex mismatch")
			}
			paramType := paramFields[1]
			paramName := paramFields[2]
			f.Parameters = append(f.Parameters, &Parameter{
				Name: paramName,
				Type: paramType,
			})
		}
		functions = append(functions, f)
	}

	f, err := os.Create("generated_ts3_function_wrappers.h")
	if err != nil {
		panic("Could not create generated_ts3_function_wrappers.h")
	}
	defer f.Close()
	codeTpl.Execute(f, functions)
}
