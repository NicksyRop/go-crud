package middleware

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	log.Println("Executing middleware requireAuth")
	//get token from headers
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	// Expect format: "Bearer <token>"
	fields := strings.Fields(authHeader)
	if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
		c.JSON(401, gin.H{"error": "Invalid Authorization header format"})
		c.Abort()
		return
	}
	tokenString := fields[1]

	log.Println("Token is", tokenString)

	//decode / validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {

		log.Printf("jwt parse error: %v", err)
		//check expiry
		if errors.Is(err, jwt.ErrTokenExpired) {
			c.AbortWithStatusJSON(401, gin.H{"message": "Token expired"})
			return
		}

		// other invalid token errors
		c.AbortWithStatusJSON(401, gin.H{"message": "Invalid token"})
		return

	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//find user
		sub, ok := claims["sub"].(float64)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid token subject"})
			return
		}

		var user models.User
		if err := initializers.DB.First(&user, uint(sub)).Error; err != nil {
			c.AbortWithStatusJSON(404, gin.H{"message": "User not found"})
			return
		}

		//nb you can set the user object to the request like below if required
		c.Set("user", user) //pass the value and not pointer

		//execute the pending handlers in the chain inside the calling handler / continue
		c.Next()
	} else {
		fmt.Println(err)
	}
}
