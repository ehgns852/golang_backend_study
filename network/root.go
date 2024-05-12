package network

import "github.com/gin-gonic/gin"

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r)

	return r
}

func (n *Network) ServerStart(port string) {
	if err := n.engin.Run(port); err != nil {
		panic(err)
	}
}
