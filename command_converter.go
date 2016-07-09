package cli

type CommandConverter interface {
	Convert() (*Command, error)
}
