package blog

import "github.com/jinzhu/gorm"

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=test dbname=test password=test")
	defer db.Close()
	if err != nil {

	}
}
