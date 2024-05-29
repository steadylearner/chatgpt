package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

var (
	OPENAI_API_KEY string
	OPENAI_BOT     *openai.Client

	// images or Use yours instead /Users/<YOURS>/Desktop/images in production
	IMAGES_FOLDER = "/Users/steadylearner/Desktop/images"
	IMAGE_SIZE    = "1792x1024" // 1024x1024
	QUALITY       = "hd"        // standard, hd

	// Use yours instead
	DESKTOP_WIDTH  = 1440
	DESKTOP_HEIGHT = 900
)

func init() {
	// DEV
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set environment variables
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

	if OPENAI_API_KEY == "" {
		log.Fatal("OPENAI_API_KEY environment variable not set")
	}

	// PROD
	// $go build -o chatgpt
	// $./chatgpt
	OPENAI_API_KEY = "YOURS"

	OPENAI_BOT = openai.NewClient(OPENAI_API_KEY)
}
