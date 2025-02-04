package middleware

import (
	"net/http"

	. "backend/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(reqRole string) gin.HandlerFunc{
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		role := session.Get("role")

		if username == nil || role == nil {
			c.JSON(http.StatusUnauthorized, Message("silahkan login terlebih dahulu"))
			c.Abort()
			return
		}
		if role != reqRole{
			c.JSON(http.StatusForbidden, Message("user tidak diizinkan"))
			c.Abort()
			return
		}
		c.Next()
	}
}