package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gyan1230/Golang/services"
	"github.com/gyan1230/Golang/utils"
)

//Ping :
func Ping(c *gin.Context) {
	resp := services.PingTest()
	c.JSON(resp.HTTPCode, utils.APIData{
		Data: resp,
	})

}
