package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rescDev/actions-docs/internal/markdown"
	"github.com/rescDev/actions-docs/internal/update"
	"github.com/urfave/cli/v2"
)

const (
	author = "Rene Schach - https://github.com/rescDev"
)

var (
	name  = "actions-docs"
	usage = "Generate documentation for your GitHub Action"

	version   string
	osarch    string
	buildTime string
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s - version %s - %s - built: %s\n", c.App.Name, c.App.Version, osarch, buildTime)
	}

	app := &cli.App{
		Name:    name,
		Usage:   usage,
		Version: version,
		Authors: []*cli.Author{
			{
				Name: author,
			},
		},
		EnableBashCompletion: true,
		HideHelpCommand:      true,
		Commands: []*cli.Command{
			markdown.MarkdownCommand(),
			update.UpdateCommand(),
		},
		UsageText: "actions-docs [global options] command [command options]",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
