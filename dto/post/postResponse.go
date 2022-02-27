package dtopost

import (
	"time"
)

type PostResponse struct {
	Id                   string
	EditedAt             time.Time
	CreatedAt            time.Time
	DiscussionType       string
	IsArchived           string
	IsHidden             bool
	IsSaved              bool
	IsScoreHidden        bool
	IsVisited            bool
	Score                int
	SuggestedCommentSort string
	Title                string
	Url                  string
	VoteState            string
	AuthorInfoResponse   string
	AuthorOnlyInfo       string
	CommentCount         int
	content              string
	subreddit            string
}
type AuthorInfoResponse struct {
	TypeName        string
	Id              string
	Name            string
	IsPremiumMember bool
}
type SubredditResponse struct {
	Id           int
	Styles       string
	Name         string
	Subscribers  int
	Title        string
	Type         string
	Path         string
	IsFavorite   bool
	IsNSFW       bool
	IsSubscribed bool
	IsEnabled    bool
}
type StylesResponse struct {
	LegacyIcon   string
	PrimaryColor string
	Icon         string
}
