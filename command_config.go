package cli

type CommandConfig struct {
	// the name of the Command
	Name string

	// the description of the Command
	Desc string

	// the help for the Command
	Help string

	// the Runner for the Command
	Run Runner
}

func (cfg *CommandConfig) Convert() (*Command, error) {
	return NewCommand(cfg.Name, cfg.Desc, cfg.Help, cfg.Run)
}
