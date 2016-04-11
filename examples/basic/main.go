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
    handleErr := func(err error) {
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }

    app, err := cli.New("Do some stuff and things", &MainCmd{})
    handleErr(err)

    err = app.Run(nil)
    handleErr(err)
}
