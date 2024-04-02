package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"raihanki/learn_go_api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learngoapi")
	helper.PanicIfError(err)

	return db
}
