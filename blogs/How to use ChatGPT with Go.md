<!-- # How to use ChatGPT with Go -->

[ChatGPT]: https://chatgpt.com
[openai-go]: https://github.com/sashabaranov/go-openai
[openai-api-keys]: https://platform.openai.com/account/api-keys
[openai-api-limits]: https://platform.openai.com/settings/organization/limits 
[openai-api-billing]: https://platform.openai.com/settings/organization/billing
[openai-api-usage]: https://platform.openai.com/usage 

[Steadylearner GitHub]: https://github.com/steadylearner
[Steadylearner ChatGPT repository]: https://github.com/steadylearner/chatgpt

[Hire me]: https://www.onlycoiners.com/user/steadylearner/product/i-will-help-you-to-build-a-full-stack-app-bot-or-smart-contr

[ChatGPT YouTube]: https://www.youtube.com/watch?v=JY7H286Jubg

In this post, you will learn how to use ChatGPT with Go using the official [openai-go] package. 

[You can find the code used here in this repository.][Steadylearner ChatGPT repository] It also has examples in other programming languages and blog posts for them.

[If you haven't used ChatGPT for images, you might find this useful.][ChatGPT YouTube]

I plan to include more examples with other programming languages, so feel free to like the repository if you find it useful and want to see more examples.

[You can also hire me if you need a full stack developer for your projects.][Hire me]

## Prerequisites

You can skip these and go to **Code** if you have done similar things already.

1. Install Go
2. Set up dev envrionment
3. Get OpenAI API key
4. Install required packages

### 1. Install Go

We will use Go for this post. [Visit this page to install it](https://go.dev/doc/install) and follow the guide to download it. To check if Go is installed, use your console and type this.

```console
$go version
```

### 2. Set up dev environment

Create a Go dev environment. [You can also refer to this.](https://go.dev/doc/tutorial/getting-started)

```console
$go mod init chatgpt/main
```

Then, paste this to your main.go file.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Then, use `$go run .` and you will see "Hello, World!" in your console.

[You can include more files similar to what you can see here.](https://github.com/steadylearner/chatgpt/tree/main/golang/chatgpt)

### 3. Get OpenAI API key

Sign up on the OpenAI platform and get your API key from the [API keys page][openai-api-keys]. Store it securely as you will need it to access the OpenAI API later.

You can find more information about OpenAI API usage, billing, and limits at the following links.

- [API Limits][openai-api-limits]
- [API Billing][openai-api-billing]
- [API Usage][openai-api-usage]

### 4. Install required packages

You can install pacakges required for this post with `$go get` command.

```console
https://github.com/sashabaranov/go-openai
https://github.com/urfave/cli
```

Then, your go.mod and go.sum file will be similar to this.

```console
// go.mod
module example/hello

go 1.22.3

require (
	github.com/joho/godotenv v1.5.1
	github.com/sashabaranov/go-openai v1.24.1
	github.com/urfave/cli/v2 v2.27.2
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
)

// go.sum
github.com/cpuguy83/go-md2man/v2 v2.0.4 h1:wfIWP927BUkWJb2NmU/kNDYIBTh/ziUX91+lVfRxZq4=
github.com/cpuguy83/go-md2man/v2 v2.0.4/go.mod h1:tgQtvFlXSQOSOSIRvRPT7W67SCa46tRHOmNcaadrF8o=
github.com/joho/godotenv v1.5.1 h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=
github.com/joho/godotenv v1.5.1/go.mod h1:f4LDr5Voq0i2e/R5DDNOoa2zzDfwtkZa6DnEwAbqwq4=
github.com/russross/blackfriday/v2 v2.1.0 h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=
github.com/russross/blackfriday/v2 v2.1.0/go.mod h1:+Rmxgy9KzJVeS9/2gXHxylqXiyQDYRxCVz55jmeOWTM=
github.com/sashabaranov/go-openai v1.24.1 h1:DWK95XViNb+agQtuzsn+FyHhn3HQJ7Va8z04DQDJ1MI=
github.com/sashabaranov/go-openai v1.24.1/go.mod h1:lj5b/K+zjTSFxVLijLSTDZuP7adOgerWeFyZLUhAKRg=
github.com/urfave/cli/v2 v2.27.2 h1:6e0H+AkS+zDckwPCUrZkKX38mRaau4nL2uipkJpbkcI=
github.com/urfave/cli/v2 v2.27.2/go.mod h1:g0+79LmHHATl7DAcHO99smiR/T7uGLw84w8Y42x+4eM=
github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 h1:+qGGcbkzsfDQNPPe9UDgpxAWQrhbbBXOYJFQDq/dtJw=
github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913/go.mod h1:4aEEwZQutDLsQv2Deui4iYQ6DWTxR14g6m8Wv88+Xqk=
```

It will be easier if you just clone the repository and start from there.

```console
$git clone https://github.com/steadylearner/chatgpt.git
```

## Code

### Configuration

First, create a .env file in your chatgpt folder with the following text.

```console
OPENAI_API_KEY=<YOURS>
```

Then, create a settings.go file to manage your settings.

```ts
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

	// texts or Use yours instead /Users/<YOURS>/Desktop/texts in production
	TEXTS_FOLDER  = "texts" // "/Users/steadylearner/Desktop/texts"
	TEXT_FILE_EXT = "md"

	// images or Use yours insted /Users/<YOURS>/Desktop/images in production
	IMAGES_FOLDER = "iamges" // "/Users/steadylearner/Desktop/images"
	IMAGE_SIZE    = "1024x1024"
	QUALITY       = "hd" // standard, hd
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
	// OPENAI_API_KEY = "YOURS"

	OPENAI_BOT = openai.NewClient(OPENAI_API_KEY)
}
```

If you want to make the CLI binary you can use in your console later, you can use other folder paths for `images` and `texts`` and use your **OPENAI_API_KEY = "YOURS"** similar to the example.

Finally, we will see two examples that you might find useful and make some variations later.

But before we code on, we will write **main.go** file similar to this to help you test the code snippets we will have later.

```go
package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "text",
				Aliases: []string{"t"},
				Usage:   "Create a text based on a prompt",
				Action: func(cCtx *cli.Context) error {
					CreateText()
					return nil
				},
			},
			{
				Name:    "image",
				Aliases: []string{"i"},
				Usage:   "Create an image based on a prompt",
				Action: func(cCtx *cli.Context) error {
					CreateImage()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
```

With it, you can test the CLI with `$go run . text` for text and `$go run . image` later.

### A CLI chatbot

First, we will see a simple code snippet for a CLI chatbot.

You can save the text below as **text.go** file.

```go
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
    // 1.
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

    // 2.
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

    // 3.
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
```

We already included this code snippet to index.ts file.

```go
{
  Name:    "text",
  Aliases: []string{"t"},
  Usage:   "Create a text based on a prompt",
  Action: func(cCtx *cli.Context) error {
    CreateText()
    return nil
  },
},
```

So we can test it with `$go run . text` and you will see some messages like this.

You can use your own question instead of "What does steadylearner mean?"

```console
🤖 Bot 
How can I help you?

🧑 You
What does steadylearner mean?

🤖 Bot
Steadylearner likely means someone who is constantly learning and seeking knowledge in a consistent and steady manner. This person is dedicated to their personal growth and development through continuous learning and improvement.

🤖 Bot 
Anything else I can help?
```

So how does it work? There are a few important parts but we will see only important ones.

1. We want to ask questions to the bot until we stop the bot with `Ctrl + C`, `!quit` or others that results in SIGINT signal.
2. The time for the ChatGPT API to respond varies depending on your question and it might take very long so we will have some feedbacks to decrease response time by measuring how much it takes.
3. If you liked the response and want to save it, you can optionally do that here.

So that was the simple code snippet needed to respond to your questions.

The **gpt-3.5-turbo** option was used here to save resources but you can also use **gpt-4o** option.

Not so interesting because you can use more features at [ChatGPT] website. Therefore, we will make something a little bit more useful variation.

### A CLI image creator

The code snippet used here is almost the same to the previous one but we will use **dall-e-3** model to make images. You can use it for profile image, blog cover, logo etc.

```go
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

		BotMessage(fmt.Sprintf("It took %v seconds to create the image.", duration))

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
```

You can test it with `$go run . image` command. Then, describe the image you want to create with details.

You will see similar messages you read in the previous process but you will see an image created with url and how much time it did take.

You can use another message instead of **Can you make a cover image for blog post with title "How to use chatgpt with Go"** used here.

```console
🤖 Bot 
Can you describe the image you want to create with details?

🧑 You
Can you make a cover image for blog post with title "How to use chatgpt with Go"

🤖 Bot
Here is the link to the image.

https://oaidalleapiprodscus.blob.core.windows.net/private

Do you want to save the image?
y
```

The profile and the cover image on this blog was created with it. 

You can also make your own images.

## Make a binary file for the CLI

With Go, you can easily make a CLI binary file you can easily use with your console.

Use these commands.

```console
$go build -o chatgpt
$sudo mv chatgpt /usr/local/bin/

or

$pwd
$sudo ln -s <YOURS> /usr/local/bin/chatgpt
``````

Then, you can use `$chatgpt text` or `$chatgpt image` to use the CLI you made here directly regardless of the current path in your console.

## Conclusion

You have learned how to create simple CLIs for text and image using TypeScript and [openai-go]. 

You can extend these examples to create more complex applications. 

If you thought this post helpful, please share it and like the [ChatGPT repository][Steadylearner ChatGPT repository] and blog post.

[You can also hire me if you need a full stack developer for your projects.][Hire me]