package helperMain

import (
	"eh_teh_mewa/master/model/entity"
	"eh_teh_mewa/master/web"
)

func ToUserResponses(users []entity.Users) []web.UsersResponse {
	var userResponses []web.UsersResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToUserResponse(user entity.Users) web.UsersResponse {
	return web.UsersResponse{
		Id:       int(user.Id),
		Username: user.UserName,
		Password: user.Password,
	}
}
