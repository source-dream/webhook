package main

import (
    "fmt"
	"os"
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
    // 加载 .env 文件
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
}

type WebhookPayload struct {
    Data struct {
        Message string `json:"message"`
    } `json:"data"`
}

func sendToMeow(payload WebhookPayload) {
	userId := os.Getenv("MEOW_USER_ID")
	title := os.Getenv("MEOW_TITLE")
	url := fmt.Sprintf("http://api.chuckfang.com/%s/%s/%s", userId, title, payload.Data.Message)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get err:", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("sendToMeow success")
}

func webhookHandler(c *gin.Context) {
    var payload WebhookPayload

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	fmt.Println("Received message:", payload.Data.Message)

	sendToMeow(payload)

    c.JSON(http.StatusOK, gin.H{"message": "WebHook received"})
}

func main() {
    r := gin.Default()
    r.POST("/", webhookHandler)
    r.Run(":6004")
}