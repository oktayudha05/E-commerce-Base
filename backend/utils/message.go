package utils

import (
	"github.com/gin-gonic/gin"
)

func Message(pesan string, data... any) gin.H {
	if len(data) > 0 {
		return gin.H{"message": pesan, "data": data[0]}
	}
	return gin.H{"message": pesan}
}