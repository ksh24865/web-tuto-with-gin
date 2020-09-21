package memory

import (
	"errors"

	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
)

type userRepo struct {
	db *memoryRepo
}

func NewUserRepo() *userRepo {
	return &userRepo{
		db: memoryDB,
	}
}

func (ur *userRepo) GetByID(id int) (u *model.User, err error) {
	ur.db.umu.Lock()
	defer ur.db.umu.Unlock()

	for _, u := range ur.db.userList {
		if u.ID == id {
			return &u, nil
		}
	}
	return &model.User{Username: ""}, errors.New("User not found")
}

func (ur *userRepo) GetByName(name string) (u *model.User, err error) {
	ur.db.umu.Lock()
	defer ur.db.umu.Unlock()

	for _, u := range ur.db.userList {
		if u.Username == name {
			return &u, nil
		}
	}
	return &model.User{Username: ""}, errors.New("User not found")
}

func (ur *userRepo) Create(user *model.User) (*model.User, error) {
	ur.db.umu.Lock()
	defer ur.db.umu.Unlock()

	user.ID = ur.db.userID
	ur.db.userID++

	ur.db.userList = append(ur.db.userList, *user)
	return user, nil
}
