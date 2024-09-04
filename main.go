package main

import (
	"os"

	"github.com/stevancvetkovic/go-grpc-addressbook/pkg/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name: "addressbook",
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "serve the agenda server on port 50051",
				Action: serve,
				Flags:  []cli.Flag{},
			},
		},
	}

	app.Run(os.Args)
}

func serve(c *cli.Context) error {
	srv, err := server.New()
	if err != nil {
		return cli.Exit(err, 1)
	}

	if err = srv.Serve(":50051"); err != nil {
		return cli.Exit(err, 1)
	}
	return nil
}
