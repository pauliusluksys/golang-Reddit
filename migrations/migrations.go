package migrations

import (
	"fmt"
	"github.com/pauliusluksys/golang-Reddit/domain"
)

var schema = `
    create table post_comments(post_comments_id bigint unsigned not null 
        auto_increment primary key,
        post_id bigint unsigned not null references posts(id),
        author_id bigint unsigned not null references users(id),
        parent_id bigint unsigned references post_comments(post_comments_id),
        text text, created_at datetime, updated_at datetime, deleted_at datetime);`

func PostMigration() {
	db := domain.GormDbConnections()
	err := db.AutoMigrate(&domain.PostGorm{}, &domain.PostCategoryGorm{}, &domain.UserGorm{})
	if err != nil {
		fmt.Println("Error while trying to migrate posts: " + err.Error())
	}
}
func PostCommentMigration() {
	db := domain.SqlxDbConnections()
	db.MustExec(schema)
}
