package main

import (
    "fmt"
    "github.com/ronelliott/go-cli"
    "os"
    "sort"
    "strings"
)

type MainCmd struct {}

func (opts *MainCmd) Run() error {
    fmt.Println(app.Description, "\n")
    fmt.Println("\nUsage:")
    fmt.Println("\t", app.Name, "<subcommand> [options]")
    fmt.Println("\nThe subcommands are:\n")

    maxNameLen := 0

    // find the command with the largest length
    for _, cmd := range app.Subs {
        curLen := len(cmd.Name)

        if curLen > maxNameLen {
            maxNameLen = curLen
        }
    }

    names := make([]string, len(app.Subs))
    idx := 0

    for name, _ := range app.Subs {
        names[idx] = name
        idx++
    }

    sort.Strings(names)

    for idx := range names {
        cmd := app.Subs[names[idx]]
        fmt.Printf("\t%s", cmd.Name)
        fmt.Print(strings.Repeat(" ", maxNameLen - len(cmd.Name)))
        fmt.Println("\t", cmd.Description)
    }

    return nil
}

var app *cli.Command = cli.New("Do some stuff and things", &MainCmd{})

func main() {
    if err := app.Run(nil); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
