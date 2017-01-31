package cli

import (
	"fmt"
	"os"
)

// Run the defined app
func run(cfg *AppConfig, withHelp bool) error {
	app, err := cfg.Convert(withHelp)

	if err != nil {
		return err
	}

	err = app.CategoryCommand.Command.Run(nil)

	if err != nil {
		return err
	}

	return nil
}

// Run the defined app without adding a help command
func Run(cfg *AppConfig) error {
	return run(cfg, false)
}

// Run the defined app and adds a help command
func RunWithHelp(cfg *AppConfig) error {
	return run(cfg, true)
}

// Run the defined app, adding the help command, and printing errors to stderr
func RunWithHelpAndErrors(cfg *AppConfig) error {
	err := RunWithHelp(cfg)

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%s\n", err.Error()))
	}

	return err
}

// Run the defined app, printing errors to stderr
func RunWithErrors(cfg *AppConfig) error {
	err := Run(cfg)

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%s\n", err.Error()))
	}

	return err
}

// Run the defined app, adding the help command, printing errors to stderr, and
// exiting with a non-zero code if an error occurs
func RunWithHelpErrorsAndExit(cfg *AppConfig) error {
	err := RunWithHelpAndErrors(cfg)

	if err != nil {
		os.Exit(1)
	}

	return err
}

// Run the defined app, printing errors to stderr, and exiting with a non-zero
// code if an error occurs
func RunWithErrorsAndExit(cfg *AppConfig) error {
	err := RunWithErrors(cfg)

	if err != nil {
		os.Exit(1)
	}

	return err
}
