package controller

import "github.com/gin-gonic/gin"

func GetAllUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "List all users v1", 
	})
}

func GetDetailUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "List user detail information", 
	})
} 

func CreateUser(ctx *gin.Context) {
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
