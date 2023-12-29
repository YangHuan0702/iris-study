package service

import (
	"iris-study/src/domain"
	"iris-study/src/storage"
)

type UserService struct {
	dao *storage.UserStorage
}

var userService = &UserService{dao: storage.GetUserRepo()}

func GetUserService() *UserService {
	return userService
}

type UserServiceInter interface {
	Insert(user *domain.User)
}

func (userService *UserService) Insert(user *domain.User) {
	// that need set the user id
	user.Id = 1
	userService.dao.Insert(user)
}
