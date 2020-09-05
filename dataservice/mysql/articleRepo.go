package mysql

import (
	"time"

	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
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

func (ar *articleRepo) GetAll() (al []model.Article) {
	err := ar.db.Find(&al).Error
	if err != nil {
		return []model.Article{}
	}
	return
}

func (ar *articleRepo) GetByID(id int) (a *model.Article, err error) {
	a = new(model.Article)
	return a, ar.db.Where("id=?", id).First(a).Error
}

func (ar *articleRepo) Create(title, content string) (a *model.Article, err error) {
	a = &model.Article{
		Title: title, Content: content,
		CreatedAt: time.Now(),
	}

	return a, ar.db.Create(a).Error
}

func (ar *articleRepo) DeleteByID(id int) error {
	a := model.Article{
		ID: id,
	}
	return ar.db.Delete(&a).Error
}
