package manageArticle

import (
	"errors"
	"time"

	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
	"github.com/KumKeeHyun/web-tuto-with-gin/domain/repository"
)

type manageArticle struct {
	ar repository.ArticleRepo
	ur repository.UserRepo
}

func NewManageArticleUsecase(ar repository.ArticleRepo, ur repository.UserRepo) *manageArticle {
	return &manageArticle{
		ar: ar,
		ur: ur,
	}
}

func (mauc *manageArticle) GetAllArticles() ([]model.Article, error) {
	return mauc.ar.GetAll()
}

func (mauc *manageArticle) GetArticleByID(id int) (*model.Article, error) {
	return mauc.ar.GetByID(id)
}

func (mauc *manageArticle) CreateNewArticle(title, content string, writerID int) (*model.Article, error) {
	a := model.Article{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		WriterID:  writerID,
	}
	if err := mauc.isUserExist(writerID); err != nil {
		return nil, err
	}
	return mauc.ar.Create(&a)
}

func (mauc *manageArticle) isUserExist(id int) error {
	u, err := mauc.ur.GetByID(id)
	if u == nil {
		errors.New("Unknown user")
	}
	if err != nil {
		return err
	}
	return nil
}

func (mauc *manageArticle) DeleteArticleByID(id int) error {
	a := model.Article{
		ID: id,
	}

	return mauc.ar.Delete(&a)
}
