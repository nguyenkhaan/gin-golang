package controller

import (
	"cloudian/cloudian-restful/internal/dto"

	"github.com/go-fuego/fuego"
)

func GetAllUsers(ctx fuego.ContextWithParams[dto.GetAllUserQuery]) (map[string]any , error) {
	limit := ctx.QueryParam("limit") 
	offset := ctx.QueryParam("offset")  
	println(limit) 
	println(offset) 
	//Hoc co the duc toasn bo struct params thong qua ctx.Params() 
	params, err := ctx.Params() 
	println(params) 
	println(err) 
	return map[string]any{
		"message": "Cloudian Love Cloud", 
	}, nil
}

//ContextWithParams, ContextWithBody, ContextWithNoBody -> Parameter, Body (JSON), NoBody (No JSON)
func GetDetailUser(ctx fuego.ContextWithParams[dto.GetDetailUserParam]) (map[string]any , error) {
	return map[string]any{
		"message": "Cloudian Love Cloud", 
	}, nil
}

func CreateUser(ctx fuego.ContextWithBody[map[string]any]) (map[string]any , error) {
	//QueryParam, PathParam, Body 
	body, err := ctx.Body() 
	println(body) 
	println(err) 
	return body, nil 
}