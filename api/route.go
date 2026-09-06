package api

import (
	"cloudian/cloudian-restful/internal/controller"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

func RegisterRouter(f *fuego.Server) {
	api := fuego.Group(f , "/api") 
	user := fuego.Group(api , "/user" , option.Tags("User")) 
	fuego.Get(user , "/" , controller.GetAllUsers) 
}