package model

import "time"

type Article struct {
	ArticleID  int64     `db:"article_id"`
	AuthorID   int64     `db:"author_id"`
	Title      string    `db:"title"`
	Content    string    `db:"content"`
	CoverImage string    `db:"cover_image"`
	CreateTime time.Time `db:"create_time"`
}
