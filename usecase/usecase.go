package usecase

import "github.com/KumKeeHyun/web-tuto-with-gin/domain/model"

type ManageArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
	GetArticleByID(id int) (*model.Article, error)
	CreateNewArticle(title, content string, writerID int) (*model.Article, error)
	DeleteArticleByID(id int) error
}

type RegistrationUsecase interface {
	RegisterUser(name, pass string) (*model.User, error)
	MatchUser(name, pass string) (*model.User, error)
}
