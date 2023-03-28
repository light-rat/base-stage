package dao

import (
	"stage/http_server/model"
)

type UserQuery struct {
	Name  string
	Email string
}

func GetUserList(query UserQuery) ([]*model.User, int64, error) {
	modelArr := make([]*model.User, 0)
	count := int64(0)

	db := DB.Model(&model.User{})

	if query.Name != "" {
		db.Where("name like ?", query.Name+"%")
	}

	if query.Email != "" {
		db.Where("email like ?", query.Email+"%")
	}

	db.Count(&count)

	err := db.Scan(&modelArr).Error
	return modelArr, count, err
}
