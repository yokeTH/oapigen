package fiberparser

import (
	"go/ast"
	"go/token"
	"strings"

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
		log.Debug().Str("func", f).Msg("Found")
	}

	for fname, f := range files {
		log.Debug().Str("filename", fname).Msg("Processing")

		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			method := strings.ToLower(sel.Sel.Name)
			if !shared.IsHTTPMethod(method) {
				return true
			}
			if len(call.Args) < 2 {
				return true
			}
			pathLit, ok := call.Args[0].(*ast.BasicLit)
			if !ok || pathLit.Kind != token.STRING {
				return true
			}

			path := strings.Trim(pathLit.Value, `"`)
			last := call.Args[len(call.Args)-1]
			handlerName := ""
			bodyType := ""
			respType := ""
			switch h := last.(type) {
			case *ast.SelectorExpr:
				var receiver string
				if ident, ok := h.X.(*ast.Ident); ok {
					receiver = ident.Name
				} else {
					receiver = shared.ExtractTypeName(h.X)
				}
				handlerName = receiver + "." + h.Sel.Name
				// method declarations might be stored with receiver prefixed by "*"
				if fd, ok := funcDecls["*"+handlerName]; ok {
					bt, rt := inspectHandlerBody(fd, structs)
					bodyType = bt
					respType = rt
				} else if fd, ok := funcDecls[handlerName]; ok {
					bt, rt := inspectHandlerBody(fd, structs)
					bodyType = bt
					respType = rt
				}

			case *ast.FuncLit:
				handlerName = "anonymous"
				bt, rt := inspectFuncLit(h, structs)
				bodyType = bt
				respType = rt
			}
			params := shared.ExtractColonPathParam(path)

			route = append(route, shared.Route{
				Method:   method,
				Path:     path,
				Handler:  handlerName,
				BodyType: bodyType,
				RespType: respType,
				Params:   params,
			})
			return true
		})
	}

	return route
}

func inspectFuncLit(h *ast.FuncLit, structs map[string]shared.StructDef) (string, string) {
	panic("unimplemented")
}

func inspectHandlerBody(fd *ast.FuncDecl, structs map[string]shared.StructDef) (string, string) {
	log.Debug().Str("func", fd.Name.Name).Msg("Processing")
	var bt, rt string

	ast.Inspect(fd.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		name := sel.Sel.Name

		switch name {
		case "JSON", "Body", "All":
			if recvCall, ok := sel.X.(*ast.CallExpr); ok {
				if recvSel, ok := recvCall.Fun.(*ast.SelectorExpr); ok {
					recvName := recvSel.Sel.Name

					if recvName == "Bind" && len(call.Args) > 0 { // Ctx.Bind().XXX(XXX)
						bt = shared.ResolveArgType(call.Args[0], fd)

					} else if recvName == "Status" && len(call.Args) > 0 { // Ctx.Status().JSON(XXX)
						rt = shared.ResolveArgType(call.Args[0], fd)
					}
				}
			} else {
				// ctx.JSON(...)
				if len(call.Args) > 0 {
					rt = shared.ResolveArgType(call.Args[0], fd)
				}
			}
		}
		return true

	})
	return bt, rt
}
