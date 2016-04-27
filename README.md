# go-cli

[![GoDoc](https://godoc.org/github.com/ronelliott/go-cli?status.png)](https://godoc.org/github.com/ronelliott/go-cli)
[![Build Status](https://travis-ci.org/ronelliott/go-cli.svg?branch=master)](https://travis-ci.org/ronelliott/go-cli)
[![Coverage Status](https://coveralls.io/repos/github/ronelliott/go-cli/badge.svg?branch=master)](https://coveralls.io/github/ronelliott/go-cli?branch=master)

A go library for writing command line interfaces. Only supports go versions newer than, or equal to, 1.5

## Installation

    $ go get github.com/ronelliott/go-cli

## Examples

### Basic:

```go
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
```

### Subcommands:

* a more complete example (with subcommands) can be found here: [examples/subcommands](https://github.com/ronelliott/go-cli/tree/master/examples/subcommands)

[![Analytics](https://ga-beacon.appspot.com/UA-59523757-2/go-cli/readme?pixel)](https://github.com/igrigorik/ga-beacon)
