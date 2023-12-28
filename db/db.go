package _db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// type Post struct {
// 	gorm.Model
// 	id        int
// 	author_id int
// 	created   string
// 	title     string
// 	body      string
// }

type PostWithUser struct {
	gorm.Model
	id        int
	author_id int
	created   string
	title     string
	body      string
	username  string
}

// Структура для таблицы "user"
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

var NilUser = User{}

// Структура для таблицы "post"
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	AuthorID  uint      `gorm:"not null"`
	Created   time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	Title     string    `gorm:"not null"`
	Body      string    `gorm:"not null"`
	Author    User      `gorm:"foreignKey:AuthorID"`
}

func initializeDB() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Post{}, &User{})

	// sqlStmt, err := os.ReadFile("D:\\GoProjects\\blog-server\\schemas\\schema.sql")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = db.Exec(string(sqlStmt))
	// if err != nil {
	// 	panic(err)
	// }

}

func init() {
	// initializeDB()
}


func GetAllPosts() []Post {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	var posts []Post
	db.Table("post").Select("post.id, title, body, created, author_id, username").
	Joins("JOIN user ON post.author_id = user.id").
	Order("created DESC").
	Scan(&posts)

	// post_json, err := json.Marshal(posts)
	// if err != nil {
	// 	panic(err)
	// }

	// result := []map[string]any{}
	// for i := range posts {
	// 	result = append(result, structs.Map(i))
	// }

	return posts

}

func AddNewUser(username, password string) {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.Table("user").Create(&User{Username: username, Password: password}).Commit()
}


func GetUser(username string) User {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	var user User
	// err = db.Table("user").First(&user, &User{Username: username}).Error
	err = db.Table("user").Where("username = ?", string(username)).First(&user).Error
	if err != nil {
		return User{}
	}

	return user
}

func GetUserById(id uint) User {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	var user User
	db.Table("user").First(&user, &User{ID: id})
	// db.Where("username = ?", string(username)).First(&user)

	return user
}
