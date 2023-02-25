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
