package parser

import (
	"fmt"
	"go/ast"
	"strings"
)

func exprToString(e ast.Expr) string {
	if e == nil {
		return ""
	}
	switch x := e.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.StarExpr:
		return "*" + exprToString(x.X)
	case *ast.SelectorExpr:
		return exprToString(x.X) + "." + x.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToString(x.Elt)
	case *ast.MapType:
		return "map[" + exprToString(x.Key) + "]" + exprToString(x.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct"
	case *ast.FuncType:
		return "func"
	case *ast.BasicLit:
		return x.Value
	case *ast.CallExpr:
		return exprToString(x.Fun)
	case *ast.CompositeLit:
		return exprToString(x.Type)
	case *ast.IndexExpr:
		return exprToString(x.X) + "[" + exprToString(x.Index) + "]"
	default:
		return fmt.Sprintf("%T", e)
	}
}

func extractTypeName(expr ast.Expr) string {
	switch a := expr.(type) {
	case *ast.UnaryExpr:
		return exprToString(a.X)
	case *ast.CompositeLit:
		return exprToString(a.Type)
	case *ast.Ident:
		return a.Name
	case *ast.StarExpr:
		return exprToString(a.X)
	case *ast.CallExpr:
		return exprToString(a.Fun)
	default:
		return ""
	}
}

func extractPathParams(path string) []string {
	var out []string
	parts := strings.Split(path, "/")
	for _, p := range parts {
		if strings.HasPrefix(p, ":") {
			out = append(out, strings.TrimPrefix(p, ":"))
		}
	}
	return out
}

func inspectHandlerBody(fd *ast.FuncDecl, structs map[string]StructDef) (string, string) {
	var bt, rt string
	ast.Inspect(fd.Body, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if se, ok := call.Fun.(*ast.SelectorExpr); ok {
			sel := se.Sel.Name
			if sel == "BodyParser" && len(call.Args) >= 1 {
				if t := extractTypeName(call.Args[0]); t != "" {
					bt = t
					if _, ok := structs[t]; !ok {
						// nothing
					}
				}
			}
			if sel == "JSON" && len(call.Args) >= 1 {
				if t := extractTypeName(call.Args[0]); t != "" {
					rt = t
				}
			}
		}
		return true
	})
	return bt, rt
}

func isHTTPMethod(m string) bool {
	switch m {
	case "get", "post", "put", "delete", "patch", "options", "head":
		return true
	default:
		return false
	}
}

func parseTag(raw, key string) string {
	st := raw
	for st != "" {
		st = strings.TrimLeft(st, " ")
		if !strings.HasPrefix(st, key+`:"`) {
			idx := strings.Index(st, " ")
			if idx < 0 {
				break
			}
			st = st[idx+1:]
			continue
		}
		st = st[len(key)+2:]
		end := strings.Index(st, `"`)
		if end < 0 {
			return ""
		}
		val := st[:end]
		return val
	}
	return ""
}
