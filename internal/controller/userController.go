package controller

import (
	"cloudian/cloudian-restful/internal/dto"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	var query dto.GetAllUserQuery 
	//Su dung dto de co the thuc hien Biding du lieu truyen den 
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(), 
		})
		return 
	}
	ctx.JSON(200 , gin.H{
		"message": "Get all users", 
	})
}

func GetDetailUser(ctx *gin.Context) {
	
	var param dto.GetDetailUserParam 
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(), 
		})
		return 
	}
	ctx.JSON(200, gin.H{
		"message": "List user detail information", 
	})
} 

func CreateUser(ctx *gin.Context) {
	var request dto.CreateUserDto 
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(400 , gin.H{
			"error": err.Error(), 
		}) 
		return 
	}
	ctx.JSON(201, gin.H{
		"message": "Create a new user", 
	})
}

func UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Update user information", 
	})
}

func DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete user successfully", 
	})
}
