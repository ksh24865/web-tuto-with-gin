package manageArticle

import (
	"time"

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

func (mauc *manageArticle) GetAllArticles() ([]model.Article, error) {
	return mauc.ar.GetAll()
}

func (mauc *manageArticle) GetArticleByID(id int) (*model.Article, error) {
	return mauc.ar.GetByID(id)
}

func (mauc *manageArticle) CreateNewArticle(title, content string) (*model.Article, error) {
	a := model.Article{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	return mauc.ar.Create(&a)
}

func (mauc *manageArticle) DeleteArticleByID(id int) error {
	a := model.Article{
		ID: id,
	}
	return mauc.ar.Delete(&a)
}
