package manageArticle

import (
	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
	"github.com/KumKeeHyun/web-tuto-with-gin/domain/repository"
)

type manageArticle struct {
	ar repository.ArticleRepo
}

func NewManageArticleUsecase(ar repository.ArticleRepo) *manageArticle {
	return &manageArticle{
		ar: ar,
	}
}

func (mauc *manageArticle) GetAllArticles() []model.Article {
	return mauc.ar.GetAll()
}

func (mauc *manageArticle) GetArticleByID(id int) (*model.Article, error) {
	return mauc.ar.GetByID(id)
}

func (mauc *manageArticle) CreateNewArticle(title, content string) (*model.Article, error) {
	return mauc.ar.Create(title, content)
}

func (mauc *manageArticle) DeleteArticleByID(id int) error {
	return mauc.ar.DeleteByID(id)
}
