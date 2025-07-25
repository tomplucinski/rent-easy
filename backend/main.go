package main

import (
	"chatgpt-api-server/clients"
	"chatgpt-api-server/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY not set")
	}

	clients.InitOpenAIClient(apiKey)

	router := gin.Default()

	router.POST("/chat", controllers.ChatHandler)

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
