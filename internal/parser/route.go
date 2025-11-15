package parser

import (
	"go/ast"
	"go/token"
	"strings"
)

type Route struct {
	Method   string   `json:"method"`
	Path     string   `json:"path"`
	Handler  string   `json:"handler"`
	BodyType string   `json:"body,omitempty"`
	RespType string   `json:"response,omitempty"`
	Params   []string `json:"params,omitempty"`
}

func findRoutes(files map[string]*ast.File, appNames map[string]struct{}, structs map[string]StructDef) []Route {
	var out []Route
	funcDecls := map[string]*ast.FuncDecl{}
	for _, f := range files {
		for _, d := range f.Decls {
			if fd, ok := d.(*ast.FuncDecl); ok && fd.Name != nil {
				funcDecls[fd.Name.Name] = fd
			}
		}
	}
	for _, f := range files {
		ast.Inspect(f, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			x := exprToString(sel.X)
			if _, ok := appNames[x]; !ok {
				return true
			}
			method := strings.ToLower(sel.Sel.Name)
			if !isHTTPMethod(method) {
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
				handlerName = h.Name
				if fd, ok := funcDecls[handlerName]; ok {
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
			params := extractPathParams(path)
			out = append(out, Route{Method: method, Path: path, Handler: handlerName, BodyType: bodyType, RespType: respType, Params: params})
			return true
		})
	}
	return out
}

func inspectFuncLit(fl *ast.FuncLit, structs map[string]StructDef) (string, string) {
	var bt, rt string
	ast.Inspect(fl.Body, func(n ast.Node) bool {
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
