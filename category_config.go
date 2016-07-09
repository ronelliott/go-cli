package cli

type CategoryConfig struct {
	// the name of the CategoryCommand
	Name string

	// the description of the CategoryCommand
	Desc string

	// the help for the CategoryCommand
	Help string

	// the subcommands for the CategoryCommand
	Commands []CommandConverter
}

func (cfg *CategoryConfig) Convert() (*Command, error) {
	cmd, err := NewCategoryCommand(cfg.Name, cfg.Desc, cfg.Help)

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
