package handler

import (
    "net/http"
    "50HW/internal/qrcode"
    "github.com/gin-gonic/gin"
)

type QRCodeRequest struct {
    Size    int    `form:"size"`
    Content string `form:"content"`
}

func GenerateQRCode(c *gin.Context) {
    var req QRCodeRequest

    if err := c.ShouldBind(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if req.Size == 0 {
        req.Size = 256 
    }

    png, err := qrcode.Generate(req.Content, req.Size)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
        return
    }

    c.Header("Content-Type", "image/png")
    c.Header("Content-Disposition", "attachment; filename=qrcode.png")
    c.Writer.Write(png)
}
