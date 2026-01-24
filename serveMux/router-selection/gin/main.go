package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- Handlers ---

// UserHandler は、ルーターに依存しない形式で実装されたハンドラの例です
func UserHandler(id string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Detail for ID: %s (Gin Bridge)\n", id)
	})
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		// 標準ハンドラにブリッジする例
		UserHandler(id).ServeHTTP(c.Writer, c.Request)
	})

	fmt.Println("Server (gin) starting on :8080...")
	r.Run(":8080")
}
