package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

func (d *Dao) CreateArticle(title, desc, content, coverImageUrl string, state uint8) (*model.Article, error) {
	article := model.Article{Title: title, Desc: desc, Content: content, CoverImageUrl: coverImageUrl, State: state}
	return article.Create(d.engine)
}

func (d *Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) GetArticleListByIDs(ids []uint32, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{State: state}
	return article.ListByIDs(d.engine, ids, app.GetPageOffset(page, pageSize), pageSize)
}
