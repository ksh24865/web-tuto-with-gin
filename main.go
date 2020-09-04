package main

import (
	"github.com/KumKeeHyun/web-tuto-with-gin/handler"
	"github.com/KumKeeHyun/web-tuto-with-gin/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.Use(middleware.SetUserStatus())
	loggedIn := middleware.EnsureLoggedIn()
	notLoggedIn := middleware.EnsureNotLoggedIn()

	r.GET("/", handler.ShowIndexPage)
	user := r.Group("/u")
	{
		user.GET("/login", notLoggedIn, handler.ShowLoginPage)
		user.POST("/login", notLoggedIn, handler.PerformLogin)
		user.GET("/logout", loggedIn, handler.Logout)
		user.GET("/register", notLoggedIn, handler.ShowRegistrationPage)
		user.POST("/register", notLoggedIn, handler.Register)
	}
	article := r.Group("/article")
	{
		article.GET("/view/:article_id", handler.GetArticle)
		article.GET("/create", loggedIn, handler.ShowArticleCreationPage)
		article.POST("/create", loggedIn, handler.CreateArticle)
	}

	r.Run(":8080")
}
