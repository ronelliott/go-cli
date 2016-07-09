package cli

import (
	"fmt"
	"sort"
	"strings"
)

type CategoryCommand struct {
	Command
}

func NewCategoryCommand(name, description, help string) (*CategoryCommand, error) {
	catCmd := &CategoryCommand{}
	cmd, err := NewCommand(name, description, help, catCmd)

	if err != nil {
		return nil, err
	}

	catCmd.Command = *cmd

	return catCmd, nil
}

func (opts *CategoryCommand) Run() error {
	if len(opts.Command.Subs) == 0 {
		return nil
	}

	fmt.Println("The available subcommands are:\n")
	maxNameLen := 0

	// find the command with the largest length
	for _, cmd := range opts.Command.Subs {
		curLen := len(cmd.Name)

		if curLen > maxNameLen {
			maxNameLen = curLen
		}
	}

	names := make([]string, len(opts.Command.Subs))
	idx := 0

	for name, _ := range opts.Command.Subs {
		names[idx] = name
		idx++
	}

	sort.Strings(names)

	for idx := range names {
		cmd := opts.Command.Subs[names[idx]]
		fmt.Printf("\t%s", cmd.Name)
		fmt.Print(strings.Repeat(" ", maxNameLen - len(cmd.Name)))
		fmt.Println("\t", cmd.Description)
	}

	fmt.Println()
	return nil
}
