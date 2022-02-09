package domain

func GetAllPosts() []Post {
	db := GormDbConnections()
	var posts []Post
	db.Find(&posts)
	return posts
}
