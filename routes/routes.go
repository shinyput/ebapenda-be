package routes

import (
	"Go-Starter-Project/controllers"

	"github.com/labstack/echo"
)

// InitGetRoutes - Dichiara tutte le route GET
func InitGetRoutes(e *echo.Echo) {
	e.GET("user/all", controllers.GetAllUser)
}

// InitPostRoutes - Dichiara tutte le route POST
func InitPostRoutes(e *echo.Echo) {}