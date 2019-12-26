package app

import (
	"github.com/gyan1230/Golang/controllers/ping"
	"github.com/gyan1230/Golang/controllers/users"
	"github.com/gyan1230/Golang/socket"
)

func mapURLs() {
	router.GET("/ping", ping.PingEndpoint)
	router.GET("/", users.Index)
	router.GET("/user/:user_id", users.GetUser)
	router.POST("/user", users.CreateUser)
	router.POST("/ws", socket.WsEndPoint)

}
