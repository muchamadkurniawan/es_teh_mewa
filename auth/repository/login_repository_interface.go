package repository

import (
	"context"
	"database/sql"
	"es_teh_mewa/auth/model"
)

type LoginRepository interface {
	Login(ctx context.Context, sql *sql.Tx, user model.UserLogin) model.UserLogin
}
