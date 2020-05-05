package main

import (
	"net/http"
	"strconv"

	"github.com/Zaki_Khairi/https-service/model"
	"github.com/labstack/echo"
)

func main() {
	store := model.NewArticleStoreInMemory()

	e := echo.New()

	e.GET("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		idconv, _ := strconv.Atoi(id)
		articles := store.ArticleMap

		return c.JSON(http.StatusOK, articles[idconv-1])
	})

	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap

		return c.JSON(http.StatusOK, articles)
	})

	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("Title")
		body := c.FormValue("Body")

		article, _ := model.CreateArticle(title, body)

		store.Save(article)

		return c.JSON(http.StatusOK, article)
	})

	e.PUT("/articles/:id", func(c echo.Context) error {
		title := c.FormValue("Title")
		body := c.FormValue("Body")
		id := c.Param("id")
		idconv, _ := strconv.Atoi(id)

		articles := store.ArticleMap
		articles[idconv-1] = model.Article{ID: idconv, Title: title, Body: body}

		return c.JSON(http.StatusOK, store.ArticleMap)
	})

	e.DELETE("/articles/:id", func(c echo.Context) error {
		id := c.Param("id")
		idconv, _ := strconv.Atoi(id)
		store.ArticleMap = append(store.ArticleMap[:idconv-1], store.ArticleMap[idconv:]...)

		return c.JSON(http.StatusOK, store.ArticleMap)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
