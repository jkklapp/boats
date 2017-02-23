package app

import (
    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
)

func main() {
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()
    router.LoadHTMLGlob("templates/**/*")

    // Group of API calls
    api := router.Group("/api")
    {
        api.GET("/healthcheck", healthcheck)
    }

    web := router.Group("/")
    {
        web.GET("/", func(c *gin.Context) {
            // TODO actual posts
            var posts [0]int
            c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts})
        })
    }

    // By default it serves on :8080 unless a
    // PORT environment variable was defined.
    router.Run()
    // router.Run(":3000") for a hard coded port
}

func healthcheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status":  "ok"})
}