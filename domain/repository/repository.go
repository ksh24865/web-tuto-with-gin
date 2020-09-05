package repository

import "github.com/KumKeeHyun/web-tuto-with-gin/domain/model"

type ArticleRepo interface {
	GetAll() []model.Article
	GetByID(id int) (*model.Article, error)
	Create(title, content string) (*model.Article, error)
	DeleteByID(id int) error
}
