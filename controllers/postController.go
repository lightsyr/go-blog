package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"syrlight/go-blog/models"
	"syrlight/go-blog/utils"

	"github.com/gomarkdown/markdown"
	"github.com/labstack/echo/v4"
)

func PostIndex(c echo.Context) error {

	results := models.ListAllPosts()

	for i, post := range results.Posts {

		content := markdown.ToHTML([]byte(post.Content), nil, nil)

		results.Posts[i].Content = template.HTML(content)

	}

	return c.Render(http.StatusOK, "index.html", results)
}

func CreatePost(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	file, err := c.FormFile("file")

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Required fields not present in requrest",
		})
	}

	fileName, err := utils.UploadFile("./images", file)

	if err != nil {
		fmt.Println(err)
	}

	if len(title) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Required fields not present in requrest",
		})
	}

	if len(content) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Required fields not present in requrest",
		})
	}

	models.InsertOnePost(title, content, true, fileName)

	return c.Redirect(302, "/")
}
