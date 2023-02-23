package service

import (
	"go-blog/dao/mysql"
	"go-blog/model"
)

// QueryArticleList 查询文章列表
func QueryArticleList() ([]*model.Article, error) {
	return mysql.QueryArticleList()
}
