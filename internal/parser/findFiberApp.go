package parser

import (
	"go/ast"
	"strings"
)

func FindFiberAppNames(files map[string]*ast.File) map[string]struct{} {
	apps := map[string]struct{}{}
	for _, f := range files {
		ast.Inspect(f, func(n ast.Node) bool {
			as, ok := n.(*ast.AssignStmt)
			if !ok {
				return true
			}
			for i, rhs := range as.Rhs {
				call, ok := rhs.(*ast.CallExpr)
				if !ok {
					continue
				}
				switch fun := call.Fun.(type) {
				case *ast.SelectorExpr:
					x := exprToString(fun.X)
					sel := fun.Sel.Name
					if strings.HasSuffix(x, "fiber") && sel == "New" {
						if len(as.Lhs) > i {
							if id, ok := as.Lhs[i].(*ast.Ident); ok {
								apps[id.Name] = struct{}{}
							}
						}
					}
				case *ast.Ident:
					if fun.Name == "New" {
						if len(as.Lhs) > i {
							if id, ok := as.Lhs[i].(*ast.Ident); ok {
								apps[id.Name] = struct{}{}
							}
						}
					}
				}
			}
			return true
		})
	}
	return apps
}
