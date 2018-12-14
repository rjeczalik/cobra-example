package command

import (
	mw "github.com/rjeczalik/cobra-example/command/middleware"

	"io/ioutil"
)

func (app *App) ReadConfig(next mw.CobraFunc) mw.CobraFunc {
	if app.ConfigFile == "" {
		return mw.Errorf("missing config file; provide with --config|-c flag")
	}
	var err error
	app.Config, err = ioutil.ReadFile(app.ConfigFile)
	if err != nil {
		return mw.Errorf("error reading config file: %s", err)
	}
	return next
}
