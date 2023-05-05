package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/auth/model"
)

type LoginRepositoryImpl struct {
	DB *sql.DB
}

func NewLoginRepository() *LoginRepositoryImpl {
	return &LoginRepositoryImpl{}
}

func (LoginRepositoryImpl) Login(ctx context.Context, sql *sql.Tx, user model.UserLogin) model.UserLogin {
	SQL := "SELECT username, password, type FROM users WHERE username = ? limit 1;"
	row, _ := sql.QueryContext(ctx, SQL, user.Nama)
	defer row.Close()
	userLog := model.UserLogin{}
	if row.Next() {
		err := row.Scan(&userLog.Nama, &userLog.Password, &userLog.Type)
		if err != nil {
			return model.UserLogin{
				Nama:     "",
				Password: "",
				Type:     "",
			}
		}
	}
	return userLog
}
