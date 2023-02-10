package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/cafe_mawah?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20)

	return db
}
