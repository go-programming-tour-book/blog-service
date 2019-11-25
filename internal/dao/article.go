package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type Article struct {
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
		Model: &model.Model{
			CreatedBy: param.CreatedBy,
		},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) (*model.Article, error) {

}

func (d *Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) GetArticleListByIDs(ids []uint32, state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{State: state}
	return article.ListByIDs(d.engine, ids, app.GetPageOffset(page, pageSize), pageSize)
}
