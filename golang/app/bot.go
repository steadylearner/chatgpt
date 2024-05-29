package main

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

func CreateImage(yourMessage string) (openai.ImageResponse, error) {
	chatGptResponse, err := OPENAI_BOT.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:  yourMessage,
			Model:   openai.CreateImageModelDallE3,
			N:       1,
			Quality: QUALITY,
			Size:    IMAGE_SIZE,
		},
	)

	return chatGptResponse, err
}
