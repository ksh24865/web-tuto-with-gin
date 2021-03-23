package memory

import (
	"errors"
	"sync"

	"github.com/web-tuto-with-gin/domain/model"
)

type articleRepo struct {
	mu          *sync.Mutex
	articleList []model.Article
	id          int
}

func NewArticleRepo() *articleRepo {
	return &articleRepo{
		mu:          &sync.Mutex{},
		articleList: make([]model.Article, 0),
		id:          0,
	}
}

func (ar *articleRepo) GetAll() ([]model.Article, error) {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	return ar.articleList, nil
}

func (ar *articleRepo) GetByID(id int) (*model.Article, error) {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	for _, a := range ar.articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func (ar *articleRepo) Create(article *model.Article) (*model.Article, error) {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	article.ID = ar.id
	ar.id++

	ar.articleList = append(ar.articleList, *article)

	return article, nil
}

func (ar *articleRepo) Delete(article *model.Article) error {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	for i, a := range ar.articleList {
		if a.ID == article.ID {
			ar.articleList = append(ar.articleList[:i], ar.articleList[i+1:]...)
			return nil
		}
	}
	return errors.New("Article not found")
}
