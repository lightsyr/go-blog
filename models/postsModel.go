package models

import (
	"database/sql"
	"fmt"
	"html/template"
	"syrlight/go-blog/config"
	"time"
)

type Post struct {
	Id         int           `json:"id"`
	Title      string        `json:"title"`
	Content    template.HTML `json:"content"`
	Avaliable  bool          `json:"avaliable"`
	Created_at time.Time     `json:"created_at"`
	Updated_at time.Time     `json:"updated_at"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

func CreatePost(*Post, *sql.DB) {
	//TODO

}

func ListAllPosts() Posts {

	rows, err := config.DB.Query("SELECT * FROM posts")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Posts{}

	for rows.Next() {
		post := Post{}

		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Avaliable, &post.Created_at, &post.Updated_at)

		if err != nil {
			fmt.Println(err)
		}

		result.Posts = append(result.Posts, post)
	}

	return result

}
