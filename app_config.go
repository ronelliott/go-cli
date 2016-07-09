package cli

type AppConfig struct {
	// the name of the App
	Name string

	// the description of the App
	Desc string

	// the help for the App
	Help string

	// the subcommands for the App
	Commands []CommandConverter
}

func (cfg *AppConfig) Convert(withHelp bool) (*App, error) {
	app, err := NewApp(cfg.Name, cfg.Desc, cfg.Help)

	if err != nil {
		return nil, err
	}

	for _, sub := range cfg.Commands {
		subCmd, err := sub.Convert()

		if err != nil {
			return nil, err
		}

		app.NewSubCommand(subCmd)
	}

	if withHelp {
		err = app.AddHelpCommand()

		if err != nil {
			return nil, err
		}
	}

	return app, nil
}
