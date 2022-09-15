package controller

import (
	"day2/config"
	"day2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	req := model.UserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	user := model.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
	err = config.DB.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "successfuly create user",
	})
}

func GetUser(c echo.Context) error {
	user := []model.User{}
	err := config.DB.Model(model.User{}).Scan(&user).Error
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "successfuly retrive user data",
		"data":    user,
	})
}

func GetUserById(c echo.Context) error {
	userId := c.Param("id")
	user := model.User{}
	err := config.DB.Model(model.User{}).Where("id = ?", userId).Scan(&user).Error
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "successfuly retrive user data",
		"data":    user,
	})
}

func UpdateUserById(c echo.Context) error {
	userId := c.Param("id")

	req := model.UserRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}
	intUserId, _ := strconv.Atoi(userId)
	user := model.User{
		ID:       uint(intUserId),
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
	err = config.DB.Model(model.User{}).Where("id = ?", userId).Updates(&user).Error
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "successfuly update user data",
		"data":    user,
	})
}

func DeleteUserById(c echo.Context) error {
	userId := c.Param("id")
	user := model.User{}
	err := config.DB.Model(model.User{}).Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "successfuly delete user data",
	})
}
