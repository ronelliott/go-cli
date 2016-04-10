package main

import (
    "errors"
    "fmt"
)

type HelpCmd struct {
    Args []string `positional:"true"`
}

func (opts *HelpCmd) Run() error {
    if len(opts.Args) != 1 {
        return errors.New("Please define a subcommand!")
    }

    name := opts.Args[0]

    if cmd, valid := app.Subs[name]; valid {
        if cmd.Help == "" {
            fmt.Printf("No help for subcommand: %s\n", name)
        } else {
            fmt.Println(cmd.Help, "\n")
            fmt.Printf("Usage: %s %s [options]\n", app.Name, name)
            fmt.Println("\nThe available options are:")
            // cmd.Flags.SetOutput(os.Stdout)
            // cmd.Flags.PrintDefaults()
        }
    } else {
        fmt.Printf("Invalid subcommand: %s\n", name)
    }

    return nil
}

func init() {
    app.NewSub("help", "Help on using my awesome CLI", "See `help` :P", &HelpCmd{})
}
