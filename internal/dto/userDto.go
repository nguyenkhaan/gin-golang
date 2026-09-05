package dto

type GetAllUserQuery struct {
	Limit int `form:"limit" binding:"required,gte=0,numeric"`
	Offset int `form:"offset" binding:"required,gte=0,numeric"`
}

type GetDetailUserParam struct {
	UserID int `uri:"userID" binding:"required,numeric"`
}

type CreateUserDto struct {
	Username string `json:"username" binding:"required"` 
	Password string `json:"password" binding:"required"`
}