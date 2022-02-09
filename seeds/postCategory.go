package seeds

import (
	"github.com/bxcodec/faker/v3"
	"github.com/gosimple/slug"
)

func (s Seed) PostCategorySeed() {

	for i := 0; i < 100; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO categories(created_at,name,title,type,path) VALUES (?,?,?,?,?)`)
		// execute query
		categoryName := faker.Paragraph()
		text := slug.Make(categoryName)
		categoryUrLPath := "categories/" + text
		_, err := stmt.Exec(faker.Date(), faker.Word(), faker.Paragraph(), "cool", categoryUrLPath)
		if err != nil {
			panic(err)
		}
	}
}
