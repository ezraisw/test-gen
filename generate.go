package testgen

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/ezraisw/test-gen/internal/analyzer"
	"github.com/sanity-io/litter"
)

type Test struct {
	Name        string
	MockConfigs []*analyzer.MockConfig
	Run         func(mocks []any)
	TestRun     string
	TestRunPkg  string
}

var fileTmpl = template.Must(template.New("file").Parse(
	`// Code generated by test-gen. DO NOT EDIT.
package {{ .packageName }}

import (
	{{ range .imports }}"{{ . }}"
	{{ end }}
)

{{ range .testFuncs }}{{ . }}
{{ end }}
`))

var testFuncTmpl = template.Must(template.New("test-func").Parse(
	`func {{ .testName }}(t *testing.T) {
	cases := []struct{
		setup func(ctrl *gomock.Controller) []any
	}{
		{{ range .testCases }}{{ . }}
		{{ end }}
	}

	for i, c := range cases {
		c := c
		t.Run("Case #"+strconv.Itoa(i), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mocks := c.setup(ctrl)

			{{ .testRun }}(mocks)
		})
	}
}`))

var testCaseTmpl = template.Must(template.New("test-case").Parse(`{
	setup: func(ctrl *gomock.Controller) []any {
		{{ range .mockVarDecls }}{{ . }}
		{{ end }}

		{{ range .mockCallStmts }}{{ . }}
		{{ end }}

		return []any{
			{{ range .mockVarNames }}{{ . }},
			{{ end }}
		}
	},
},`))

var sq = litter.Options{
	Compact:                   true,
	HidePrivateFields:         true,
	HideZeroValues:            true,
	DisablePointerReplacement: true,
}

func Generate(pkgName, fileName string, tests []*Test) {
	importSet := map[string]struct{}{
		"go.uber.org/mock/gomock": {},
		"testing":                 {},
		"strconv":                 {},
	}

	testFuncs := make([]string, 0, len(tests))
	for _, test := range tests {
		if !strings.HasPrefix(test.Name, "Test") {
			panic("test name must start with 'Test'")
		}

		res := analyzer.Multiply(test.MockConfigs, test.Run)

		testCases := make([]string, 0, len(res))
		for _, capturedCalls := range res {
			mockVars := make(map[string]string)
			mockCallStmts := make([]string, 0, len(capturedCalls))

			for _, capturedCall := range capturedCalls {
				declName := getVarName("mock", capturedCall.Type)
				if _, ok := mockVars[declName]; !ok {
					pkgName := filepath.Base(capturedCall.Type.PkgPath())
					varDecl := declName + " := " + pkgName + ".New" + capturedCall.Type.Name() + "(ctrl)"
					mockVars[declName] = varDecl
				}

				callArgs := make([]string, 0, capturedCall.MethodType.NumIn())
				for i := 0; i < capturedCall.MethodType.NumIn(); i++ {
					callArgs = append(callArgs, "gomock.Any()")
				}

				retValStrs := make([]string, 0, len(capturedCall.Returns))
				for _, ret := range capturedCall.Returns {
					var retValStr string
					if err, ok := ret.(error); ok {
						retValStr = "errors.New(\"" + err.Error() + "\")"
						importSet["errors"] = struct{}{}
					} else {
						retValStr = sq.Sdump(ret)
					}
					retValStrs = append(retValStrs, retValStr)
				}

				importSet[capturedCall.Type.PkgPath()] = struct{}{}
				callStmt := declName + ".EXPECT()." + capturedCall.MethodName + "(" + strings.Join(callArgs, ", ") + ")"
				if len(retValStrs) > 0 {
					callStmt += ".Return(" + strings.Join(retValStrs, ", ") + ")"
				}
				mockCallStmts = append(mockCallStmts, callStmt)
			}

			mockVarDecls := make([]string, 0, len(mockVars))
			mockVarNames := make([]string, 0, len(mockVars))

			for mockVarName, mockVarDecl := range mockVars {
				mockVarNames = append(mockVarNames, mockVarName)
				mockVarDecls = append(mockVarDecls, mockVarDecl)
			}

			sort.Strings(mockVarDecls)
			sort.Strings(mockVarNames)

			var buf bytes.Buffer
			testCaseTmpl.Execute(&buf, map[string]any{
				"mockVarDecls":  mockVarDecls,
				"mockCallStmts": mockCallStmts,
				"mockVarNames":  mockVarNames,
			})

			testCases = append(testCases, buf.String())
		}

		testRun := test.TestRun
		if test.TestRunPkg != "" {
			importSet[test.TestRunPkg] = struct{}{}
			testRun = filepath.Base(test.TestRunPkg) + "." + testRun
		}

		var buf bytes.Buffer
		testFuncTmpl.Execute(&buf, map[string]any{
			"testName":  test.Name,
			"testCases": testCases,
			"testRun":   testRun,
		})

		testFuncs = append(testFuncs, buf.String())
	}

	imports := make([]string, 0, len(importSet))
	for im := range importSet {
		imports = append(imports, im)
	}
	sort.Strings(imports)

	var buf bytes.Buffer
	fileTmpl.Execute(&buf, map[string]any{
		"packageName": pkgName,
		"imports":     imports,
		"testFuncs":   testFuncs,
	})

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	file.Write(formatted)
}

func getVarName(prefix string, rt reflect.Type) string {
	name := rt.Name()

	if prefix != "" {
		return prefix + name
	}

	return strings.ToLower(name[0:1]) + name[1:]
}