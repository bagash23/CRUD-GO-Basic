package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) SemuaBookHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var booksResponse []book.BookResoponse
	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	buku, err := h.bookService.FindByID(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	booksResponse := convertToBookResponse(buku)

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) CreateBooksHandler(c *gin.Context) {
	var bookReq book.BookRequest
	err := c.ShouldBindJSON(&bookReq)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := []string{}
			for _, e := range errors {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	book, err := h.bookService.Create(bookReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    convertToBookResponse(book),
		"message": "book created",
	})
}

func (h *bookHandler) UpdateBooksHandler(c *gin.Context) {

	var bookReq book.BookRequest
	err := c.ShouldBindJSON(&bookReq)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			errorMessages := []string{}
			for _, e := range errors {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    convertToBookResponse(book),
		"message": "book update",
	})
}

func (h *bookHandler) DeleteByID(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	buku, err := h.bookService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	booksResponse := convertToBookResponse(buku)

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func convertToBookResponse(books book.Book) book.BookResoponse {
	return book.BookResoponse{
		ID:          books.ID,
		Title:       books.Title,
		Price:       books.Price,
		Description: books.Description,
		Rating:      books.Rating,
		Discount:    books.Discount,
	}
}
