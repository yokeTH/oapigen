package shared

import "go/ast"

func ResolveArgType(expr ast.Expr, fd *ast.FuncDecl) string {
	switch a := expr.(type) {
	case *ast.UnaryExpr: // &req
		if ident, ok := a.X.(*ast.Ident); ok {
			return lookupVarType(fd, ident.Name)
		}
	case *ast.Ident: // resp
		return lookupVarType(fd, a.Name)
	case *ast.CallExpr: // Success(resp)
		if len(a.Args) > 0 {
			return ResolveArgType(a.Args[0], fd)
		}
	}
	return ""
}
