package bar

import (
	"github.com/rjeczalik/cobra-example/command"

	"github.com/spf13/cobra"
)

func NewCommand(app *command.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bar",
		Short: "Bar service",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewServeCommand(app),
	)

	return cmd
}
