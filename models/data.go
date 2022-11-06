package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test comment2",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ArticleID:   1,
		Title:       "first article",
		Contents:    "This is the test article.",
		UserName:    "John Doe",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ArticleID: 2,
		Title:     "second article",
		Contents:  "This is the test article.",
		UserName:  "John Smith",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
