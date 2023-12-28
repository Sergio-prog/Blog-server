package index

import (
	_db "blog/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	posts := _db.GetAllPosts()

	user, _ := c.Get("user")
	isLogin := user != _db.NilUser && user != nil

	log.Println(user, isLogin)

	c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts, "logged": isLogin, "user": user})
}
