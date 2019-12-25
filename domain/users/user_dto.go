package users

import (
	"strings"

	"github.com/gyan1230/Golang/utils"
)

//UserData :
type UserData struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Salary float32 `json:"salary"`
}

//Validate :
func (u *UserData) Validate() *utils.APIResponse {
	u.Name = strings.TrimSpace(strings.ToLower(u.Name))
	if u.Name == "" {
		resterr := utils.NewBadrequest("Invalid name")
		return resterr
	}
	return nil
}
