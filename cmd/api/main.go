package main

import (
	"cloudian/cloudian-restful/api" 
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default() 
	api.RegisterRouter(r)  
	r.Run(":8000")
}