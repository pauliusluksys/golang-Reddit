package domain

func GetAllPosts() []PostGorm {
	db := GormDbConnections()
	var posts []PostGorm
	db.Find(&posts)
	return posts
}
func GetPostById(postId string) Post {
	post := Post{}
	query := "SELECT * FROM posts where id=$1 LIMIT 1;"
	SqlxDbConnections().Select(&post, query, postId)
	return post
}
