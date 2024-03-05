package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nauish/go-auth-docker/models"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		var user models.User
		result := models.Client.First(&user, "id = ?", userID)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(404, gin.H{"error": "User not found"})
			return
		} else if result.Error != nil {
			ctx.JSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		userInfo := UserInfo{
			ID:        int(user.ID),
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
		}

		ctx.JSON(200, userInfo)
	}
}
