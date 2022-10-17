package controller

import (
	"fmt"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}

	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"user_name": User.UserName,
		"age":       User.Age,
	})
}

func LoginUser(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}

	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	password := User.Password
	err := db.Debug().Where("email = ? ", User.Email).Take(&User).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": fmt.Sprintf("Email Not registered :%s", err.Error()),
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": "Wrong password",
		})
		return
	}

	token, err := helpers.GenerateToken(User.ID, User.Email, User.UserName, User.Age)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": fmt.Sprintf("Error generating token :%s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUserData(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("userId")
	User := models.User{}

	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	err := db.Model(&User).Where("id=?", id).Updates(models.User{
		Email:    User.Email,
		UserName: User.UserName,
		Age:      User.Age,
	}).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         id,
		"email":      User.Email,
		"user_name":  User.UserName,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}

func DeleteUserAccount(ctx *gin.Context) {
	db := database.GetDB()
	User := models.User{}
	userID := ctx.Param("userId")

	db.Delete(&User, userID)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})

}
