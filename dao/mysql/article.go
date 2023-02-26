package mysql

import (
	"go-blog/api"
	"go-blog/model"
)

// QueryArticleList 查询所有的文章
func QueryArticleList(pageNum int64, pageSize int64) (articleList []*model.Article, err error) {
	sqlStr := `select article_id, title, content, cover_image, author_id, create_time
			   from article
			   order by create_time desc
			   limit ?, ?`
	err = db.Select(&articleList, sqlStr, (pageNum-1)*pageSize, pageSize)
	return
}

// CreateArticle 创建文章
func CreateArticle(param *api.PostArticle) (err error) {
	sqlStr := `insert into article(article_id, title, content, cover_image, author_id)
			   values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr,
		param.ArticleID,
		param.Title,
		param.Content,
		param.CoverImage,
		param.AuthorID,
	)
	return
}

// DeleteArticle 删除文章
func DeleteArticle(articleID int64) (err error) {
	sqlStr := `delete from article
			   where article_id = ?`
	_, err = db.Exec(sqlStr, articleID)
	return
}
