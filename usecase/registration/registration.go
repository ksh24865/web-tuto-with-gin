package registration

import (
	"errors"
	"strings"

	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
	"github.com/KumKeeHyun/web-tuto-with-gin/domain/repository"
)

type registrationUsecase struct {
	ur repository.UserRepo
}

func NewRegistrationUsecase(ur repository.UserRepo) *registrationUsecase {
	return &registrationUsecase{
		ur: ur,
	}
}

func (ru *registrationUsecase) RegisterUser(name, pass string) (*model.User, error) {
	u := model.User{
		Username: name,
		Password: pass,
	}
	if err := ru.isValid(&u); err != nil {
		return nil, err
	}
	// TODO : Password must be encrypted through an algorithm like hash
	res, err := ru.ur.Create(&u)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (ru *registrationUsecase) isValid(newUser *model.User) error {
	u, _ := ru.ur.GetByName(newUser.Username)
	if u.Username != "" {
		return errors.New("username is already exist")
	}
	if strings.TrimSpace(newUser.Password) == "" {
		return errors.New("The password can't be empty")
	}
	return nil
}

func (ru *registrationUsecase) MatchUser(name, pass string) (*model.User, error) {
	u, err := ru.ur.GetByName(name)
	if err != nil {
		return nil, err
	}
	if u.Password != pass {
		return nil, errors.New("wrong password")
	}
	return u, nil
}
