package domain

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"strconv"
)

type PostCommentGormRepo struct {
	client *gorm.DB
}
type PostCommentSqlxRepo struct {
	client *sqlx.DB
}

func NewPostCommentSqlxRepo(db *sqlx.DB) PostCommentSqlxRepo {
	return PostCommentSqlxRepo{client: db}
}
func (pcsqlx PostSqlxRepo) NewPostComment(newPostComment PostComment) PostComment {
	//query := `INSERT INTO post_comments (post_id,author_id,parent_id,content,created_at) VALUES (:postId,:AuthorId,:ParentId.Int64,:Text)`
	tx := pcsqlx.client.MustBegin()
	var result sql.Result
	if newPostComment.ParentId.Valid {
		result = tx.MustExec("INSERT INTO post_comments (post_id,author_id,parent_id,content,created_at) VALUES (?,?,?,?,?)", newPostComment.PostId, newPostComment.AuthorId, newPostComment.ParentId, newPostComment.Content, newPostComment.CreatedAt.Time)
	} else {
		result = tx.MustExec("INSERT INTO post_comments (post_id,author_id,content,created_at) VALUES (?,?,?,?)", newPostComment.PostId, newPostComment.AuthorId, newPostComment.Content, newPostComment.CreatedAt.Time)
	}
	err := tx.Commit()
	if err != nil {
		pcsqlx.Logger.Debug("err: ", err.Error())
	}
	//_, err := SqlxDbConnections().NamedExec(query, &newPostComment)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	postCommentId, err := result.LastInsertId()
	var insertedPostcomment PostComment
	err = pcsqlx.client.Get(&insertedPostcomment, "SELECT * FROM post_comments WHERE post_comments_id=?", postCommentId)
	if err != nil {
		pcsqlx.Logger.Debug(err.Error())
	}
	return insertedPostcomment
}
func (pcsqlx PostSqlxRepo) GetTotalComments(URLParams map[string]string) int {
	var postCommentsCount []int
	postIdInt, err := strconv.Atoi(URLParams["postId"])
	if err != nil {
		pcsqlx.Logger.Debug(err.Error())
	}
	query := `SELECT COUNT(post_comments_id) AS comments_count FROM post_comments where post_id = ?;`

	err = pcsqlx.client.Select(&postCommentsCount, query, postIdInt)
	if err != nil {
		pcsqlx.Logger.Debug("error while querying post comments: " + err.Error())
	}
	return postCommentsCount[0]
}
func (pcsqlx PostSqlxRepo) GetPostComments(URLParams map[string]string) []PostComment {
	postComments := []PostComment{}
	postIdInt, err := strconv.Atoi(URLParams["postId"])
	if err != nil {
		pcsqlx.Logger.Debug(err.Error())
	}
	var next int
	queryOffset := 0
	if val, ok := URLParams["next"]; ok {
		next, err = strconv.Atoi(val)
		queryOffset = next * 20
	}
	if err != nil {
		pcsqlx.Logger.Debug(err.Error())
	}

	query := `WITH RECURSIVE post_comments_path (post_comments_id,content, parent_id, post_id,author_id,created_at,updated_at,deleted_at ) AS
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

	err = pcsqlx.client.Select(&postComments, query, postIdInt, queryOffset)
	if err != nil {
		pcsqlx.Logger.Debug("error while querying post comments: " + err.Error())
	}
	return postComments
}
