package services

import (
	"github.com/gyan1230/Golang/domain/users"
	"github.com/gyan1230/Golang/utils"
)

//CreateUserService :
func CreateUserService(u users.UserData) (*users.UserData, *utils.APIResponse) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}
