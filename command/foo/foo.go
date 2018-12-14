package foo

import (
	"github.com/rjeczalik/cobra-example/command"

	"github.com/spf13/cobra"
)

func NewCommand(app *command.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "foo",
		Short: "Foo resource",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewLsCommand(app),
		NewAddCommand(app),
		NewDelCommand(app),
	)

	return cmd
}
