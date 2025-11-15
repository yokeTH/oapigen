package shared

import (
	"go/ast"
	"go/token"
)

func lookupVarType(fd *ast.FuncDecl, varName string) string {
	var typ string

	ast.Inspect(fd.Body, func(n ast.Node) bool {
		if n == nil {
			return true
		}

		// var declarations: var req CreateBookRequest
		if declStmt, ok := n.(*ast.DeclStmt); ok {
			if genDecl, ok := declStmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
				for _, spec := range genDecl.Specs {
					if valSpec, ok := spec.(*ast.ValueSpec); ok {
						for i, name := range valSpec.Names {
							if name.Name == varName {
								if valSpec.Type != nil {
									typ = ExprToString(valSpec.Type)
								} else if len(valSpec.Values) > i {
									typ = ExprToString(valSpec.Values[i])
								}
								return false
							}
						}
					}
				}
			}
		}

		// short variable declarations: req := SomeType{}
		if assignStmt, ok := n.(*ast.AssignStmt); ok && assignStmt.Tok == token.DEFINE {
			for i, lhs := range assignStmt.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok && ident.Name == varName {
					if i < len(assignStmt.Rhs) {
						typ = ExprToString(assignStmt.Rhs[i])
						return false
					}
				}
			}
		}

		return true
	})

	return typ
}
