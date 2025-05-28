package astralis

import (
    "bytes"
    "io"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

var astralisURL = os.Getenv("ASTRALIS_API_URL")

func ProxyAstralis(c *gin.Context) {
    body, err := io.ReadAll(c.Request.Body)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }

    url := astralisURL + "/search/query"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upstream request"})
        return    
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.JSON(http.StatusBadGateway, gin.H{"error": "LLM backend unreachable"})
        return
    }
    defer resp.Body.Close()

    c.Writer.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
    c.Writer.Header().Set("Transfer-Encoding", "chunked")
    c.Writer.WriteHeader(http.StatusOK)

    writer := c.Writer
    flusher, ok := writer.(http.Flusher)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "streaming not supported"})
        return
    }

    buf := make([]byte, 1024)
    for {
        n, err := resp.Body.Read(buf)
        if n > 0 {
            writer.Write(buf[:n])
            flusher.Flush()
        }
        if err != nil {
            break
        }
    }

}

func RegisterRoutes(rg *gin.RouterGroup) {
    rg.POST("/astralis/query", ProxyAstralis)
}

