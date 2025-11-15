package shared

import "go/ast"

func GetCallChain(expr ast.Expr) []string {
	var parts []string

	// helper to unwind wrapper expressions
	unwrap := func(e ast.Expr) ast.Expr {
		for {
			switch v := e.(type) {
			case *ast.ParenExpr:
				e = v.X
			case *ast.StarExpr:
				e = v.X
			case *ast.TypeAssertExpr:
				e = v.X
			case *ast.IndexExpr:
				e = v.X
			default:
				return e
			}
		}
	}

	// We'll walk from the outermost expression inward.
	e := unwrap(expr)
	for {
		switch n := e.(type) {
		case *ast.CallExpr:
			// A CallExpr wraps a Fun: continue inspecting the Fun
			e = unwrap(n.Fun)
			continue

		case *ast.SelectorExpr:
			// record method/field name
			parts = append(parts, n.Sel.Name)
			// continue into the receiver X
			e = unwrap(n.X)
			continue

		case *ast.Ident:
			parts = append(parts, n.Name)
			// we reached the base; reverse parts (currently [lastMethod,...,base])
			// but we appended in reverse order: JSON, Bind, ctx -> we want ctx, Bind, JSON
			reverse(parts)
			return parts

		default:
			// unknown base (could be a composite literal, call to func(), index, literal...)
			// try to get a printable fallback (e.g., for "myPkg.Some().Method()")
			// If the base is another SelectorExpr (e.g. pkg.Var), let the selector case above handle it.
			// Otherwise reverse whatever we collected and return.
			reverse(parts)
			return parts
		}
	}
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
