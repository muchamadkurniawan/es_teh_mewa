package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/master/model/entity"
)

type usersRepositoryImpl struct {
	DB *sql.DB
}

func NewUsersRepository() *usersRepositoryImpl {
	return &usersRepositoryImpl{}
}

func (usersRepositoryImpl) InsertUsers(ctx context.Context, tx *sql.Tx, user entity.Users) {
	SQL := "INSERT INTO users (username, password, type) values (?,?,?);"
	_, err := tx.ExecContext(ctx, SQL, user.UserName, user.Password, user.Type_user)
	if err != nil {
		panic(err)
	}

}

func (usersRepositoryImpl) FindByIdUsers(ctx context.Context, tx *sql.Tx, id int32) (entity.Users, error) {
	SQL := "SELECT id, username, password, type FROM users WHERE id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	user := entity.Users{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Type_user)
		if err != nil {
			return entity.Users{}, err
		}
	}
	return user, nil
}

func (usersRepositoryImpl) FindByAllUsers(ctx context.Context, tx *sql.Tx) ([]entity.Users, error) {
	SQL := "SELECT id, username, password, type FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []entity.Users
	for rows.Next() {
		user := entity.Users{}
		err := rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Type_user)
		if err != nil {
			panic(err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (usersRepositoryImpl) UpdateUsers(ctx context.Context, tx *sql.Tx, user entity.Users) {
	SQL := "UPDATE users SET username = ?, password = ?, type = ? where id = ?;"
	_, err := tx.ExecContext(ctx, SQL, user.UserName, user.Password, user.Type_user, user.Id)
	helperMain.PanicIfError(err)
}

func (usersRepositoryImpl) DeleteUsers(ctx context.Context, tx *sql.Tx, id int32) {
	SQL := "DELETE FROM users where id = ?"
	_, err := tx.ExecContext(context.Background(), SQL, id)
	if err != nil {
		panic(err)
	}
}
