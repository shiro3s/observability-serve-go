package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"template_app/factory"
	"template_app/models"
	"template_app/utils"

	"github.com/labstack/echo/v4"
)

func FindTodoAll(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	cond := models.NewTodoSearchParameter()
	err := c.Bind(&cond)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	todos, count := service.FindAll(cond)
	c.Response().Header().Set("x-count", fmt.Sprint(count))
	return c.JSON(http.StatusOK, todos)
}

func FindTodoById(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	cond := models.TodoParam{
		IsDeleted: false,
	}
	err := c.Bind(&cond)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Path Parameter")
	}
	todo, err := service.FindById(id, cond)
	if err != nil {
		e := utils.ErrorWrap(err)
		return c.JSON(e.Code, e.Msg.Error())
	}

	return c.JSON(http.StatusOK, &todo)
}

func CreateTodo(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	postBody := models.TodoBody{}
	err := c.Bind(&postBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = service.Create(postBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}

func UpdateTodo(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	putBody := models.TodoBody{}
	err := c.Bind(&putBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Path Parameter")
	}

	err = service.Update(id, putBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}

func DeleteTodo(c echo.Context) error {
	service := factory.NewTodoFactory(c).TodoService()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Path Parameter")
	}

	err = service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "OK")
}
