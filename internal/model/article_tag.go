package model

import (
	"github.com/jinzhu/gorm"
)

type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (t ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND is_del = ?", t.ArticleID, 0).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

func (t ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ? AND is_del = ?", t.TagID, 0).Find(&articleTags).Error; err != nil {
		return nil, err
	}

	return articleTags, nil
}

func (t ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleTags, nil
}

func (t ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&t).Error; err != nil {
		return err
	}

	return nil
}

func (t ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&t).Where("article_id = ? AND is_del = ?", t.ArticleID, 0).Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (t ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error; err != nil {
		return err
	}

	return nil
}

func (t ArticleTag) DeleteOne(db *gorm.DB) error {
	if err := db.Where("article_id = ? AND is_del = ?", t.ArticleID, 0).Delete(&t).Limit(1).Error; err != nil {
		return err
	}

	return nil
}
