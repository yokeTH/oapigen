package shared

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

type StructDef struct {
	Name   string
	Fields []StructField
}

type StructField struct {
	Name       string
	Type       string
	SchemaName string
	Omit       bool
}

func CollectStruct(files map[string]*ast.File) map[string]StructDef {
	result := map[string]StructDef{}

	for _, file := range files {
		for _, decl := range file.Decls {
			// GenDecl is General declaration (import, type, var, const).
			ts, ok := decl.(*ast.GenDecl)
			if !ok || ts.Tok != token.TYPE {
				continue // Skip if it is not type
			}

			for _, spec := range ts.Specs {
				tsp, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				st, ok := tsp.Type.(*ast.StructType)
				if !ok {
					continue
				}

				// Struct definition
				sd := StructDef{Name: tsp.Name.Name}

				// Loop for struct fields
				for _, field := range st.Fields.List {
					if len(field.Names) == 0 {
						continue
					}

					name := field.Names[0].Name
					typ := ExprToString(field.Type)
					schemaName := name
					omit := false

					if field.Tag != nil {
						if unq, err := strconv.Unquote(field.Tag.Value); err == nil {
							jTag := ParseTag(unq, "json")
							if jTag != "" {
								parts := strings.Split(jTag, ",")
								if parts[0] == "-" {
									omit = true
								} else if parts[0] != "" {
									schemaName = parts[0]
								}
							}
						}
					}

					if omit {
						continue
					}

					sd.Fields = append(sd.Fields, StructField{Name: name, Type: typ, SchemaName: schemaName})
				}
				result[sd.Name] = sd
			}
		}
	}

	return result
}
