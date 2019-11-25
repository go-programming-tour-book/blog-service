package service

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

type Article struct {
	ID            uint32 `json:"id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`

	Tag *model.Tag
}

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=2,max=100"`
	Desc          string `form:"desc" binding:"min=2,max=255"`
	Content       string `form:"content" binding:"min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"url"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	articleTag, err := svc.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(articleTag.TagID, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}

	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           tag,
	}, nil
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int, error) {
	articleTags, err := svc.dao.GetArticleTagListByTID(param.TagID)
	if err != nil {
		return nil, 0, err
	}

	var articleIDs []uint32
	var tagIDs []uint32
	var articleTagRelation = make(map[uint32]uint32)
	for _, articleTag := range articleTags {
		articleIDs = append(articleIDs, articleTag.ArticleID)
		tagIDs = append(tagIDs, articleTag.TagID)
		articleTagRelation[articleTag.ArticleID] = articleTag.TagID
	}

	articles, err := svc.dao.GetArticleListByIDs(articleIDs, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	tags, err := svc.dao.GetTagListByIDs(tagIDs, model.STATE_OPEN)
	if err != nil {
		return nil, 0, err
	}

	var articleResults []*Article
	for _, article := range articles {
		if tagId, ok := articleTagRelation[article.ID]; ok {
			for _, tag := range tags {
				if tag.ID == tagId {
					articleResults = append(articleResults, &Article{
						ID:            article.ID,
						Title:         article.Title,
						Desc:          article.Desc,
						Content:       article.Content,
						CoverImageUrl: article.CoverImageUrl,
						State:         article.State,
						Tag:           tag,
					})

					continue
				}
			}
		}
	}

	return articleResults, len(articleIDs), nil
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(
		param.Title,
		param.Desc,
		param.Content,
		param.CoverImageUrl,
		param.State,
	)
	if err != nil {
		return err
	}

	err = svc.dao.CreateArticleTag(article.ID, param.TagID)
	if err != nil {
		return err
	}

	return nil
}
