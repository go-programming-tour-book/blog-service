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

func (t Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(t).Updates(values).Where("id = ? AND is_del = ?", t.ID).Error; err != nil {
		return err
	}

	return nil
}

func (t Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", t.ID, t.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}

	return article, nil
}

func (t Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table("`blog_article_tag` AS tb_article_tag").
		Joins("LEFT JOIN `blog_tag` AS tb_tag ON tb_article_tag.tag_id = tb_tag.id").
		Joins("LEFT JOIN `blog_article` AS tb_article ON tb_article_tag.article_id = tb_article.id ").
		Where("tb_article_tag.`tag_id` = ?", tagID).
		Where("tb_article.state = ?", t.State).
		Where("tb_article.is_del = 0").
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (t Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{
		"tb_article.id AS article_id",
		"tb_article.title AS article_title",
		"tb_article.desc AS article_desc",
		"tb_article.cover_image_url",
		"tb_article.content",
		"tb_article_tag.id AS tag_id",
		"tb_tag.name AS tag_name",
	}
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table("`blog_article_tag` AS tb_article_tag").
		Joins("LEFT JOIN `blog_tag` AS tb_tag ON tb_article_tag.tag_id = tb_tag.id").
		Joins("LEFT JOIN `blog_article` AS tb_article ON tb_article_tag.article_id = tb_article.id ").
		Where("tb_article_tag.`tag_id` = ?", tagID).
		Where("tb_article.state = ?", t.State).
		Where("tb_article.is_del = 0").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articleRows []*ArticleRow
	for rows.Next() {
		articleRow := &ArticleRow{}
		if err := rows.Scan(
			&articleRow.ArticleID,
			&articleRow.ArticleTitle,
			&articleRow.ArticleDesc,
			&articleRow.CoverImageUrl,
			&articleRow.Content,
			&articleRow.TagID,
			&articleRow.TagName,
		); err != nil {
			return nil, err
		}

		articleRows = append(articleRows, articleRow)
	}

	return articleRows, nil
}

func (t Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error; err != nil {
		return err
	}

	return nil
}
