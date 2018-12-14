package foo

import (
	"fmt"

	"github.com/rjeczalik/cobra-example/command"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewAddCommand(app *command.App) *cobra.Command {
	c := &addCmd{App: app}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add foo items",
		RunE:  c.run,
	}

	c.register(cmd.Flags())

	return cmd
}

type addCmd struct {
	*command.App
	force bool // --force
}

func (c *addCmd) register(flags *pflag.FlagSet) {
	flags.BoolVar(&c.force, "force", false, "Force addition")
}

func (c *addCmd) run(_ *cobra.Command, args []string) error {
	for _, arg := range args {
		fmt.Println(arg, "was added")
	}
	return nil
}
