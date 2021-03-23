package handler

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/web-tuto-with-gin/domain/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *GinHandler) Login(c *gin.Context) {
	session := sessions.Default(c)

	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	var (
		u   *model.User
		err error
	)
	if u, err = h.ruc.MatchUser(username, password); err == nil {
		token := generateSessionToken()
		session.Set(token, u.ID)
		if err = session.Save(); err == nil {
			// TODO : This is example code for using session, cookie
			// It is not actually used like this in the public
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)
			render(c, gin.H{
				"title": "Successful Login"}, "login-successful.html")
		}
	}
	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		"ErrorTitle":   "Login Failed",
		"ErrorMessage": err.Error()})
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// TODO : using jwt or something..
	return strconv.FormatInt(rand.Int63(), 16)
}

func (h *GinHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)

	tempToken, _ := c.Get("token")
	token := tempToken.(string)

	session.Delete(token)
	if err := session.Save(); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *GinHandler) Register(c *gin.Context) {
	session := sessions.Default(c)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if u, err := h.ruc.RegisterUser(username, password); err == nil {
		token := generateSessionToken()
		session.Set(token, u.ID)
		if err = session.Save(); err == nil {
			c.SetCookie("token", token, 3600, "", "", false, true)
			c.Set("is_logged_in", true)
			render(c, gin.H{
				"title": "Successful Login"}, "login-successful.html")
		}
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}
