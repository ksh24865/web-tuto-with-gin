package main

import (
	"github.com/KumKeeHyun/web-tuto-with-gin/dataservice/mysql"
	"github.com/KumKeeHyun/web-tuto-with-gin/rest/handler"
	"github.com/KumKeeHyun/web-tuto-with-gin/rest/middleware"
	"github.com/KumKeeHyun/web-tuto-with-gin/usecase/manageArticle"
	"github.com/KumKeeHyun/web-tuto-with-gin/usecase/registration"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysql.Setup()

	//ar := memory.NewArticleRepo()
	ar := mysql.NewArticleRepo()
	ur := mysql.NewUserRepo()
	mauc := manageArticle.NewManageArticleUsecase(ar, ur)
	ruc := registration.NewRegistrationUsecase(ur)
	h := handler.NewGinHandler(mauc, ruc)

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.Use(middleware.SetUserStatus())
	loggedIn := middleware.EnsureLoggedIn()
	notLoggedIn := middleware.EnsureNotLoggedIn()

	r.GET("/", h.ShowIndexPage)
	article := r.Group("/article")
	{
		article.GET("/view/:article_id", h.ShowArticle)
		article.GET("/create", loggedIn, h.ShowArticleCreationPage)
		article.POST("/create", loggedIn, h.NewArticle)

		// 메소드는 DELETE가 되어야 하지만 html의 한계로 GET으로 대체함.
		article.GET("/delete/:article_id", h.RemoveArticle)
		//article.DELETE("/delete/:article_id", h.RemoveArticle)
	}
	user := r.Group("/u")
	{
		user.GET("/login", notLoggedIn, h.ShowLoginPage)
		user.POST("/login", notLoggedIn, h.Login)
		user.GET("/logout", loggedIn, h.Logout)
		user.GET("/register", notLoggedIn, h.ShowRegistrationPage)
		user.POST("/register", notLoggedIn, h.Register)
	}

	r.Run(":8080")
}
