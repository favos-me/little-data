package routes

import (
	"github.com/favos-me/little-data/web/controllers"
	"github.com/labstack/echo/v4"
)

func NewRoutes() *echo.Echo {
	e := echo.New()
	e.GET("", controllers.Index)
	e.POST("/task/create", controllers.CreateTask)
	e.POST("/task/start", controllers.StartTask)
	e.GET("/log/:task_id", controllers.OperateLogIndex)
	return e
}
