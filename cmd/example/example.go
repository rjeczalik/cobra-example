package main

import (
	"github.com/rjeczalik/cobra-example/command"
	"github.com/rjeczalik/cobra-example/command/bar"
	"github.com/rjeczalik/cobra-example/command/foo"

	"github.com/spf13/cobra"
)

func NewCommand(app *command.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "Example cobra project",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		foo.NewCommand(app),
		bar.NewCommand(app),
	)

	app.Register(cmd.PersistentFlags())

	return cmd
}
