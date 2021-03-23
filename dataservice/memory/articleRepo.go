package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/web-tuto-with-gin/domain/model"
)

type memoryRepo struct {
	amu         *sync.Mutex
	articleList []model.Article
	articleID   int
	umu         *sync.Mutex
	userList    []model.User
	userID      int
}

var memoryDB *memoryRepo

func Setup() {
	memoryDB = &memoryRepo{
		amu:         &sync.Mutex{},
		articleList: make([]model.Article, 0),
		articleID:   0,
		umu:         &sync.Mutex{},
		userList:    make([]model.User, 0),
		userID:      0,
	}
}

type articleRepo struct {
	db *memoryRepo
}

func NewArticleRepo() *articleRepo {
	return &articleRepo{
		db: memoryDB,
	}
}

func (ar *articleRepo) GetAll() ([]model.Article, error) {
	ar.db.amu.Lock()
	defer ar.db.amu.Unlock()

	for i, a := range ar.db.articleList {
		for _, u := range ar.db.userList {
			if u.ID == a.WriterID {
				ar.db.articleList[i].Writer = u
				break
			}
		}
	}
	return ar.db.articleList, nil
}

func (ar *articleRepo) GetByID(id int) (*model.Article, error) {
	ar.db.amu.Lock()
	defer ar.db.amu.Unlock()
	ar.db.umu.Lock()
	defer ar.db.umu.Unlock()

	for _, a := range ar.db.articleList {
		if a.ID == id {
			for _, u := range ar.db.userList {
				if u.ID == a.WriterID {
					a.Writer = u
					break
				}
			}
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func (ar *articleRepo) Create(article *model.Article) (*model.Article, error) {
	ar.db.amu.Lock()
	defer ar.db.amu.Unlock()

	article.ID = ar.db.articleID
	article.CreatedAt = time.Now()
	ar.db.articleID++

	ar.db.articleList = append(ar.db.articleList, *article)

	return article, nil
}

func (ar *articleRepo) Delete(article *model.Article) error {
	ar.db.amu.Lock()
	defer ar.db.amu.Unlock()

	for i, a := range ar.db.articleList {
		if a.ID == article.ID {
			ar.db.articleList = append(ar.db.articleList[:i], ar.db.articleList[i+1:]...)
			return nil
		}
	}
	return errors.New("Article not found")
}
