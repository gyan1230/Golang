package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/gyan1230/Golang/services"
	"github.com/gyan1230/Golang/utils"
)

//PingEndpoint :
func PingEndpoint(c *gin.Context) {
	resp := services.PingTest()
	c.JSON(resp.HTTPCode, utils.APIData{
		Data: resp,
	})

}
