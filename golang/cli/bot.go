package main

import (
	"bufio"
	"fmt"
	"os"
)

func BotQuestion(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	userResponse, _ := reader.ReadString('\n')
	return userResponse
}

func UserMessage(text string) {
	fmt.Printf("\n%s User\n%s\n", USER_EMOJI, text)
}

func BotMessage(text string) {
	fmt.Printf("\n%s Bot\n%s\n", BOT_EMOJI, text)
}
