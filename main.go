package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
	"github.com/yokeTH/oapigen/internal/fiberparser"
)

func main() {
	cmd := &cli.Command{
		Name:  "oapigen",
		Usage: "scan a Fiber v3 project and emit routes + structs as JSON",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "debug", Aliases: []string{"d"}, Value: false, Usage: "Log the debug level"},
			&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Value: ".", Usage: "project root to scan"},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			console := zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			}
			debug := cmd.Bool("debug")
			log.Logger = log.Output(console).With().Timestamp().Logger()
			if debug {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			} else {
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
			}
			return ctx, nil
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			root := cmd.String("path")

			return fiberparser.Parse(root)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err)
	}
}
