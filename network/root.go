package network

import (
	"github.com/gin-gonic/gin"
	"golang_backend_study/service"
)

type Network struct {
	engin   *gin.Engine
	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engin:   gin.New(),
		service: service,
	}

	newUserRouter(r, service.User)

	return r
}

func (n *Network) ServerStart(port string) {
	if err := n.engin.Run(port); err != nil {
		panic(err)
	}
}
