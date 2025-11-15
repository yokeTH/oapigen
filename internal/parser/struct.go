package parser

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

type StructField struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	SchemaName string `json:"json,omitempty"`
	Omit       bool   `json:"omit,omitempty"`
}

type StructDef struct {
	Name   string        `json:"name"`
	Fields []StructField `json:"fields"`
}

func collectStructs(files map[string]*ast.File) map[string]StructDef {
	out := map[string]StructDef{}
	for _, f := range files {
		for _, d := range f.Decls {
			ts, ok := d.(*ast.GenDecl)
			if !ok || ts.Tok != token.TYPE {
				continue
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
				sd := StructDef{Name: tsp.Name.Name}
				for _, fld := range st.Fields.List {
					if len(fld.Names) == 0 {
						continue
					}
					name := fld.Names[0].Name
					typ := exprToString(fld.Type)
					schemaName := name
					omit := false
					if fld.Tag != nil {
						if unq, err := strconv.Unquote(fld.Tag.Value); err == nil {
							tag := parseTag(unq, "json")
							if tag != "" {
								parts := strings.Split(tag, ",")
								if parts[0] == "-" {
									omit = true
								} else if parts[0] != "" {
									schemaName = parts[0]
								}
							}
							if t := parseTag(unq, "oapi"); t != "" {
								if t == "-" {
									omit = true
								} else {
									if strings.Contains(t, "=") {
										kv := strings.SplitN(t, "=", 2)
										if kv[1] != "" {
											schemaName = kv[1]
										}
									} else if t != "" {
										schemaName = t
									}
								}
							}
						}
					}
					if omit {
						continue
					}
					sd.Fields = append(sd.Fields, StructField{Name: name, Type: typ, SchemaName: schemaName})
				}
				out[sd.Name] = sd
			}
		}
	}
	return out
}
