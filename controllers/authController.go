package controllers

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nauish/go-auth-docker/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON format",
			})
			return
		}

		if len(user.Username) == 0 || len(user.Password) == 0 || len(user.Email) == 0 {
			c.JSON(400, gin.H{
				"message": "All fields are required",
			})
			return
		}

		err := models.Client.First(&models.User{}).Where("username = ?", user.Username).Error
		if err == nil {
			c.JSON(400, gin.H{
				"message": "Username already taken",
			})
			return
		}

		err = models.Client.First(&models.User{}).Where("email = ?", user.Email).Error
		if err == nil {
			c.JSON(400, gin.H{
				"message": "Email already taken",
			})
			return
		}

		user.Password = HashPassword(user.Password)
		models.Client.Create(&user)

		if models.Client.Error != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Sign up successfully",
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid JSON format",
			})
			return
		}

		var user models.User

		err := models.Client.First(&user, "username = ?", credentials.Username).Error
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid username or password",
			})
			return
		}

		if !VerifyPassword(credentials.Password, user.Password) {
			c.JSON(400, gin.H{
				"message": "Invalid username or password",
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		JwtSecret := os.Getenv("JWT_SECRET")
		tokenString, err := token.SignedString([]byte(JwtSecret))
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Login successfully",
			"token":   tokenString,
		})
	}
}

func HashPassword(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hashed)
}

func VerifyPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
