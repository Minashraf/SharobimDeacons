package middleware

import (
	"Backend/internal/consts"
	"Backend/internal/utils"
	"context"
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

		ctx := context.WithValue(c.Request.Context(), consts.UserIDKey, claims.UserId)
		ctx = context.WithValue(ctx, consts.Roles, claims.Roles)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func AllowedRoles(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesValue, exists := c.Get(string(consts.Roles))
		if !exists {
			err := errors.New("Roles not found in context")
			_ = c.Error(err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		roles, ok := rolesValue.([]string)
		if !ok {
			err := errors.New("invalid Roles type in context")
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
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}
}

func UserIDFromContext(ctx context.Context) (int64, error) {
	val := ctx.Value(consts.UserIDKey)
	if val == nil {
		return 0, errors.New("missing user id")
	}

	userID, ok := val.(int64)
	if !ok {
		return 0, errors.New("invalid user id type")
	}

	return userID, nil
}
