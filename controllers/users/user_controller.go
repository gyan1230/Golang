package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gyan1230/Golang/datasource/mysql/usersdb"
	"github.com/gyan1230/Golang/domain/users"
	"github.com/gyan1230/Golang/services"
	"github.com/gyan1230/Golang/utils"
	"log"
	"net/http"
	"strconv"
)

//Index :
func Index(c *gin.Context) {
	log.Println("In home:::")
	c.JSON(200, "In home")
}

func ping() error {
	err := usersdb.UsersDB.Ping()
	if err != nil {
		log.Printf("Ping error : %s\n", err)
		return err
	}
	log.Println("Ping sucessfull")
	return nil
}

//GetUser :
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Println("url parse error:")
		resterr := utils.InternalServerError("url parse error:")
		c.JSON(resterr.HTTPCode, utils.APIData{
			Data: *resterr,
		},
		)
		return
	}
	log.Printf("User Id: %d\n", userID)
	result, res := services.GetUserService(userID)
	if err != nil {
		log.Println("user not found error:", res.Msg)
		resterr := utils.InternalServerError(res.Msg)
		c.JSON(resterr.HTTPCode, utils.APIData{
			Data: *resterr,
		},
		)
		return
	}
	log.Println(result)
	c.JSON(200, result)
	return
}

//CreateUser :
func CreateUser(c *gin.Context) {
	var u users.UserData

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	log.Println("error:", err)
	// 	return
	// }
	// err = json.Unmarshal(bytes, &u)
	// if err != nil {
	// 	log.Println("error:", err)
	// 	return
	// }

	//above thing replaced by
	if err := c.ShouldBindJSON(&u); err != nil {
		resterr := utils.NewBadrequest("Invalid JSON body")
		c.JSON(resterr.HTTPCode, utils.APIData{
			Data: *resterr,
		},
		)
		return
	}
	res, err := services.CreateUserService(u)
	if err != nil {
		log.Println("user creation error:", err.Msg)
		resterr := utils.InternalServerError(err.Msg)
		c.JSON(resterr.HTTPCode, utils.APIData{
			Data: *resterr,
		},
		)
		return
	}
	c.JSON(http.StatusCreated, res)
}
