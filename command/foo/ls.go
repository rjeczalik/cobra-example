package foo

import (
	"errors"

	"github.com/rjeczalik/cobra-example/command"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewLsCommand(app *command.App) *cobra.Command {
	c := &lsCmd{App: app}

	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List foo items",
		RunE:  c.run,
	}

	c.register(cmd.Flags())

	return cmd
}

type lsCmd struct {
	*command.App
	unique bool // --unique|-u
	limit  int  // --limit|-n
}

func (c *lsCmd) register(flags *pflag.FlagSet) {
	flags.BoolVarP(&c.unique, "unique", "u", false, "Skip duplicates")
	flags.IntVarP(&c.limit, "limit", "n", 100, "Show at most")
}

func (c *lsCmd) run(*cobra.Command, []string) error {
	return errors.New("no foo items found")
}
