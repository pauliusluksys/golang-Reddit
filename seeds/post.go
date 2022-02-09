package seeds

import (
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

type userIds struct {
	ID uint `db:"id"`
}
type postCategoryIds struct {
	ID uint `db:"id"`
}

func (s Seed) PostSeed() {

	postIds := []postCategoryIds{}
	userIds := []userIds{}
	err := s.db.Select(&postIds, "SELECT id FROM categories")
	if err != nil {
		panic("error in post seed while selecting post category: " + err.Error())
	}
	err = s.db.Select(&userIds, "SELECT id FROM users")
	if err != nil {
		panic("error in post seed while selecting user: " + err.Error())
	}
	userscount := len(userIds)
	postCatCount := len(postIds)
	for i := 0; i < 100; i++ {
		randomUserId := userIds[rand.Intn(userscount)].ID
		randomPostCatId := postIds[rand.Intn(postCatCount)].ID

		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO posts(author_id,created_at,is_draft,discussion_type,title,category_id,content,is_nfs) VALUES (?,?,?,?,?,?,?,?)`)
		// execute query

		_, err := stmt.Exec(randomUserId, faker.Date(), false, "comments", faker.Sentence(), randomPostCatId, faker.Paragraph(), false)
		if err != nil {
			panic(err)
		}
	}
}
