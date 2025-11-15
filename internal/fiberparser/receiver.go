package fiberparser

import (
	"go/ast"

	"github.com/rs/zerolog/log"
	"github.com/yokeTH/oapigen/internal/shared"
)

func inspectReceiver(fd *ast.FuncDecl) (string, string) {
	log.Debug().Str("func", fd.Name.Name).Msg("Processing")
	var bt, rt string

	ast.Inspect(fd.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		chain := shared.GetCallChain(call)
		var lastCall, beforeLast string
		if len(chain) > 1 {
			lastCall = chain[len(chain)-1]
		}
		if len(chain) > 2 {
			beforeLast = chain[len(chain)-2]
		}

		switch lastCall {
		case "JSON":
			if beforeLast == "Bind" { // ctx.Bind().JSON() -> req
				bt = shared.ResolveArgType(call.Args[0], fd)
			} else { // ctx.JSON() -> res
				rt = shared.ResolveArgType(call.Args[0], fd)
			}

		case "Body":
			bt = shared.ResolveArgType(call.Args[0], fd)
		}

		return true
	})
	return bt, rt
}
