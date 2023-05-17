package models

import (
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
	Image      string        `json:"image"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

func InsertOnePost(title string, content string, avaliable bool, imageName string) {
	sql, err := config.DB.Prepare(`
		INSERT INTO posts 
		(title, content, avaliable, created_at, updated_at, image) 
		values (?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = sql.Exec(title, content, avaliable, time.Now(), time.Now(), imageName)

	if err != nil {
		fmt.Println(err.Error())
	}

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

		err := rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.Avaliable,
			&post.Created_at,
			&post.Updated_at,
			&post.Image,
		)

		if err != nil {
			fmt.Println(err)
		}

		result.Posts = append(result.Posts, post)
	}

	return result

}
