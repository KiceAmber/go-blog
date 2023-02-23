package mysql

import "go-blog/model"

// QueryArticleList 查询所有的文章
func QueryArticleList() (articleList []*model.Article, err error) {
	sqlStr := `select article_id, title, content, cover_image, author_id, create_time
			   from article`
	err = db.Select(&articleList, sqlStr)
	if err != nil {
		return
	}

	return
}
