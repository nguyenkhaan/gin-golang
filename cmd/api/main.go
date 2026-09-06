package main

import (
	"cloudian/cloudian-restful/api"

	"github.com/go-fuego/fuego"
)

// @title           Todo API
// @version         1.0
// @description     Todo API using Gin
// @host            localhost:8000
// @BasePath        /api
func main() {
	s := fuego.NewServer() 
	//Tien hanh Register lai route vao  
	fuego.Get(s , "/health" , func(ctx fuego.ContextNoBody) (map[string]any , error) {
		return map[string]any{
			"message": "Your app is running", 
			"notification": "Build with Cloudian Love Cloud", 
		}, nil 
	})
	api.RegisterRouter(s) 
	s.Run() 
}