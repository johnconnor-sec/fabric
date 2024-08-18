package main

import (
    "io"
    "github.com/danielmiessler/fabric/internal/api"
    "github.com/danielmiessler/fabric/internal/core"
    "html/template"
    "github.com/danielmiessler/fabric/internal/db"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "os/user"
    "path/filepath"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
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

    fabric, err := core.NewFabric(fabricDB)
    if err != nil {
        panic(err)
    }

    // Initialize Echo
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Serve static files
    e.Static("/static", "web/static")

    // Register the renderer
    t := &Template{
        templates: template.Must(template.ParseGlob("web/templates/*.html")),
    }
    e.Renderer = t

    // Setup routes
    api.SetupRoutes(e, fabric)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}