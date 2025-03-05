package cmd

import (
	"fmt"
	"go_crud/config"
	"go_crud/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	fmt.Println(c.config.Server.Port)
	c.network.ServerStart(c.config.Server.Port)

	return c
}
