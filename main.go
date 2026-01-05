package main

import (
	"time"
)

type PostLikedEvent struct {
	PostID string
	Liker  string // userID
	at     time.Time
}

func (e PostLikedEvent) Name() string {
	return "PostLiked"
}

func (e PostLikedEvent) Time() time.Time {
	return e.at
}

type Comment struct {
	commentId  string
	auther     User
	content    string
	timestamp  string
	likes      []User
	parentPost Post
}
type TimeLine struct {
	owner User
	posts []Post
}

func main() {

}
