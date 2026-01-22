package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- Handlers ---

// 標準の http.HandlerFunc を返す設計
func userDetailHandler(id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Detail for ID: %s (Gin Bridge)\n", id)
	}
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		// 標準ハンドラにブリッジする例
		userDetailHandler(id)(c.Writer, c.Request)
	})

	fmt.Println("Server (gin) starting on :8080...")
	r.Run(":8080")
}

