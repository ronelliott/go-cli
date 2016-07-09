package cli

// Run the defined app
func run(cfg *AppConfig, withHelp bool) error {
	app, err := cfg.Convert(withHelp)

	if err != nil {
		return err
	}

	err = app.CategoryCommand.Command.Run(nil)

	if err != nil {
		return err
	}

	return nil
}

// Run the defined app without adding a help command
func Run(cfg *AppConfig) error {
	return run(cfg, false)
}

// Run the defined app and adds a help command
func RunWithHelp(cfg *AppConfig) error {
	return run(cfg, true)
}
