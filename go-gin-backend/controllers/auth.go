package controllers

import (
	"bookweb/models"
	"fmt"
	"net/http"
	"bookweb/utils/token"
	"github.com/gin-gonic/gin"
)

// Register input validation
type RegisterInput struct {
	Username  string `json:"username" binding:"required"`
	Password1 string `json:"password1" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

// Register controller
func Register(ctx *gin.Context) {
	var input RegisterInput
	var user models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if input.Password1 != input.Password2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not match",
		})
		return
	}
	exist := models.DB.Where("username = ?", input.Username).First(&user).Error
	if exist == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password1
	_, err := u.SaveUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "registration success",
	})
}

// Login Input validation
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login controller
func Login(ctx *gin.Context) {
	var input LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	fmt.Println(u.Username,u.Password)
	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or password incorrect",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// Retieve currentuser.
func CurrentUser(ctx *gin.Context){
	user_id, err := token.ExtractTokenID(ctx)

	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	user,err := models.GetUserByID(user_id)
	if err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"success","data":user})

}

