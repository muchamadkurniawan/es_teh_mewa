package service

import (
	"context"
	"database/sql"
	"eh_teh_mewa/helperMain"
	"eh_teh_mewa/master/model/entity"
	"eh_teh_mewa/master/repository"
	"eh_teh_mewa/master/web"
)

type UsersServiceImpl struct {
	UsersRepo repository.UsersRepository
	DB        *sql.DB
}

func NewUsersService(usersRepository repository.UsersRepository, DB *sql.DB) UsersService {
	return &UsersServiceImpl{
		UsersRepo: usersRepository,
		DB:        DB,
	}
}

func (service *UsersServiceImpl) Create(ctx context.Context, request web.UsersCreateRequest) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()
	user := entity.Users{
		UserName:  request.Username,
		Password:  request.Password,
		Type_user: request.Type,
	}
	service.UsersRepo.InsertUsers(ctx, tx, user)
}

func (service *UsersServiceImpl) Update(ctx context.Context, response web.UsersResponse) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	user := entity.Users{
		Id:        int32(response.Id),
		UserName:  response.Username,
		Password:  response.Password,
		Type_user: response.Type,
	}
	_, err = service.UsersRepo.UpdateUsers(ctx, tx, user)
	if err != nil {
		panic(err)
	}
}

func (service *UsersServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	err = service.UsersRepo.DeleteUsers(context.Background(), tx, int32(id))
	if err != nil {
		panic(err)
	}
}

func (service *UsersServiceImpl) FindAll(ctx context.Context) []map[string]interface{} {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()

	users, err := service.UsersRepo.FindByAllUsers(ctx, tx)
	if err != nil {
		panic(err)
	}
	usersResponse := helperMain.ToUserResponses(users)
	datas := helperMain.StructSliceToMap_Users(usersResponse)
	return datas
}

func (service *UsersServiceImpl) FindById(ctx context.Context, id int) map[string]interface{} {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()
	user, err := service.UsersRepo.FindByIdUsers(ctx, tx, int32(id))
	if err != nil {
		panic(err)
	}
	userResponse := helperMain.ToUserResponse(user)
	data := helperMain.StructToMap_Users(userResponse)
	return data
}
