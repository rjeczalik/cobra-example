package command

import (
	"context"

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
