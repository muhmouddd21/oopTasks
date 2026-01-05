package domain

import "time"

type Post struct {
	postId       string
	auther       User
	content      string
	timestamp    time.Time
	likes        []User
	comments     []Comment
	isEdited     bool
	lastEditTime time.Time
}
