package shared

import "go/ast"

func ExtractTypeName(expr ast.Expr) string {
	switch a := expr.(type) {
	case *ast.UnaryExpr:
		return ExprToString(a.X)
	case *ast.CompositeLit:
		return ExprToString(a.Type)
	case *ast.Ident:
		return a.Name
	case *ast.StarExpr:
		return ExprToString(a.X)
	case *ast.CallExpr:
		return ExprToString(a.Fun)
	default:
		return ""
	}
}
