package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"syrlight/go-blog/models"

	"github.com/gomarkdown/markdown"
	"github.com/labstack/echo/v4"
)

func PostIndex(c echo.Context) error {

	results := models.ListAllPosts()

	for i, post := range results.Posts {

		content := markdown.ToHTML([]byte(post.Content), nil, nil)

		results.Posts[i].Content = template.HTML(content)

		fmt.Println(post.Avaliable)
	}

	return c.Render(http.StatusOK, "index.html", results)
}
