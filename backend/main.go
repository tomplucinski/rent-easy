package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	router := gin.Default()

	router.POST("/chat", func(c *gin.Context) {
		var req struct {
			Prompt string `json:"prompt"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		chatCompletion, err := client.Chat.Completions.New(
			context.TODO(), openai.ChatCompletionNewParams{
				Messages: []openai.ChatCompletionMessageParamUnion{
					openai.ChatCompletionMessageParamUser{
						Content: "Say this is a test",
					},
				},
			},
		)

		if err != nil {
			log.Println("OpenAI error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI API error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"reply": chatCompletion.Choices[0].Message.Content,
		})
	})

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
