package middleware

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if token, err := c.Cookie("token"); err == nil || token != "" {
			uid := session.Get(token)
			if uid != nil {
				c.Set("is_logged_in", true)
				c.Set("token", token)
				c.Set("uid", uid)

				c.Next()
			} else {
				c.SetCookie("token", "", -1, "", "", false, true)
			}
		}
		c.Set("is_logged_in", false)
	}
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's an error or if the token is empty
		// the user is not logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			//if token, err := c.Cookie("token"); err != nil || token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's no error or if the token is not empty
		// the user is already logged in
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			// if token, err := c.Cookie("token"); err == nil || token != "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
