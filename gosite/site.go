package main

import (
  "log"
  "net/http"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

func main() {
  gin.SetMode(gin.ReleaseMode)

  r := gin.Default()

  // Define a health GET endpoint
  r.Any("/health", func(c *gin.Context) {
  	c.Status(http.StatusNoContent)
  })

  // Static file hosting for portfolio
  r.Use(static.Serve("/", static.LocalFile("./portfolio", false)))
	r.GET("/", func(c *gin.Context) {
		c.File("./portfolio/index.html")
	})

  // Start server on port 80 (may need elevated permissions)
  // Server will listen on localhost
	if err := r.Run(":80"); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}
