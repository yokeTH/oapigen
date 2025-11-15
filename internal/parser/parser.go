package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func Parser(root string) error {
	fset := token.NewFileSet()
	files := make(map[string]*ast.File)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}
		f, e := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if e != nil {
			return e
		}
		files[path] = f
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	appNames := FindFiberAppNames(files)
	structs := collectStructs(files)
	routes := findRoutes(files, appNames, structs)

	fmt.Printf("appNames: %v\n", appNames)
	fmt.Printf("structs: %v\n", structs)
	fmt.Printf("routes: %v\n", routes)

	return nil
}
