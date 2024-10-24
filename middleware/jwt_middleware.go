package middleware

import (
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "strings"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the Authorization header
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Request does not contain an access token"})
            c.Abort()
            return
        }

        // Extract the token from the "Bearer" scheme
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        // Get the JWT secret from the environment
        jwtSecret := os.Getenv("JWT_SECRET")
        if jwtSecret == "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret is not configured"})
            c.Abort()
            return
        }

        // Parse and validate the token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(jwtSecret), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Proceed to the next middleware/handler if the token is valid
        c.Next()
    }
}
