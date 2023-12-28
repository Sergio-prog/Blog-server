package middleware

import (
	"blog/sessions"
	"blog/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем сессию
		session, err := sessions.CookieStore().Get(c.Request, "session-name")
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Проверяем наличие информации о пользователе в сессии
		userID, exist := session.Values["user_id"]
		if !exist || userID == nil {
			c.Set("user", nil)
		} else {
			// Устанавливаем объект пользователя в контекст
			c.Set("user", _db.GetUserById(userID.(uint)))
		}

		// Продолжаем выполнение запроса
		c.Next()
	}
}
