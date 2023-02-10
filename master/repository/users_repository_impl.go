package repository

import (
	"context"
	"database/sql"
	"eh_teh_mewa/master/model/entity"
)

type usersRepositoryImpl struct {
	DB *sql.DB
}

func NewUsersRepository() *usersRepositoryImpl {
	return &usersRepositoryImpl{}
}

func (usersRepositoryImpl) InsertUsers(ctx context.Context, tx *sql.Tx, user entity.Users) (entity.Users, error) {
	SQL := "INSERT INTO users (username, password, type) values (?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.UserName, user.Password, user.Type_user)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.Id = int32(int(id))
	return user, nil
}

func (usersRepositoryImpl) FindByIdUsers(ctx context.Context, tx *sql.Tx, id int32) (entity.Users, error) {
	SQL := "SELECT username, password, type FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	user := entity.Users{}
	if rows.Next() {
		err = rows.Scan(&user.UserName, &user.Password, &user.Type_user)
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

func (usersRepositoryImpl) UpdateUsers(ctx context.Context, tx *sql.Tx, user entity.Users) (entity.Users, error) {
	SQL := "update users set username = ?, password = ?, type = ? where id = ?"
	_, err := tx.ExecContext(context.Background(), SQL, user.UserName, user.Password, user.Type_user, user.Id)
	if err != nil {
		return entity.Users{}, err
		panic(err)
	}
	return user, nil
}

func (usersRepositoryImpl) DeleteUsers(ctx context.Context, tx *sql.Tx, id int32) error {
	SQL := "DELETE FROM users where id = ?"
	_, err := tx.ExecContext(context.Background(), SQL, id)
	if err != nil {
		panic(err)
	}
	return err
}

//func (repository *usersRepositoryImpl) InsertUser(ctx context.Context, user entity.Users) (entity.Users, error) {
//	script := "INSERT INTO users(nama, password, type) VALUES(?, ?, ?);"
//	result, err := repository.DB.ExecContext(ctx, script, user.UserName, user.Password, user.Type_user)
//	defer repository.DB.Close()
//	if err != nil {
//		return user, err
//	}
//	id, err := result.LastInsertId()
//	if err != nil {
//		return user, err
//	}
//	user.Id = int32(id)
//	return user, nil
//}
//
//func (repository *usersRepositoryImpl) FindByIdUser(ctx context.Context, id int32) (entity.Users, error) {
//	script := "SELECT id, nama, password, type FROM users WHERE id = ? LIMIT 1;"
//	result, err := repository.DB.QueryContext(ctx, script, id)
//	defer repository.DB.Close()
//	user := entity.Users{}
//	defer result.Close()
//	if err != nil {
//		return user, err
//	}
//	if result.Next() {
//		result.Scan(&user.Id, &user.UserName, &user.Password, &user.Type_user)
//		return user, nil
//	} else {
//		return user, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
//	}
//}
//
//func (repository *usersRepositoryImpl) FindByAllUser(ctx context.Context, tx sql.Tx) ([]entity.Users, error) {
//	script := "SELECT id, nama, password, type FROM users;"
//	result, err := repository.DB.QueryContext(ctx, script)
//	if err != nil {
//		return nil, err
//	}
//	var users []entity.Users
//	defer result.Close()
//	defer repository.DB.Close()
//	for result.Next() {
//		user := entity.Users{}
//		result.Scan(&user.Id, &user.UserName, &user.Password, &user.Type_user)
//		users = append(users, user)
//	}
//	return users, nil
//}
