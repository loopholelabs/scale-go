/*
	Copyright 2023 Loophole Labs
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at
		   http://www.apache.org/licenses/LICENSE-2.0
	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package typescript

import (
	"text/template"

	"github.com/loopholelabs/scale/compile/typescript/templates"
)

type Generator struct {
	template *template.Template
}

func New() *Generator {
	return &Generator{
		template: template.Must(template.New("main").ParseFS(templates.FS, "*ts.templ")),
	}
}

// func (g *Generator) GenerateTypescriptPackageJSON(schema *scalefile.Schema, parsedSignatureDependency *ParsedDependency, functionPath string, dependencies []*scalefunc.Dependency, packageName string, packageVersion string) ([]byte, error) {
//	if parsedSignatureDependency.Path != "" && !strings.HasPrefix(parsedSignatureDependency.Path, "/") && !strings.HasPrefix(parsedSignatureDependency.Path, "./") && !strings.HasPrefix(parsedSignatureDependency.Path, "../") {
//		parsedSignatureDependency.Path = "./" + parsedSignatureDependency.Path
//	}
//
//	if !strings.HasPrefix(functionPath, "/") && !strings.HasPrefix(functionPath, "./") && !strings.HasPrefix(functionPath, "../") {
//		functionPath = "./" + functionPath
//	}
//
//	buf := new(bytes.Buffer)
//	err := g.template.ExecuteTemplate(buf, "packagejson.ts.templ", map[string]interface{}{
//		"version":              packageVersion,
//		"package":              packageName,
//		"dependencies":         dependencies,
//		"signature_dependency": "signature",
//		"signature_package":    parsedSignatureDependency.Package,
//		"signature_path":       parsedSignatureDependency.Path,
//		"signature_version":    parsedSignatureDependency.Version,
//		"function_dependency":  schema.Name,
//		"function_path":        functionPath,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return buf.Bytes(), nil
//}

//
// func (g *Generator) GenerateRustLib(signature *signature.Schema, schema *scalefile.Schema, scaleVersion string) ([]byte, error) {
//	buf := new(bytes.Buffer)
//	err := g.template.ExecuteTemplate(buf, "lib.rs.templ", map[string]interface{}{
//		"version":   scaleVersion,
//		"signature": signature,
//		"schema":    schema,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return buf.Bytes(), nil
//}
