package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Index :
func Index(c *gin.Context) {
	log.Println("In home:::")
	c.JSON(200, "In home")
}

//GetLikes :
func GetLikes(c *gin.Context) {
	c.JSON(200, "Likes Received...")
}
