package main

import (
	"github.com/djzwieback/go-demo/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	defineRestInterfaces(r)
	err := r.Run()
	if err != nil {
		log.Fatal("There was an error starting the server", err);
	}
}

func defineRestInterfaces(ginEngine *gin.Engine) {

	ginEngine.GET("/ping", controller.GetPing)
}
