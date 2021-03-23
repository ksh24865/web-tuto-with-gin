package handler

import (
	"log"

	"github.com/web-tuto-with-gin/usecase"
	"github.com/gin-gonic/gin"
)

func catchPanic() {
	if p := recover(); p != nil {
		log.Printf("%+v\n", p)
	}
}

type GinHandler struct {
	mauc usecase.ManageArticleUsecase
	ruc  usecase.RegistrationUsecase
}

func NewGinHandler(mauc usecase.ManageArticleUsecase, ruc usecase.RegistrationUsecase) *GinHandler {
	return &GinHandler{
		mauc: mauc,
		ruc:  ruc,
	}
}

func (h *GinHandler) ShowArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func (h *GinHandler) ShowLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func (h *GinHandler) ShowRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "register.html")
}
