package route

import (
	"day2/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	booksV1 := v1.Group("/books")
	booksV1.GET("", controller.GetBooks)
	booksV1.GET("/:id", controller.GetBookById)
	booksV1.POST("", controller.CreateBook)
	booksV1.PUT("/:id", controller.UpdateBookById)
	booksV1.DELETE("/:id", controller.DeleteBook)

	usersV1 := v1.Group("/users")
	usersV1.GET("", controller.GetUser)
	usersV1.GET("/:id", controller.GetUserById)
	usersV1.POST("", controller.CreateUser)
	usersV1.PUT("/:id", controller.UpdateUserById)
	usersV1.DELETE("/:id", controller.DeleteUserById)

	return e
}
