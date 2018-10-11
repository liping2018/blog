package service

import (
	"my_blog/models"
)

//根据用户名查找用户信息
func QueryUserInfoByUername(username string) *models.User {
	return models.QueryUserInfoByUsername(username)
}
