// $go run . text

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func CreateText() {
	var questionedAlready = false
	for {
		if !questionedAlready {
			BotMessage("How can I help you?")
			questionedAlready = true
		} else {
			BotMessage("Anything else I can help?")
		}

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

		resp, err := OPENAI_BOT.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: yourMessage,
					},
				},
			},
		)

		if err != nil {
			BotMessage(fmt.Sprintf("ChatCompletion error: %v\n", err))
			return
		}

		endTime := time.Now()
		duration := endTime.Sub(startTime)

		BotMessage(fmt.Sprintf("It took %v seconds to create the response.", duration))

		var text = resp.Choices[0].Message.Content
		BotMessage(text)

		saveText := BotQuestion(fmt.Sprintf("\n%s Do you want to save it?\n", BOT_EMOJI))
		if strings.ToLower(saveText) == "y" || strings.HasPrefix(strings.ToLower(saveText), "y") {
			if _, err := os.Stat(TEXTS_FOLDER); os.IsNotExist(err) {
				err := os.Mkdir(TEXTS_FOLDER, os.ModePerm)
				if err != nil {
					BotMessage(fmt.Sprintf("Error creating folder: %v", err))
					return
				}
			}

			textFileName := BotQuestion(fmt.Sprintf("\n%s What is the name of the text file?\n", BOT_EMOJI))
			if textFileName == "" {
				currentTimestamp := time.Now().Unix()
				textFileName = fmt.Sprintf("%d", currentTimestamp)
			}

			textFileName = strings.TrimSpace(textFileName)
			// textFilePath := filepath.Join(TEXTS_FOLDER, textFileName+"."+TEXT_FILE_EXT)
			textFilePath := filepath.Join(TEXTS_FOLDER, fmt.Sprintf("%s.%s", textFileName, TEXT_FILE_EXT))

			err := os.WriteFile(textFilePath, []byte(text), 0644)
			if err != nil {
				BotMessage(fmt.Sprintf("Error writing to file: %v", err))
				return
			}

			fmt.Printf("\nThe response %s was saved to %s\n", textFileName, textFilePath)
		}
	}
}
