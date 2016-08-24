package cli

type CategoryConfig struct {
	// the name of the CategoryCommand
	Name string

	// the description of the CategoryCommand
	Desc string

	// the help for the CategoryCommand
	Help string

	// the default command for the CategoryCommand
	Default Runner

	// the subcommands for the CategoryCommand
	Commands []CommandConverter
}

func (cfg *CategoryConfig) Convert() (*Command, error) {
	cmd, err := NewCategoryCommandWithDefault(
		cfg.Name, cfg.Desc, cfg.Help, cfg.Default)

	if err != nil {
		return nil, err
	}

	for _, sub := range cfg.Commands {
		subCmd, err := sub.Convert()

		if err != nil {
			return nil, err
		}

		cmd.NewSubCommand(subCmd)
	}

	return &cmd.Command, nil
}
