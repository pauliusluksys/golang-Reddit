package domain

func GetAllPosts() []PostGorm {
	db := GormDbConnections()
	var posts []PostGorm
	db.Find(&posts)
	return posts
}
