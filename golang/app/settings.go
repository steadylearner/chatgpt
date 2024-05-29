package main

import (
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
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("No .env file found")
	// }

	// // Set environment variables
	// OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

	// if OPENAI_API_KEY == "" {
	// 	log.Fatal("OPENAI_API_KEY environment variable not set")
	// }

	// PROD
	// $go build -o chatgpt
	// $./chatgpt

	// Or this
	// https://docs.fyne.io/started/packaging.html
	// $go install fyne.io/fyne/v2/cmd/fyne@latest
	// $go get fyne.io/fyne/v2/cmd/fyne

	// $ls $HOME/go/bin
	// $vim ~/.zshrc
	// export PATH=$PATH:$HOME/go/bin
	// $source ~/.zshrc
	// $fyne package -os darwin -icon icon.png --name SteadylearnerChatGPT
	// $fyne package -os darwin -icon icon.png --name ChatGPT
	// $cd SteadylearnerChatGPT.app/Contents/MacOS && ./app
	// $cd ChatGPT.app/Contents/MacOS && ./app

	OPENAI_API_KEY = "YOURS"

	OPENAI_BOT = openai.NewClient(OPENAI_API_KEY)
}
