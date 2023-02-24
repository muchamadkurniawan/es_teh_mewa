package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type UsersRepository interface {
	InsertUsers(ctx context.Context, tx *sql.Tx, user entity.Users)
	FindByIdUsers(ctx context.Context, tx *sql.Tx, id int32) (entity.Users, error)
	FindByAllUsers(ctx context.Context, tx *sql.Tx) ([]entity.Users, error)
	UpdateUsers(ctx context.Context, tx *sql.Tx, user entity.Users)
	DeleteUsers(ctx context.Context, tx *sql.Tx, id int32)
}
