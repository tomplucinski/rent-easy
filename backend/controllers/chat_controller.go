package controllers

import (
	"chatgpt-api-server/clients"
	"chatgpt-api-server/models"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
)

func ChatHandler(c *gin.Context) {
	var req models.ChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client := clients.GetOpenAIClient()

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(req.Prompt),
		},
		Model: openai.ChatModelGPT4o,
	})

	if err != nil {
		log.Println("OpenAI error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI API error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reply": chatCompletion.Choices[0].Message.Content,
	})
}
