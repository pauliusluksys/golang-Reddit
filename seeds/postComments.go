package seeds

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
)

type postIds struct {
	ID uint `db:"id"`
}

func (s Seed) PostCommentsSeed() {

	postIds := []postIds{}
	userIds := []userIds{}
	err := s.db.Select(&postIds, "SELECT id FROM posts")
	if err != nil {
		panic("error in post seed while selecting post: " + err.Error())
	}
	err = s.db.Select(&userIds, "SELECT id FROM users")
	if err != nil {
		panic("error in post seed while selecting user: " + err.Error())
	}
	userscount := len(userIds)

	for i := 0; i < 10; i++ {

		randomUserId := userIds[rand.Intn(userscount)].ID

		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO post_comments(post_id,author_id,content,created_at) VALUES (?,?,?,?)`)
		// execute query
		fmt.Println(postIds[0].ID)
		fmt.Println(randomUserId)
		fmt.Println(faker.Sentence())
		fmt.Println(faker.Date())
		rows, err := stmt.Exec(postIds[0].ID, randomUserId, faker.Sentence(), faker.Date())
		if err != nil {
			panic(err)
		}
		fmt.Println(rows.RowsAffected())
		commentParentId, err := rows.LastInsertId()
		if err != nil {
			fmt.Println("Error during post comment insert into database: " + err.Error())
		}

		for l := 0; l < 10; l++ {
			randomUserId := userIds[rand.Intn(userscount)].ID

			//prepare the statement
			stmt, _ := s.db.Prepare(`INSERT INTO post_comments(post_id,author_id,parent_id,content,created_at) VALUES (?,?,?,?,?)`)
			// execute query

			rows, err := stmt.Exec(postIds[0].ID, randomUserId, commentParentId, faker.Sentence(), faker.Date())
			if err != nil {
				panic(err)
			}
			commentParentId, err = rows.LastInsertId()
			if err != nil {
				fmt.Println("Error during post comment insert into database: " + err.Error())
			}

			for k := 0; k < 10; k++ {
				randomUserId := userIds[rand.Intn(userscount)].ID

				//prepare the statement
				stmt, _ := s.db.Prepare(`INSERT INTO post_comments(post_id,author_id,parent_id,content,created_at) VALUES (?,?,?,?,?)`)
				// execute query

				_, err := stmt.Exec(postIds[0].ID, randomUserId, commentParentId, faker.Sentence(), faker.Date())
				if err != nil {
					panic(err)
				}
			}
		}

	}

}
