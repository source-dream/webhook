package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

type WebhookPayload struct {
    Data struct {
        Message string `json:"message"`
    } `json:"data"`
}

func sendToMeow(payload WebhookPayload) {
	userId := os.Getenv("MEOW_USER_ID")
	url := fmt.Sprintf("http://api.chuckfang.com/%s/站点监测/%s", userId, payload.Data.Message)
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