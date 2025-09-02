package controllers

import (
	"os"
	"time"

	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	//get email and password from req body
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		//return so it stops here
		return
	}

	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to hash password"})
	}

	//create user
	user := models.User{Email: body.Email, Password: string(hashedPassword)}

	result := initializers.DB.Create(&user)

	//check if we have error
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Failed to create user"})
		return
	}

	//return response
	c.JSON(200, gin.H{"message": "User created successfully"})

}

func Login(c *gin.Context) {
	//get email and password from req
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		//return so it stops here
		return
	}
	//get user with that email
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	//check if it found the user
	if user.ID == 0 {
		c.JSON(404, gin.H{"error": "User does not exist"})
		return
	}

	//compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	//generate jwt token and return

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	//convert the secret to byte
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to sign token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})

}
