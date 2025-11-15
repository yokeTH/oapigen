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
			case *ast.Ident:
				// function declaretion func Handler(ctx fiber.Ctx) error
				panic("unimplement")
			case *ast.SelectorExpr:
				// reciver method (e.g. bookHandler.CreateBook)
				var receiver string
				if ident, ok := h.X.(*ast.Ident); ok {
					receiver = ident.Name
				} else {
					receiver = shared.ExprToString(h.X)
				}
				handlerName = receiver + "." + h.Sel.Name
				// method declarations might be stored with receiver prefixed by "*"
				if fd, ok := funcDecls["*"+handlerName]; ok {
					bt, rt := inspectReceiver(fd)
					bodyType = bt
					respType = rt
				} else if fd, ok := funcDecls[handlerName]; ok {
					bt, rt := inspectReceiver(fd)
					bodyType = bt
					respType = rt
				}

			case *ast.FuncLit:
				// anonymous function (e.g. app.Get("/path", func (c fiber.Ctx) err {}))
				panic("unimplemented")
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
