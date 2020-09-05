package rest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/KumKeeHyun/web-tuto-with-gin/usecase"
	"github.com/gin-gonic/gin"
)

func catchPanic() {
	if p := recover(); p != nil {
		log.Printf("%+v\n", p)
	}
}

type GinHandler struct {
	mauc usecase.ManageArticleUsecase
}

func NewGinHandler(mauc usecase.ManageArticleUsecase) *GinHandler {
	return &GinHandler{
		mauc: mauc,
	}
}

func (h *GinHandler) ShowIndexPage(c *gin.Context) {
	articles := h.mauc.GetAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func (h *GinHandler) ShowArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func (h *GinHandler) ShowArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := h.mauc.GetArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (h *GinHandler) NewArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if _, err := h.mauc.CreateNewArticle(title, content); err == nil {
		// If the article is created successfully, redirect to home page
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func (h *GinHandler) RemoveArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if err := h.mauc.DeleteArticleByID(articleID); err == nil {
			// If the article is deleted successfully, redirect to home page
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			// if there was an error while deleting the article, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
