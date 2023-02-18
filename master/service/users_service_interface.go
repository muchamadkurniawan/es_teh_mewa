package service

import (
	"context"
	"eh_teh_mewa/master/web"
)

type UsersService interface {
	Create(ctx context.Context, request web.UsersCreateRequest)
	Update(ctx context.Context, response web.UsersResponse)
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []web.UsersResponse
	FindById(ctx context.Context, id int) web.UsersResponse
}
