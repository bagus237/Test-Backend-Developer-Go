package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/server"
)

// AuthMiddleware adalah middleware untuk otentikasi OAuth2
func AuthMiddleware(oauthServer *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cek header Authorization
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header kosong"})
			c.Abort()
			return
		}

		// Parse token dari header Authorization
		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid"})
			c.Abort()
			return
		}

		// Verifikasi token dengan OAuth2 server
		token := parts[1]
		if !isValidToken(oauthServer, token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		// Token valid, lanjutkan ke handler
		c.Next()
	}
}

// isValidToken untuk verifikasi token
func isValidToken(oauthServer *server.Server, token string) bool {
	ctx := context.Background()
	_, err := oauthServer.Manager.LoadAccessToken(ctx, token)
	return err == nil
}

// Middleware untuk menangani verifikasi token yang gagal
func HandleTokenVerificationError() gin.HandlerFunc {
	return func(c *gin.Context) {
		ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
		if !exists || ti == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Next()
	}
}
