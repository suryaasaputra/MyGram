package middleware

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		id, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"error":   "Bad request",
				"message": "Invalid ID",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))
		user := models.User{}
		err = db.First(&user, id).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if userID != user.ID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		id, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"error":   "Bad request",
				"message": "Invalid ID",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))
		photo := models.Photo{}
		err = db.Select("user_id").First(&photo, id).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if userID != photo.UserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		id, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"error":   "Bad request",
				"message": "Invalid ID",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))
		comment := models.Comment{}
		err = db.Select("user_id").First(&comment, id).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if userID != comment.UserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}

func SosmedAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		id, err := strconv.Atoi(ctx.Param("socialMediaId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"error":   "Bad request",
				"message": "Invalid ID",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))
		sosmed := models.SocialMedia{}
		err = db.Select("user_id").First(&sosmed, id).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Not Found",
				"message": "Data doesn't exist",
			})
			return
		}

		if userID != sosmed.UserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusNotFound,
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}
}
