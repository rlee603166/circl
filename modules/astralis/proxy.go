package astralis

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var astralisURL = os.Getenv("ASTRALIS_API_URL")

func ProxyAstralisResponse(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	fmt.Printf("Request body: %s\n", string(body))
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
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "LLM backend unreachable"})
		return
	}
	defer res.Body.Close()

	c.Writer.Header().Set("Content-Type", res.Header.Get("Content-Type"))
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
		n, err := res.Body.Read(buf)
		if n > 0 {
			writer.Write(buf[:n])
			flusher.Flush()
		}
		if err != nil {
			break
		}
	}

}

func ProxyAstralisSummary(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	url := astralisURL + "/search/summarize"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upstream request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "LLM backend unreachable"})
		return
	}
	defer res.Body.Close()

	for key, values := range res.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	c.Writer.WriteHeader(res.StatusCode)

	_, err = io.Copy(c.Writer, res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to copy response body"})
		return
	}
}

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/astralis", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Astralis health",
		})
	})
	rg.POST("/astralis/query", ProxyAstralisResponse)
	rg.POST("/astralis/summarize", ProxyAstralisSummary)
}
