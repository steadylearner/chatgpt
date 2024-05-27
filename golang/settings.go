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

	TEXTS_FOLDER  = "texts"
	TEXT_FILE_EXT = "md"
	IMAGES_FOLDER = "images"
	IMAGE_SIZE    = "1024x1024"
	QUALITY       = "hd" // standard, hd
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set environment variables
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

	if OPENAI_API_KEY == "" {
		log.Fatal("OPENAI_API_KEY environment variable not set")
	}

	OPENAI_BOT = openai.NewClient(OPENAI_API_KEY)
}
