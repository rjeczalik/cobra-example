package bar

import (
	"log"

	"github.com/rjeczalik/cobra-example/command"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewServeCommand(app *command.App) *cobra.Command {
	c := &barCmd{App: app}

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run Bar server",
		Args:  cobra.NoArgs,
		RunE:  c.run,
	}

	c.register(cmd.Flags())
	command.Use(cmd, app.ReadConfig)

	return cmd
}

type barCmd struct {
	*command.App
	http string
}

func (c *barCmd) register(flags *pflag.FlagSet) {
	flags.StringVar(&c.http, "http", "localhost:3000", "Network address to listen on")
}

func (c *barCmd) run(*cobra.Command, []string) error {
	log.Printf("Serving on %s ...", c.http)
	<-c.Done()
	log.Print("Serving done")
	return nil
}
