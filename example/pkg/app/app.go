package app

import (
	"github.com/rancher/wrangler-cli"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	root := cli.Command(&App{}, cobra.Command{
		Long: "Add some long description",
	})
	root.AddCommand(
		NewSubCommand(),
	)
	return root
}

type App struct {
	OptionOne string `usage:"Some usage description"`
	OptionTwo string `name:"custom-name"`
}

func (a *App) Run(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}