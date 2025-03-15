package cmd

import (
	"go_crud/config"
	"go_crud/network"
	"go_crud/repository"
	"go_crud/service"
	"os"
)

type Cmd struct {
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd() *Cmd {
	config.LoadEnvVariables()
	config.ConnectToDb()
	//config.SyncDatabase()
	var c = &Cmd{}
	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(os.Getenv("PORT"))

	return c
}
