# go-cli

[![GoDoc](https://godoc.org/github.com/ronelliott/go-cli?status.png)](https://godoc.org/github.com/ronelliott/go-cli)
[![Build Status](https://travis-ci.org/ronelliott/go-cli.svg?branch=master)](https://travis-ci.org/ronelliott/go-cli)
[![Coverage Status](https://img.shields.io/coveralls/ronelliott/go-cli.svg)](https://coveralls.io/r/ronelliott/go-cli?branch=master)

a go library for writing command line interfaces

## Installation

    $ go get github.com/ronelliott/go-cli

## Examples

### Basic:

```go
package main

import (
    "github.com/ronelliott/go-cli"
    "fmt"
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
```

### Subcommands:

* a more complete example (with subcommands) can be found here: [examples/subcommands](https://github.com/ronelliott/go-cli/examples/subcommands)
