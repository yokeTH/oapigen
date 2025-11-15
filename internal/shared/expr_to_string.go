package shared

import (
	"go/ast"
)

func ExprToString(e ast.Expr) string {
	if e == nil {
		return ""
	}
	switch x := e.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.StarExpr:
		return "*" + ExprToString(x.X)
	case *ast.SelectorExpr:
		return ExprToString(x.X) + "." + x.Sel.Name
	case *ast.ArrayType:
		return "[]" + ExprToString(x.Elt)
	case *ast.MapType:
		return "map[" + ExprToString(x.Key) + "]" + ExprToString(x.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct"
	case *ast.FuncType:
		return "func"
	case *ast.BasicLit:
		return x.Value
	case *ast.CallExpr:
		return ExprToString(x.Fun)
	case *ast.CompositeLit:
		return ExprToString(x.Type)
	case *ast.IndexExpr:
		return ExprToString(x.X) + "[" + ExprToString(x.Index) + "]"
	default:
		return ""
	}
}
