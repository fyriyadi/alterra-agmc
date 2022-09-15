package controller

import (
	"day2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooks(c echo.Context) error {
	books := []model.Book{
		{
			ID:        1,
			Title:     "How to Read",
			Author:    "William N.",
			Publisher: "Anthropy",
		}, {
			ID:        2,
			Title:     "How to Unread",
			Author:    "Naritha W.",
			Publisher: "Ganesh",
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "successfuly retrive book data",
		"books":   books,
	})
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book := model.Book{
		ID:        uint(id),
		Title:     "How to Read",
		Author:    "William N.",
		Publisher: "Anthropy",
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfuly retrive book data",
		"books":   book,
	})
}

func CreateBook(c echo.Context) error {
	book := model.Book{
		ID:        1,
		Title:     "How to Read",
		Author:    "William N.",
		Publisher: "Anthropy",
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "successfuly create book data",
		"books":   book,
	})
}

func UpdateBookById(c echo.Context) error {
	book := model.Book{
		ID:        1,
		Title:     "How to Read",
		Author:    "William N.",
		Publisher: "Anthropy",
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "successfuly update book data",
		"books":   book,
	})
}

func DeleteBook(c echo.Context) error {
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "successfuly delete book data",
	})
}
