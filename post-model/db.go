package post_model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type Post struct {
	Content   string
	GroupName string
	GroupId   int
}

func getPosts() []Post {
	db, err := sqlx.Connect("mysql", os.Getenv("DB_STRING"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var posts []Post

	err = db.Select(&posts, os.Getenv("SQL_QUERY"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return posts
}
