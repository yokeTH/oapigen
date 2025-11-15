package fiberparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/yokeTH/oapigen/internal/shared"
)

func Parse(root string) error {
	fset := token.NewFileSet()
	files := make(map[string]*ast.File)

	// walk to every .go files and store to files as filename: *ast.File
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

	// find all struct to be generate as open api component
	structs := shared.CollectStruct(files)
	for _, sd := range structs {
		log.Debug().Str("name", sd.Name).Any("field", sd.Fields).Msg("Found Struct")
	}

	routes := findRoute(files, structs)
	for _, route := range routes {
		log.Debug().
			Str("body", route.BodyType).
			Str("response", route.RespType).
			Str("handler", route.Handler).
			Str("path", route.Path).
			Str("method", route.Method).
			Msg("Found route")
	}

	return nil
}
