package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/yokeTH/oapigen/internal/parser"
)

func main() {
	cmd := &cli.Command{
		Name:  "oapigen",
		Usage: "scan a Fiber v3 project and emit routes + structs as JSON",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Value: ".", Usage: "project root to scan"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			root := cmd.String("path")
			parser.Parser(root)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
