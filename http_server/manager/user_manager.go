package manager

import (
	"stage/http_server/dao"
	"stage/http_server/model"
	"stage/http_server/request"
	"stage/http_server/response"
)

func GetUserList(request request.GetUserListRequest) (response.GetUserListResponse, error) {
	result := response.GetUserListResponse{List: make([]response.UserListItem, 0)}

	query := dao.UserQuery{
		Name:  request.Name,
		Email: request.Email,
	}

	userModels, total, err := dao.GetUserList(query)
	if err != nil {
		return result, err
	}
	result.Count = total
	GetUserListFormat(&result, userModels)
	return result, err
}

func GetUserListFormat(result *response.GetUserListResponse, models []*model.User) {
	for _, userModel := range models {
		result.List = append(result.List, response.UserListItem{
			Name:  userModel.Name,
			Home:  userModel.Home,
			Email: userModel.Email,
			Sex:   userModel.Sex,
		})
	}
}
