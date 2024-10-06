package main

import (
	"github.com/go-zoox/cli"
	httpproxy "github.com/go-zoox/http-proxy"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "http-proxy",
		Usage:   "cli for http-proxy server",
		Version: httpproxy.Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Usage:   "the port for server",
				Value:   8080,
			},
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				EnvVars: []string{"USERNAME"},
				Usage:   "the username for server",
			},
			&cli.StringFlag{
				Name: "password",
				// Aliases: []string{"u"},
				EnvVars: []string{"PASSWORD"},
				Usage:   "the password for server",
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		server := httpproxy.New(func(cfg *httpproxy.Config) {
			cfg.Port = ctx.Int("port")
			//
			cfg.Username = ctx.String("username")
			cfg.Password = ctx.String("password")
		})

		return server.Run()
	})

	app.Run()
}
