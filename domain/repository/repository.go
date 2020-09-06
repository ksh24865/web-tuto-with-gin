package repository

import "github.com/KumKeeHyun/web-tuto-with-gin/domain/model"

type ArticleRepo interface {
	GetAll() ([]model.Article, error)
	GetByID(id int) (*model.Article, error)
	Create(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}

type UserRepo interface {
	GetByID(id int) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
}
