package client

import (
	"context"
	"heya/config"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	openaiClient *openai.Client
}

type Role string

type Msgs = []openai.ChatCompletionMessage

type Convo struct {
	Model       string
	Msgs        Msgs
	Temperature float32
	MaxTokens   int
	Title       string
	Slug        string
}

//ChatCompletionRequest

func NewClient() *Client {
	return &Client{
		openaiClient: openai.NewClient(config.OpenAIApiKey),
	}
}
func (co *Convo) Prompt(ctx context.Context) string {
	chat := openai.ChatCompletionRequest{
		Model:       co.Model,
		Temperature: co.Temperature,
		Messages:    co.Msgs,
	}
	// x := co.openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
	// 	Model:       co.Client.model,
	// 	Prompt:      co.Title,
	// 	Temperature: co.Temperature,
	// 	// MaxTokens: ,

	// })
	// println(x)
	// return "hi"

}

// func
