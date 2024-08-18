package api

import (
	"github.com/labstack/echo/v4"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/pkg/common"
	"net/http"
	// "encoding/json"
	"fmt"
)

type Handlers struct {
	Fabric *core.Fabric
}

func NewHandlers(fabric *core.Fabric) *Handlers {
	return &Handlers{Fabric: fabric}
}

func (h *Handlers) handleHome(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func (h *Handlers) handleListModels(c echo.Context) error {
	models := h.Fabric.GetModels()
	return c.JSON(http.StatusOK, models)
}

func (h *Handlers) handleListPatterns(c echo.Context) error {
	patterns, err := h.Fabric.Db.Patterns.GetNames()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch patterns"})
	}
	return c.JSON(http.StatusOK, patterns)
}

func (h *Handlers) handleGetPatternContent(c echo.Context) error {
	patternName := c.Param("name")
	pattern, err := h.Fabric.Db.Patterns.GetPattern(patternName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch pattern content"})
	}
	return c.String(http.StatusOK, pattern.Pattern)
}

func (h *Handlers) handleChat(c echo.Context) error {
	var chatRequest struct {
		Queries []struct {
			Model   string `json:"model"`
			Pattern string `json:"pattern"`
			Query   string `json:"query"`
		} `json:"queries"`
	}

	if err := c.Bind(&chatRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	if len(chatRequest.Queries) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No queries provided",
		})
	}

	var responses []string
	for i, query := range chatRequest.Queries {
		if query.Model == "" || query.Pattern == "" || query.Query == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("Query %d is missing required fields (model, pattern, or query)", i+1),
			})
		}

		pattern, err := h.Fabric.Db.Patterns.GetPattern(query.Pattern)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("Failed to fetch pattern '%s': %v", query.Pattern, err),
			})
		}

		fmt.Println(pattern)
		
		chatter, err := h.Fabric.GetChatter(query.Model, false)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("Failed to initialize chatter for model '%s': %v", query.Model, err),
			})
		}

		request := &common.ChatRequest{
			PatternName: query.Pattern,
			Message:     query.Query,
		}

		options := &common.ChatOptions{
			Model: query.Model,
			// Add other options as needed
		}

		response, err := chatter.Send(request, options)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("Failed to process chat request: %v", err),
			})
		}

		responses = append(responses, response)
	}

	return c.JSON(http.StatusOK, responses)
}