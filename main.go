package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env")
	}

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)

	port := os.Getenv("PORT")
	router.Run(":"+port)
}

func rootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Andromeda",
		"bio":  "Software Engineer",
	})
}

func helloHandler(ctx *gin.Context) {
	ctx.Writer.Write([]byte("Hello All"))
}
