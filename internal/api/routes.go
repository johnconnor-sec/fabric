package api

import (
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, fabric *core.Fabric) {
	handlers := NewHandlers(fabric)

	e.GET("/", handlers.handleHome)
	e.POST("/chat", handlers.handleChat)
	e.GET("/patterns", handlers.handleListPatterns)
	e.GET("/patterns/:name", handlers.handleGetPatternContent)
	e.GET("/models", handlers.handleListModels)
}