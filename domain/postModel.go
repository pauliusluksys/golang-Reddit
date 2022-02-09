package domain

import "database/sql"

type PostGorm struct {
	PostId         string
	IsDraft        bool
	DiscussionType string
	Title          string
	AuthorId       string
	SubredditId    string
	Content        string
	IsNFS          bool
	Media          string
	CreatedAt      sql.NullTime
	EditedAt       sql.NullTime
}

//Tabler and tableName are for rewriting default association of PostGorm struct with database table "post_gorms" to "posts"
