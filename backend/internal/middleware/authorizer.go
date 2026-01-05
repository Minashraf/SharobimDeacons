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
			err := errors.New("missing Authorization header")
			_ = c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			err := errors.New("invalid Authorization header format")
			_ = c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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
				_ = c.Error(err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
				return
			}
			_ = c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		if !token.Valid {
			_ = c.Error(err)
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
			err := errors.New("roles not found in context")
			_ = c.Error(err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		roles, ok := rolesValue.([]string)
		if !ok {
			err := errors.New("invalid roles type in context")
			_ = c.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, role := range allowedRoles {
			if slices.Contains(roles, role) {
				c.Next()
				return
			}
		}
		err := errors.New("unauthorized to access this resource")
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
}
