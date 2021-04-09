package main

import (
	"github.com/printesoi/docker-github-actions/internal/command"
	"github.com/printesoi/docker-github-actions/internal/options"
)

func login(cmd command.Runner) error {
	o, err := options.GetLoginOptions()
	if err != nil {
		return err
	}

	return command.RunLogin(cmd, o, options.GetRegistry())
}
