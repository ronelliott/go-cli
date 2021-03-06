package cli

import (
	"errors"
	"io"
	"os"
	"path"
	"strings"

	"github.com/ronelliott/go-opts"
)

type Runner interface {
	Run() error
}

type Command struct {
	// the function to run when calling this command
	Callback Runner

	// the (short) description of the command
	Description string

	// the (long) description of the command
	Help string

	// the name of the command
	Name string

	// the OptionSet for the command
	Options *opts.OptionSet

	// the commands subcommands
	Subs map[string]*Command
}

// Returns the first string not prefixed with a tack ('-')
func getCommandName(args []string) (string, int) {
	for idx, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			return arg, idx
		}
	}

	return "", -1
}

// Create a new root level command.
func NewCommand(
	name,
	description,
	help string,
	callback Runner) (*Command, error) {
	var set *opts.OptionSet
	var err error

	if callback != nil {
		set, err = opts.NewOptionSet(callback)

		if err != nil {
			return nil, err
		}
	}

	return &Command{
		Callback:    callback,
		Description: description,
		Help:        help,
		Name:        name,
		Options:     set,
		Subs:        map[string]*Command{},
	}, nil
}

// Create a new root level command.
func New(description string, callback Runner) (*Command, error) {
	return NewCommand(path.Base(os.Args[0]), description, "", callback)
}

// Checks if the command has a callback defined
func (this *Command) HasCallback() bool {
	return this.Callback != nil
}

// Checks if the command has subcommands.
func (this *Command) HasSubs() bool {
	return len(this.Subs) != 0
}

// Checks if the command has a subcommand with the given name.
func (this *Command) HasSub(name string) bool {
	_, ok := this.Subs[name]
	return ok
}

// Create a new sub level command.
func (this *Command) NewSub(
	name,
	description string,
	help string,
	callback Runner) (*Command, error) {
	cmd, err := NewCommand(name, description, help, callback)

	if err != nil {
		return nil, err
	}

	this.NewSubCommand(cmd)
	return cmd, nil
}

// Create a new sub level command.
func (this *Command) NewSubCommand(cmd *Command) {
	this.Subs[cmd.Name] = cmd
}

// Create new sub level commands from the given list.
func (this *Command) NewSubCommands(cmds []*Command) error {
	for _, sub := range cmds {
		this.NewSubCommand(sub)
	}

	return nil
}

// Checks if the command has options
func (this *Command) HasOptions() bool {
	return this.Options.HasOptions()
}

// Checks if the command has positional args
func (this *Command) HasPositional() bool {
	return this.Options.HasPositional()
}

// Run the command
func (this *Command) Run(args []string) error {
	if args == nil {
		args = os.Args[1:]
	}

	if this.HasSubs() {
		name, idx := getCommandName(args)

		if sub, ok := this.Subs[name]; ok {
			return sub.Run(args[idx+1:])
		}
	}

	if !this.HasCallback() {
		return errors.New("No command callback defined!")
	}

	err := this.Options.Parse(args)

	if err != nil {
		return err
	}

	return this.Callback.Run()
}

// Writes the default options and descriptions to the given io.Writer
func (this *Command) WriteHelp(out io.Writer) {
	this.Options.WriteHelp(out)
}
