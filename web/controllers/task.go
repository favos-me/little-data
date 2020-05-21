package controllers

import (
	"github.com/favos-me/little-data/dbactions"
	"github.com/favos-me/little-data/models"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var params models.Params
	tasks, err := dbactions.QueryTasks(&params)
	if err != nil {
		return err
	}
	_ = tasks
	return nil
}

func CreateTask(c echo.Context) error {
	var task models.Task
	err := c.Bind(&task)
	if err != nil {
		return err
	}
	return nil
}

func StartTask(c echo.Context) error {
	return nil
}
