package api

import (
	"github.com/labstack/echo/v4"
	"github.com/danielmiessler/fabric/internal/core"
	"github.com/danielmiessler/fabric/pkg/common"
	"net/http"
	"fmt"
	"log"
	"io"
	"bytes"
	"encoding/json"
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
	patterns, err := h.Fabric.Patterns.GetNames()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch patterns"})
	}
	return c.JSON(http.StatusOK, patterns)
}

func (h *Handlers) handleGetPatternContent(c echo.Context) error {
	patternName := c.Param("name")
	pattern, err := h.Fabric.Patterns.GetPattern(patternName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch pattern content"})
	}
	return c.String(http.StatusOK, pattern.Pattern)
}

type QueryRequest struct {
	Model   string `json:"model" form:"model"`
	Pattern string `json:"pattern" form:"pattern"`
	Query   string `json:"query" form:"query"`
	Options struct {
		Temperature      float64 `json:"temperature" form:"temperature"`
		TopP             float64 `json:"top_p" form:"top_p"`
		PresencePenalty  float64 `json:"presence_penalty" form:"presence_penalty"`
		FrequencyPenalty float64 `json:"frequency_penalty" form:"frequency_penalty"`
		Stream           bool    `json:"stream" form:"stream"`
	} `json:"options" form:"options"`
}

// handleChat handles the /chat endpoint, processing the queries and returning the response.
//
// Parameters:
// - c: The echo.Context object representing the HTTP request and response.
//
// Returns:
// - error: An error if the request is invalid or the processing fails, otherwise nil.
func (h *Handlers) handleChat(c echo.Context) error {
	// Read and log the incoming request payload
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Failed to read request body: %v", err),
		})
	}
	log.Println("Incoming request payload:", string(bodyBytes))

	// Rebind the request body
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var queries []QueryRequest

	contentType := c.Request().Header.Get("Content-Type")
	if contentType == "application/json" {
		if err := json.Unmarshal(bodyBytes, &queries); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("Invalid JSON format: %v", err),
			})
		}
	} else {
		// Handle form data
		var singleQuery QueryRequest
		if err := c.Bind(&singleQuery); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("Invalid form data: %v", err),
			})
		}
		queries = append(queries, singleQuery)
	}

	if len(queries) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No queries provided",
		})
	}

	var responses []string
	for i, query := range queries {
		if query.Model == "" || query.Pattern == "" || query.Query == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("Query %d is missing required fields (model, pattern, or query)", i+1),
			})
		}

		pattern, err := h.Fabric.Patterns.GetPattern(query.Pattern)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("Failed to fetch pattern '%s': %v", query.Pattern, err),
			})
		}

		fmt.Println(pattern)
		
		chatter, err := h.Fabric.GetChatter(query.Model, query.Options.Stream)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("Failed to initialize chatter for model '%s': %v", query.Model, err),
			})
		}

		message := common.Message{
			Role:    "user",
			Content: query.Query,
		}

		request := &common.ChatRequest{
			Message:     message.Content,
			PatternName: query.Pattern,
			ContextName: "default",
			SessionName: "default",
		}

		options := &common.ChatOptions{
			Model:            query.Model,
			Temperature:      query.Options.Temperature,
			TopP:             query.Options.TopP,
			PresencePenalty:  query.Options.PresencePenalty,
			FrequencyPenalty: query.Options.FrequencyPenalty,
			// Stream:           query.Options.Stream,
		}

		response, err := chatter.Send(request, options)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": fmt.Sprintf("Failed to process chat request: %v", err),
			})
		}

		responses = append(responses, response)
	}

	// Return the response
	return c.JSON(http.StatusOK, responses)
}