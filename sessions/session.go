package sessions

import (
	"github.com/gorilla/sessions"
)

// cookieStore function creates a cookie store with the user's secret key
func CookieStore() *sessions.CookieStore {
    SecretKey := []byte("super-secret-SecretKey")
    cookieStore := sessions.NewCookieStore(SecretKey)

    // function returns the cookie store so other functions can access it
    return cookieStore
}
