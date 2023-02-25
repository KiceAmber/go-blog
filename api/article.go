package api

// 发布文章接口
type PostArticle struct {
	ArticleID  int64  `json:"article_id"`
	AuthorID   int64  `json:"author_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CoverImage string `json:"cover_image"`
}
