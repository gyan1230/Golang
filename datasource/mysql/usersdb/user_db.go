package usersdb

import (
	"database/sql"
	"fmt"
	"log"

	//import driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	hostname = "localhost:3306"
	username = "root"
	password = "root"
	schema   = "users_db"
)

var (
	//UsersDB :
	UsersDB *sql.DB
)

func init() {
	var err error
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, hostname, schema)
	UsersDB, err = sql.Open("mysql", datasource)
	if err != nil {
		log.Printf("Database open error : %s\n", err)
		return
	}
	err = UsersDB.Ping()
	if err != nil {
		log.Printf("Ping error : %s\n", err)
		return
	}
	log.Println("Database connected sucessfully")
}
