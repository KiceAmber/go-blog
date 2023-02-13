package model

type Article struct {
	ArticleID  int64  `db:"article_id"`
	Title      string `db:"title"`
	Content    string `db:"content"`
	CoverImage string `db:"cover_image"`
	AuthorID   int64  `db:"author_id"`
}
