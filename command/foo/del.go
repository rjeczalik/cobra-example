package foo

import (
	"fmt"

	"github.com/rjeczalik/cobra-example/command"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewDelCommand(app *command.App) *cobra.Command {
	c := &delCmd{App: app}

	cmd := &cobra.Command{
		Use:   "del",
		Short: "Delete foo items",
		RunE:  c.run,
	}

	c.register(cmd.Flags())

	return cmd
}

type delCmd struct {
	*command.App
	force bool // --force
}

func (c *delCmd) register(flags *pflag.FlagSet) {
	flags.BoolVar(&c.force, "force", false, "Force deletion")
}

func (c *delCmd) run(_ *cobra.Command, args []string) error {
	for _, arg := range args {
		fmt.Println(arg, "was deleted")
	}
	return nil
}
