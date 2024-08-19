package main

import (
    "github.com/danielmiessler/fabric/internal/api"
    "github.com/danielmiessler/fabric/internal/core"
    "github.com/danielmiessler/fabric/internal/db"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "html/template"
    "path/filepath"
    "os/user"
    "os"
    "io"
)

// Template struct and Render method
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

type Fabric struct {
    *core.Fabric
    *db.Db
    *api.Handlers
}

func NewFabric(db *db.Db) (*Fabric, error) {
    fabric, err := core.NewFabric(db)
    if err != nil {
        return nil, err
    }
    handlers := api.NewHandlers(fabric)
    return &Fabric{Fabric: fabric, Db: db, Handlers: handlers}, nil
}

// Define the patternValue function
func patternValue() string {
    baseDir := "/patterns/*/system.md" // Update this to the actual base directory

    // List directories within the base directory
    files, err := filepath.Glob(baseDir)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        // Read the system.md file
        content, err := os.ReadFile(file)
        if err != nil {
            continue // Skip if there's an error reading the file
        }

        // Extract the pattern value from the file content
        patternValue := extractPatternValue(string(content))
        if patternValue != "" {
            return patternValue
        }
    }

    return "patternValue" // Return a default value if no pattern value is found
}

// Helper function to extract pattern value from file content
func extractPatternValue(content string) string {
    // Implement the logic to extract the pattern value
    // For example, if the pattern value is on the first line:
    return content // Assuming the entire content is the pattern value
}

func main() {
    // Initialize Fabric
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fabricDB := db.NewDb(filepath.Join(user.HomeDir, ".config/fabric"))
    if err := fabricDB.Configure(); err != nil {
        panic(err)
    }

    fabric, err := NewFabric(fabricDB)
    if err != nil {
        panic(err)
    }

    // Load patterns
    if err := fabric.PatternsLoader.(); err != nil {
        panic(err)
    }   

    // Initialize Echo
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Register the renderer with the function map
    funcMap := template.FuncMap{
        "patternValue": patternValue,
    }
    t := &Template{
        templates: template.Must(template.New("").Funcs(funcMap).ParseGlob("web/templates/*.html")),
    }
    e.Renderer = t

    // Setup routes
    api.SetupRoutes(e, fabric.Fabric)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}