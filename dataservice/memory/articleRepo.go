package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/KumKeeHyun/web-tuto-with-gin/domain/model"
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

func (ar *articleRepo) GetAll() []model.Article {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	return ar.articleList
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

func (ar *articleRepo) Create(title, content string) (*model.Article, error) {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	a := model.Article{
		ID:    ar.id,
		Title: title, Content: content,
		CreatedAt: time.Now(),
	}
	ar.id++

	ar.articleList = append(ar.articleList, a)

	return &a, nil
}

func (ar *articleRepo) DeleteByID(id int) error {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	for i, a := range ar.articleList {
		if a.ID == id {
			ar.articleList = append(ar.articleList[:i], ar.articleList[i+1:]...)
			return nil
		}
	}
	return errors.New("Article not found")
}
