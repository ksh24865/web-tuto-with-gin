package usecase

import "github.com/KumKeeHyun/web-tuto-with-gin/domain/model"

type ManageArticleUsecase interface {
	GetAllArticles() []model.Article
	GetArticleByID(id int) (*model.Article, error)
	CreateNewArticle(title, content string) (*model.Article, error)
	DeleteArticleByID(id int) error
}
