package service

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/dao"
)

type Service struct {
	dao *dao.Dao
}

func New() Service {
	svc := Service{}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
