package model

import (
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (t Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&t).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

func (t Article) Get(db *gorm.DB) (*Article, error) {
	var article *Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", t.ID, t.State, 0)
	if err := db.First(&article).Error; err != nil {
		return nil, err
	}

	return article, nil
}

func (t Article) ListByIDs(db *gorm.DB, ids []uint32, pageOffset, pageSize int) ([]*Article, error) {
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	var articles []*Article
	db = db.Where("id IN (?) AND state = ? AND is_del = ?", ids, t.State, 0)
	err := db.First(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}
