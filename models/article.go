package models

import "time"

type Article struct {
	ID		string
	Title	string
	Author	string
	Content string
}

type Comment struct {
	ArticleID 	string
	Content		string
	Timestamp	time.Time
}
