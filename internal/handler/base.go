package handler

import (
	"github.com/gin-gonic/gin"
)

// var (
// 	ServerAddr = flag.String("server_addr", "localhost: 3001", "The server address in the format of host:port")
// )

type Controller interface {
	Handler() gin.HandlerFunc
	Route() string
	Method() string
}
