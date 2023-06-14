package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/cafemewa_cafe_mawa?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)

	return db
}
