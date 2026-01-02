package share

import (
    "log"
    "os"
    "time"
    "strings"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func GlobalMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        // Before request
        log.Printf("Started %s %s", c.Request.Method, c.Request.URL.Path)

        c.Next()

        // After request
        latency := time.Since(t)
        status := c.Writer.Status()
        log.Printf("Completed %d in %v", status, latency)
    }
}

// GetJWTSecret retrieves the JWT secret from environment variables
func GetJWTSecret() []byte {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        log.Fatal("JWT_SECRET environment variable not set")
    }
    return []byte(secret)
}

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return GetJWTSecret(), nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }
        c.Next()
    }
}