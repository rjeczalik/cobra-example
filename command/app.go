package command

import (
	"context"
	"io/ioutil"

	"github.com/spf13/pflag"
)

type App struct {
	context.Context

	// Global flags.
	ConfigFile string // -c|--config
	Format     string // -f|--format

	// Common resources.
	Config []byte
}

func (app *App) Register(flags *pflag.FlagSet) {
	flags.StringVarP(&app.ConfigFile, "config", "c", "", "Configuration file")
	flags.StringVarP(&app.Format, "format", "f", "json", "Format type of output")
}

func (app *App) ReadConfig(next CobraFunc) CobraFunc {
	if app.ConfigFile == "" {
		return Errorf("missing config file; provide with --config|-c flag")
	}
	var err error
	app.Config, err = ioutil.ReadFile(app.ConfigFile)
	if err != nil {
		return Errorf("error reading config file: %s", err)
	}
	return next
}
