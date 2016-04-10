package main

import (
    "errors"
    "fmt"
    "strconv"
)

func fib(n int) int {
    if n == 0 || n == 1 {
        return n
    }

    return fib(n - 1) + fib(n - 2)
}

type FibCmd struct {
    Args []string `positional:"true"`
    Verbose bool `short:"v" description:"Use verbose output."`
}

func (opts *FibCmd) Run() error {
    fmt.Println("Running help subcommand!")
    if len(opts.Args) != 1 {
        return errors.New("Give me a number plzkthnx!")
    }

    max, err := strconv.Atoi(opts.Args[0])

    if err != nil {
        return errors.New(
            "Error whilest converting:" + opts.Args[0] + " " + err.Error())
    }

    for i := 0; i <= max; i++ {
        fmt.Println(i, "=>", fib(i))
    }

    return nil
}

func init() {
    app.NewSub("fib", "Fib it up yo!", "Run a fibonacci sequence", &FibCmd{})
}
