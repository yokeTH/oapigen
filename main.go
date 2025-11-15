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
			&cli.Int8Flag{Name: "log-level", Aliases: []string{"l"}, Value: 1, Usage: "set zerolog global log level (Panic=5, Fatal=4, Error=3, Warn=2, Info=1, Debug=0, Trace=-1)"},
			&cli.StringFlag{Name: "path", Aliases: []string{"p"}, Value: ".", Usage: "project root to scan"},
		},
		Before: func(ctx context.Context, cmd *cli.Command) (context.Context, error) {
			console := zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			}
			logLv := cmd.Int8("log-level")
			log.Logger = log.Output(console).With().Timestamp().Logger()
			zerolog.SetGlobalLevel(zerolog.Level(logLv))
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
