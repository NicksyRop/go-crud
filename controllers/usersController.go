package controllers

import (
	"os"
	"time"

	"example.com/m/v2/dtos"
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// SignUp godoc
// @Summary      Register a new user
// @Description  Create a new user account with email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body dtos.SignUpRequest  true  "User signup request"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Router       /signup [post]
func SignUp(c *gin.Context) {

	//get email and password from req body
	var body dtos.SignUpRequest

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

// Login godoc
// @Summary      Login user
// @Description  Authenticate user and return JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param       request  body dtos.SignUpRequest  true  "User login request"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Router       /login [post]
func Login(c *gin.Context) {
	//get email and password from req
	var body dtos.SignUpRequest
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
