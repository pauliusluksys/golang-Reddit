package domain

import (
	"fmt"
	"strconv"
)

func GetTotalComments(URLParams map[string]string) int {
	var postCommentsCount []int
	postIdInt, err := strconv.Atoi(URLParams["postId"])
	if err != nil {
		fmt.Println(err.Error())
	}
	query := `SELECT COUNT(post_comments_id) AS comments_count FROM post_comments where post_id = ?;`

	err = SqlxDbConnections().Select(&postCommentsCount, query, postIdInt)
	if err != nil {
		fmt.Println("error while querying post comments: " + err.Error())
	}
	return postCommentsCount[0]
}
func GetPostComments(URLParams map[string]string) []PostComment {
	postComments := []PostComment{}
	postIdInt, err := strconv.Atoi(URLParams["postId"])
	if err != nil {
		fmt.Println(err.Error())
	}
	var next int
	queryOffset := 0
	if val, ok := URLParams["next"]; ok {
		next, err = strconv.Atoi(val)
		queryOffset = next * 20
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	query := `WITH RECURSIVE post_comments_path (post_comments_id,content, parent_id, post_id,user_id,created_at,updated_at,deleted_at ) AS
(
  SELECT post_comments_id,content, parent_id, post_id,author_id,created_at,updated_at,deleted_at 
    FROM post_comments
    WHERE post_comments_id IN (SELECT * FROM (SELECT 
            post_comments_id 
        FROM
            post_comments pc 
        WHERE
           pc.parent_id is null)as pc) and post_id = ?
  UNION ALL
  SELECT c.post_comments_id,c.content, c.parent_id, c.post_id,c.author_id,c.created_at,c.updated_at,c.deleted_at
    FROM post_comments_path AS cp JOIN post_comments AS c
      ON cp.post_comments_id = c.parent_id 
)
SELECT * FROM post_comments_path 
ORDER BY parent_id,created_at DESC limit 100 offset ?;`

	err = SqlxDbConnections().Select(&postComments, query, postIdInt, queryOffset)
	if err != nil {
		fmt.Println("error while querying post comments: " + err.Error())
	}
	return postComments
}
