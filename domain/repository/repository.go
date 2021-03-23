package repository

import "github.com/web-tuto-with-gin/domain/model"

type ArticleRepo interface {
	GetAll() ([]model.Article, error)
	GetByID(id int) (*model.Article, error)
	Create(article *model.Article) (*model.Article, error)
	Delete(article *model.Article) error
}
