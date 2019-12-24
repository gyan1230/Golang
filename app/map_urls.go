package app

import (
	"github.com/gyan1230/Golang/controllers"
	"github.com/gyan1230/Golang/socket"
)

func mapURLs() {
	router.GET("/ping", controllers.Ping)
	router.GET("/", controllers.Index)
	router.GET("/twitter/likes", controllers.GetLikes)
	router.POST("/ws", socket.WsEndPoint)

}
