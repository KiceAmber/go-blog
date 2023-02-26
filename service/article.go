package service

import (
	"go-blog/api"
	"go-blog/dao/mysql"
	"go-blog/model"
	"go-blog/pkg/snowflake"
)

// QueryArticleList 查询文章列表
func QueryArticleList(pageNum int64, pageSize int64) ([]*model.Article, error) {
	return mysql.QueryArticleList(pageNum, pageSize)
}

// CreateArticle 创建文章
func CreateArticle(param *api.PostArticle) (err error) {
	// 生成文章ID
	ArticleID := snowflake.GenID()
	param.ArticleID = ArticleID
	return mysql.CreateArticle(param)
}

// DeleteArticle 删除文章
func DeleteArticle(articleID int64) (err error) {
	return mysql.DeleteArticle(articleID)
}
