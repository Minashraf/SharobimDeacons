package middleware

import (
	"Backend/internal/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"slices"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

		claims := &utils.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("PRIVATE_KEY")), nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("userId", claims.UserId)
		c.Set("roles", claims.Roles)
		c.Next()
	}
}

func AllowedRoles(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesValue, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "roles not found in context"})
			return
		}

		roles, ok := rolesValue.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "invalid roles type in context"})
			return
		}
		for _, role := range allowedRoles {
			if slices.Contains(roles, role) {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to access this resource"})
	}
}
