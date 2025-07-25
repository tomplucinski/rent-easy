package clients

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var client openai.Client

func InitOpenAIClient(apiKey string) {
	client = openai.NewClient(
		option.WithAPIKey(apiKey),
	)
}

func GetOpenAIClient() openai.Client {
	return client
}
