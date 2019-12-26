package services

import (
	"github.com/gyan1230/Golang/domain/users"
	"github.com/gyan1230/Golang/utils"
	"log"
)

//CreateUserService :
func CreateUserService(u users.UserData) (*users.UserData, *utils.APIResponse) {
	if err := u.Validate(); err != nil {
		log.Println("user validation error:")
		return nil, err
	}
	if res := u.SaveUser(); res != nil {
		log.Println("save user response not nil:", res.Msg)
		return nil, res
	}
	log.Println("User saved:::")
	return &u, nil
}

//GetUserService :
func GetUserService(id int64) (*users.UserData, *utils.APIResponse) {
	res := &users.UserData{ID: id}
	if err := res.GetByID(); err != nil {
		return nil, err
	}

	return res, nil

}
