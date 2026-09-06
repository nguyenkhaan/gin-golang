package dto 

type GetAllUserQuery struct {
	Limit int `query:"limit" validate:"required,gte=1,lte=100"`
	Offset int `query:"offset" validate:"required,gte=0"` //Co the su dung path:"userID"
}
type GetDetailUserParam struct {
	UserID int `path:"userID" validate:"required,numeric"`
}
type CreateUserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}