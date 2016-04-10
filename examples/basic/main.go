package main

import (
    "fmt"
    "github.com/ronelliott/go-cli"
    "os"
)

type MainCmd struct {
    Verbose bool `short:"v" long:"verbose"`
}

func (opts *MainCmd) Run() error {
    if opts.Verbose {
        fmt.Println("Doing stuff and things!!")
    } else {
        fmt.Println("SSSHHHHH!!!!")
    }

    return nil
}

func main() {
    app := cli.New("Do some stuff and things", &MainCmd{})

    if err := app.Run(nil); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
