package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"capregsoft.com/carlocator/config"
	_ "capregsoft.com/carlocator/docs"
	"capregsoft.com/carlocator/service/server"
	gateway "github.com/apex/gateway/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func main() {
	r := gin.New()
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	if InLambda() {
		fmt.Println("running aws lambda in aws")
		log.Fatal(gateway.ListenAndServe(":5000", server.NewServerImpl(r)))
	} else {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:5000/swagger/doc.json")))
		fmt.Println("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":5000", server.NewServerImpl(r)))
	}
}
