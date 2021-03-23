package mysql

import (
	"github.com/web-tuto-with-gin/domain/model"
	"github.com/jinzhu/gorm"
)

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo() *articleRepo {
	return &articleRepo{
		db: dbConn,
	}
}

func (ar *articleRepo) GetAll() (al []model.Article, err error) {
	return al, ar.db.Preload("Writer").Find(&al).Error
	//return al, ar.db.Joins("User").Find(&al).Error
}

func (ar *articleRepo) GetByID(id int) (a *model.Article, err error) {
	a = new(model.Article)
	return a, ar.db.Preload("Writer").Where("id=?", id).First(a).Error
}

func (ar *articleRepo) Create(article *model.Article) (*model.Article, error) {
	return article, ar.db.Create(article).Error
}

func (ar *articleRepo) Delete(article *model.Article) error {
	return ar.db.Delete(article).Error
}
