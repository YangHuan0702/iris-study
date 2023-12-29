package storage

import (
	"github.com/jinzhu/gorm"
	"iris-study/src/domain"
)

type UserStorage struct {
	db *gorm.DB
}

var userDao = &UserStorage{db: GetDB()}

func GetUserRepo() *UserStorage {
	return userDao
}

type UserStorageRepo interface {
	Insert(user *domain.User)
}

func (userStorage *UserStorage) Insert(user *domain.User) {
	userStorage.db.Debug().Create(user)
}
