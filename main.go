package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db kaga konek")
	}

	db.AutoMigrate(&book.Book{})

	bukuRepository := book.NewRepository(db)
	bukuService := book.NewService(bukuRepository)
	bookHandler := handler.NewBookHandler(bukuService)

	router := gin.Default()
	v1 := router.Group("/v1")

	// get
	v1.GET("/books", bookHandler.SemuaBookHandler)
	v1.GET("/books/:id", bookHandler.GetBookById)
	// post
	v1.POST("/create-book", bookHandler.CreateBooksHandler)
	// put
	v1.PUT("/update-book/:id", bookHandler.UpdateBooksHandler)
	// delete
	v1.DELETE("delete-book/:id", bookHandler.DeleteByID)

	// koneksi db

	router.Run()
}
