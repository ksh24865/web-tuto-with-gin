package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) ShowIndexPage(c *gin.Context) {
	articles, err := h.mauc.GetAllArticles()
	if err != nil {
		articles = nil
	}

	// Current articles include 'Writer.Password'.
	// Since this is information to be hidden in the client, you need to use an adapter,
	// but it is an example code, so I will skip it.

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
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

	tempID, _ := c.Get("uid")
	uid := tempID.(int)

	if _, err := h.mauc.CreateNewArticle(title, content, uid); err == nil {
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
