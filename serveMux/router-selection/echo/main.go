package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// --- Handlers ---

func userDetailHandler(id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Detail for ID: %s (Echo Bridge)\n", id)
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		// 標準ハンドラにブリッジする例
		userDetailHandler(id)(c.Response().Writer, c.Request())
		return nil
	})

	fmt.Println("Server (echo) starting on :8080...")
	e.Start(":8080")
}

