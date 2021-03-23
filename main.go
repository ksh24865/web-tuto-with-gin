package main

import (
	"github.com/web-tuto-with-gin/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/", handler.ShowIndexPage)
	article := r.Group("/article")
	{
		article.GET("/view/:article_id", handler.GetArticle)
		article.GET("/create", handler.ShowArticleCreationPage)
		article.POST("/create", handler.CreateArticle)

		// 메소드는 DELETE가 되어야 하지만 html의 한계로 GET으로 대체함. js필요
		article.GET("/delete/:article_id", handler.DeleteArticle)
		//article.DELETE("/delete/:article_id", handler.DeleteArticle)
	}

	r.Run(":8080")
}
