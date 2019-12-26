package users

import (
	"fmt"
	"github.com/gyan1230/Golang/datasource/mysql/usersdb"
	"github.com/gyan1230/Golang/utils"
	"log"
)

const (
	queryInsert = "INSERT INTO users(name,age,salary) VALUES(?,?,?);"
	querySelect = "SELECT id, name, age, salary FROM users WHERE id=?;"
)

//SaveUser :
func (u *UserData) SaveUser() *utils.APIResponse {
	stmt, err := usersdb.UsersDB.Prepare(queryInsert)
	if err != nil {
		log.Println("Db query error:", err)
		res := utils.InternalServerError("Db query error:")
		return res
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Name, u.Age, u.Salary)
	if err != nil {
		log.Println("Saving into db error", err)
		res := utils.InternalServerError(fmt.Sprintf("Saving into db error :%s", err))
		return res
	}

	ID, err := result.LastInsertId()
	if err != nil {
		log.Println("DB Last Insert ID error:", err)
		res := utils.InternalServerError("DB Last Insert ID error:")
		return res
	}
	u.ID = ID
	log.Println("Sucessfully saved user :::")
	return nil

}

//GetByID :
func (u *UserData) GetByID() *utils.APIResponse {
	stmt, err := usersdb.UsersDB.Prepare(querySelect)
	if err != nil {
		log.Println("Db query error:", err)
		res := utils.InternalServerError("Db query error:")
		return res
	}
	defer stmt.Close()
	row := stmt.QueryRow(u.ID)
	if err := row.Scan(&u.ID, &u.Name, &u.Age, &u.Salary); err != nil {
		res := utils.InternalServerError(fmt.Sprintf("Getting from db error :%s\n", err))
		return res
	}
	return nil
}
