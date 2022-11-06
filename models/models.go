package models

import "time"

type Article struct {
	ID          int       `json:"article_id"` // 記事ID
	Title       string    `json:"title"`      // 記事タイトル
	Contents    string    `json:"contents"`   // 記事本文
	UserName    string    `json:"user_name"`  // 投稿者名
	NiceNum     int       `json:"nice"`       // いいね数
	CommentList []Comment `json:"comments"`   // 記事についたコメント(スライス形式)
	CreatedAt   time.Time `json:"created_at"` // 投稿時刻
}

type Comment struct {
	CommentID int       `json:"comment_id"` // コメントID
	ArticleID int       `json:"article_id"` // コメント投稿先の記事ID
	Message   string    `json:"message"`    // コメント本文
	CreatedAt time.Time `json:"created_at"` // 投稿時刻
}
