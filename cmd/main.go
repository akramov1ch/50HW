package main

import (
    "log"
    "50HW/internal/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.POST("/generate", handler.GenerateQRCode)

    log.Println("Starting server on port 8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}
