package main

import (
	"log"
	"net/http"
	"os"
	"test-deploy-to-railway/book"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env")
	}

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&book.Book{}) // create books table

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)

	port := os.Getenv("PORT")
	router.Run(":" + port)
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
