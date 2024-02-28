package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk autentikasi sederhana
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dapatkan token dari header Authorization
		token := c.GetHeader("Authorization")

		// Periksa apakah token valid (contoh: token tidak boleh kosong)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort() // Hentikan eksekusi request jika token tidak valid
			return
		}

		// Lanjutkan eksekusi request jika token valid
		c.Next()
	}
}
