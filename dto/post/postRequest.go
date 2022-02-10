package dtopost

import "time"

type NewPostRequest struct {
	Id             string
	CreatedAt      time.Time
	DiscussionType string
	Title          string
	AuthorInfo     string
	Subreddit      string
	Content        string
	IsNFS          bool
}
