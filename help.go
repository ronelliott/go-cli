package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type HelpCommand struct {
	// The App this command is attached to
	app  *App

	// User provided positional args
	Args []string `positional:"true"`
}

func (this *HelpCommand) GetCommand(
		names []string,
		cur *Command) (*Command, error) {

	if len(names) == 0 {
		return nil, errors.New("No names given.")
	}

	name := ""
	name, names = names[0], names[1:len(names)]
	cmd, valid := cur.Subs[name];

	if !valid {
		return nil, errors.New("Not a valid subcommand: " + name)
	}

	if len(names) > 0 {
		return this.GetCommand(names, cmd)
	}

	return cmd, nil
}

func (opts *HelpCommand) Run() error {
	if len(opts.Args) == 0 {
		return errors.New("Please define a subcommand!")
	}

	cmd, err := opts.GetCommand(opts.Args, &opts.app.CategoryCommand.Command)

	if err != nil {
		return err
	}

	cmdPath := strings.Join(opts.Args, " ")

	if cmd.Help == "" {
		fmt.Printf("No help for subcommand: %s\n", cmdPath)
	} else {
		fmt.Println(cmd.Help)
		fmt.Printf("\nUsage: %s %s", opts.app.Name, cmdPath)

		if cmd.HasOptions() {
			fmt.Println(" [options]")
			fmt.Println("\nThe available options are:")
			cmd.WriteHelp(os.Stdout)
		} else {
			fmt.Println("")
		}
	}

	return nil
}
