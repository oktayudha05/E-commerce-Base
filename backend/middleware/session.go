package middleware

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func SetupSession()gin.HandlerFunc{
	store, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_URI"), os.Getenv("REDIS_PASS"), []byte(os.Getenv("REDIS_KEY")))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{
		MaxAge: 1800,
		Path: "/",
		HttpOnly: true,
		Secure: false,
	})
	return sessions.Sessions("pkm-sayur", store)
}