package cmd

import (
	"fmt"
	"golang_backend_study/config"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filePath string) *Cmd {
	c := Cmd{
		config: config.NewConfig(filePath),
	}
	fmt.Println(c.config.Server.Port)

	return &c
}
