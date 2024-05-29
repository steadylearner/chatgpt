// $go run . image

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func CreateImage() {
	for {
		BotMessage("Can you describe the image you want to create with details?")

		fmt.Printf("\n%s You\n", USER_EMOJI)
		var yourMessage = strings.TrimSpace(BotQuestion("")) // What does steadylearner mean?

		if len(yourMessage) == 0 {
			BotMessage("Please, type something.")
			break
		}

		if yourMessage == "!quit" {
			break
		}

		startTime := time.Now()

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

		if err != nil {
			BotMessage(fmt.Sprintf("ChatCompletion error: %v\n", err))
			return
		}

		endTime := time.Now()
		duration := endTime.Sub(startTime)

		BotMessage(fmt.Sprintf("It took %v to create the image.", duration))

		imageUrl := chatGptResponse.Data[0].URL
		BotMessage(fmt.Sprintf("Here is the link to the image.\n\n%s", imageUrl))

		saveImage := BotQuestion(fmt.Sprintf("\n%s Do you want to save it?\n", BOT_EMOJI))
		if strings.ToLower(saveImage) == "y" || strings.HasPrefix(strings.ToLower(saveImage), "y") {
			if _, err := os.Stat(IMAGES_FOLDER); os.IsNotExist(err) {
				err := os.Mkdir(IMAGES_FOLDER, os.ModePerm)
				if err != nil {
					BotMessage(fmt.Sprintf("Error creating folder: %v", err))
					return
				}
			}

			imageFileName := strings.TrimSpace(BotQuestion(fmt.Sprintf("\n%s What is the name of the image?\n", BOT_EMOJI)))
			if imageFileName == "" {
				currentTimestamp := time.Now().Unix()
				imageFileName = strconv.FormatInt(currentTimestamp, 10)
			}

			imageFileName = strings.TrimSpace(imageFileName)
			// textFilePath := filepath.Join(TEXTS_FOLDER, textFileName+"."+TEXT_FILE_EXT)
			imageFilePath := filepath.Join(IMAGES_FOLDER, fmt.Sprintf("%s.%s", imageFileName, "png"))

			response, err := http.Get(imageUrl)
			if err != nil {
				BotMessage(fmt.Sprintf("Unable to download the image with error below. \n\n %v", err))
				return
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				data, err := io.ReadAll(response.Body)
				if err != nil {
					BotMessage(fmt.Sprintf("Unable to read the image data with error below. \n\n %v", err))
					return
				}

				err = os.WriteFile(imageFilePath, data, 0644)
				if err != nil {
					BotMessage(fmt.Sprintf("Unable to save the image with error below. \n\n %v", err))
					return
				}

				fmt.Printf("\nThe image %s was saved to %s\n", imageFileName, imageFilePath)
			} else {
				BotMessage("Unable to save the image")
			}
		}
	}
}
