package main

import (
	"github.com/urfave/cli/v2"
)

func commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "yml",
			Usage:  "📄 Render tpl -> yml",
			Action: app.CliYml,
		},
		{
			Name:    "planfile",
			Aliases: []string{"plan"},
			Usage:   "📜 Generate planfile to plandir",
			Action:  app.CliPlan,
		},
		//{
		//	Name:    "deploy",
		//	Aliases: []string{"apply", "sync", "release"},
		//	Usage:   "🛥 Deploy your helmwave!",
		//	Action:  app.C,
		//},
	}

}

func help(c *cli.Context) error {
	args := c.Args()
	if args.Present() {
		return cli.ShowCommandHelp(c, args.First())
	}

	return cli.ShowAppHelp(c)
}
