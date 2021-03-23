package main

import (
	"github.com/web-tuto-with-gin/dataservice/mysql"
	"github.com/web-tuto-with-gin/rest"
	"github.com/web-tuto-with-gin/usecase/manageArticle"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql.Setup()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	//ar := memory.NewArticleRepo()
	ar := mysql.NewArticleRepo()
	mauc := manageArticle.NewManageArticleUsecase(ar)
	h := rest.NewGinHandler(mauc)

	r.GET("/", h.ShowIndexPage)
	article := r.Group("/article")
	{
		article.GET("/view/:article_id", h.ShowArticle)
		article.GET("/create", h.ShowArticleCreationPage)
		article.POST("/create", h.NewArticle)

		// 메소드는 DELETE가 되어야 하지만 html의 한계로 GET으로 대체함.
		article.GET("/delete/:article_id", h.RemoveArticle)
		//article.DELETE("/delete/:article_id", handler.DeleteArticle)
	}

	r.Run(":8080")
}
