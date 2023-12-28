package auth

import (
	_db "blog/db"
	"blog/sessions"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexLogin(c *gin.Context) {
	// cols, _ := result.Columns()
	// m := make(map[string]interface{})
	// for result.Next() {
	// 	columns := make([]interface{}, len(cols))
	// 	columnPointers := make([]interface{}, len(cols))
	// 	for i, _ := range columns {
	// 		columnPointers[i] = &columns[i]
	// 	}

	// 	// Получаем данные текущей строки и добавляем их в map
	// 	err := result.Scan(columnPointers...)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for i, colName := range cols {
	//     	val := columnPointers[i].(*interface{})
	//     	m[colName] = *val
	// 	}

	// 	log.Println(m)

	// }

	c.HTML(http.StatusOK, "login.html", "")

	// for _, post := range result {
	// 	log.Println(post)
	// }
}

func IndexRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", "")
}

func IsRegister(c *gin.Context) {
	_, exist := c.Get("user")
	if exist == false {
		c.Redirect(http.StatusFound, "/auth/login")
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Println(username, password)

	var _err error
	// if !validateUsername("") {
	// 	_err = errors.New("Username is required!")
	// } else if !validatePassword("") {
	// 	_err = errors.New("Passworde is required!")
	// }

	user := _db.GetUser(username)
	log.Println(user)
	nil_user := _db.User{}

	if user == nil_user {
		_err = errors.New("Inccorect Username!")
	} else if checkPassHash(user.Password, password) == nil {
		_err = errors.New("Inccorect password!")
	}

	if _err == nil {
		// Save current session
	} else {
		c.Error(_err)
	}

	session, err := sessions.CookieStore().Get(c.Request, "session-name")
	if err != nil {
		log.Fatalln(err)
	}

	session.Values["user_id"] = user.ID
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		log.Fatalln(err)
	}

	c.Redirect(http.StatusFound, "/")
}

func checkPassHash(passwordHash, password string) any {
	return 1
}

func Register(c *gin.Context) {
	// form := c.Request.Form

	username := c.PostForm("username")
	password := c.PostForm("password")

	var _err error
	if !validateUsername(username) {
		_err = errors.New("Username is required!")
	} else if !validatePassword(password) {
		_err = errors.New("Passworde is required!")
	}

	if _err == nil {
		_db.AddNewUser(username, password)
	} else {
		c.Error(_err)
	}

	c.Redirect(http.StatusFound, "/auth/login")
}

func validateUsername(name string) bool {
	return name != "" && len(name) >= 3
}

func validatePassword(pass string) bool {
	return pass != "" && len(pass) >= 3
}

func LoadUser(c *gin.Context) {
	session, err := sessions.CookieStore().Get(c.Request, "session-name")
	if err != nil {
		log.Fatalln(err)
	}

	userId, exist := session.Values["user_id"]
	if exist == false {
		c.Set("user", nil)
	} else {
		c.Set("user", _db.GetUserById(userId.(uint)))
	}
}

func Logout(c *gin.Context) {
	c.Set("user", nil)

	session, err := sessions.CookieStore().Get(c.Request, "session-name")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	delete(session.Values, "user_id")

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Redirect(http.StatusFound, "/")
}

// func makeScanArgs(dest map[string]interface{}) []interface{} {
// 	args := make([]interface{}, 0, len(dest))
// 	for key := range dest {
// 		args = append(args, dest[key])
// 	}
// 	return args
// }
