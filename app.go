package cli

type App struct {
	CategoryCommand
}

func NewApp(name, description, help string) (*App, error) {
	cmd, err := NewCategoryCommand(name, description, help)

	if err != nil {
		return nil, err
	}

	app := App{*cmd}
	return &app, nil
}

func NewAppWithDefault(name, description, help string, def Runner) (*App, error) {
	cmd, err := NewCategoryCommandWithDefault(name, description, help, def)

	if err != nil {
		return nil, err
	}

	app := App{*cmd}
	return &app, nil
}

func (this *App) AddHelpCommand() error {
	_, err := this.NewSub(
		"help",
		"Display help for a specific command.",
		"See help. :P",
		&HelpCommand{
			app: this,
		})
	return err
}
