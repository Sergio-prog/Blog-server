package main

import (
	// "database/sql"
	// "log"
	"net/http"
	// "os"

	"github.com/gin-gonic/gin"
	// "github.com/urfave/cli"
	"blog/auth"
	"blog/index"
	middleware "blog/middlewares"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	r.LoadHTMLGlob("templates/*")

	r.Static("/static", "./static")

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello world"})
	})

	r.GET("/", index.Index)

	_auth := r.Group("/auth")

	_auth.GET("/login", auth.IndexLogin)
	_auth.GET("/register", auth.IndexRegister)
	_auth.GET("/logout", auth.Logout)

	_auth.POST("/login", auth.Login)
	_auth.POST("/register", auth.Register)

	return r
}

func main() {
	// if len(os.Args) >= 1 {
	// 	app := cli.NewApp()
	// 	app.Name = "Init database for server"
	// 	app.Usage = "Init database by SQL schema"

	// 	app.Commands = []cli.Command {
	// 		{
	// 			Name: "init",
	// 			HelpName: "init",
	// 			Action: initDataBase,
	// 			ArgsUsage: ``,
	// 			Usage: `Init new database`,
	// 			Description: `Init new database`,
	// 		},
	// 	}
	// 	err := app.Run(os.Args)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	r := setupRouter()
	// 	r.Run(":8080")
	// }

	// defer _db.CloseDB()

	r := setupRouter()
	r.Run(":8080")
}
