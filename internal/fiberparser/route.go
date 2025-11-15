package fiberparser

import (
	"go/ast"

	"github.com/rs/zerolog/log"
	"github.com/yokeTH/oapigen/internal/shared"
)

func findRoute(files map[string]*ast.File, structs map[string]shared.StructDef) []shared.Route {
	var route []shared.Route

	// Collect all function/method declarations
	funcDecls := map[string]*ast.FuncDecl{}
	for _, f := range files {
		for _, d := range f.Decls {
			if fd, ok := d.(*ast.FuncDecl); ok && fd.Name != nil {
				key := fd.Name.Name
				if fd.Recv != nil && len(fd.Recv.List) > 0 {
					if starExpr, ok := fd.Recv.List[0].Type.(*ast.StarExpr); ok {
						key = "*" + shared.ExprToString(starExpr.X) + "." + fd.Name.Name
					} else {
						key = shared.ExprToString(fd.Recv.List[0].Type) + "." + fd.Name.Name
					}
				}
				funcDecls[key] = fd
			}
		}
	}

	for f, fd := range funcDecls {
		_ = fd
		log.Trace().Str("func", f).Msg("Found")
	}

	return route
}
