<!-- # How to make Go Desktop App with Fyne and ChatGPT -->

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

[Fyne]: https://github.com/fyne-io/fyne

In this post, you will learn how to make a desktop app with Go [Fyne] and ChatGPT with the [openai-go] package. 

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
$mkdir app
$go mod init app
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

[You can include more files similar to what you can see here.](https://github.com/steadylearner/chatgpt/tree/main/golang)

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
https://github.com/fyne-io/fyne
```

Then, your go.mod file will be similar to this.

```console
// go.mod
module app

go 1.22.3

require fyne.io/fyne/v2 v2.4.5

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/akavel/rsrc v0.10.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/jackmordaunt/icns/v2 v2.2.6 // indirect
	github.com/josephspurrier/goversioninfo v1.4.0 // indirect
	github.com/lucor/goinfo v0.9.0 // indirect
	github.com/mcuadros/go-version v0.0.0-20190830083331-035f6764e8d2 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.4.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/tools v0.12.0 // indirect
	golang.org/x/tools/go/vcs v0.1.0-deprecated // indirect
)

require (
	fyne.io/fyne v1.4.3
	fyne.io/systray v1.10.1-0.20231115130155-104f5ef7839e // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fredbi/uri v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fyne-io/gl-js v0.0.0-20220119005834-d2da28d9ccfe // indirect
	github.com/fyne-io/glfw-js v0.0.0-20220120001248-ee7290d23504 // indirect
	github.com/fyne-io/image v0.0.0-20220602074514-4956b0afb3d2 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20240306074159-ea2d69986ecb // indirect
	github.com/go-text/render v0.1.0 // indirect
	github.com/go-text/typesetting v0.1.0 // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/joho/godotenv v1.5.1
	github.com/jsummers/gobmp v0.0.0-20151104160322-e2ba15ffa76e // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sashabaranov/go-openai v1.24.1
	github.com/srwiley/oksvg v0.0.0-20221011165216-be6e8873101c // indirect
	github.com/srwiley/rasterx v0.0.0-20220730225603-2ab79fcdd4ef // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tevino/abool v1.2.0 // indirect
	github.com/yuin/goldmark v1.5.5 // indirect
	golang.org/x/image v0.11.0 // indirect
	golang.org/x/mobile v0.0.0-20230531173138-3c911d8e3eda // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	honnef.co/go/js/dom v0.0.0-20210725211120-f030747120f2 // indirect
)
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
```

Use the folder path you want to save `images` and use your own **OPENAI_API_KEY** instead. Then, you can build your own desktop app with commands like this.

```console
$fyne package -os darwin -icon icon.png --name ChatGPT
```

[You can find more details from the official docs website.](https://docs.fyne.io/started/packaging.html)

Then, first we create bot.go file used to generate an image with [openai-go].

```go
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
```

We can find a link to an image in `openai.ImageResponse` part.

### A Desktop app

First, we will create a main.go that includes the most logics for a desktop app to work.

```go
package main

import (
	// _ "embed"
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("ChatGPT Image Generator")

	w.Resize(fyne.NewSize(float32(DESKTOP_WIDTH), float32(DESKTOP_HEIGHT))) // Set initial window size

	r, _ := fyne.LoadResourceFromPath("icon.png")
	w.SetIcon(r)

	// Include all parts you want to use for this desktop app
	label := widget.NewLabel("Describe the image with details")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	description := widget.NewLabel("")
	image := canvas.NewImageFromFile("")
	image.FillMode = canvas.ImageFillContain // Use ImageFillContain to maintain aspect ratio and fit the window
	loading := widget.NewProgressBarInfinite()
	loading.Hide()

	// Include a callback to handle the app logics
	input.OnSubmitted = func(text string) {

		var yourMessage = strings.TrimSpace(text)

		if yourMessage == "" {
			description.SetText("Please enter some text.")
			return
		}

		if yourMessage == "!quit" {
			w.Close()
			return
		}

		description.SetText("")
		loading.Show()

		startTime := time.Now()
		chatGptResponse, err := CreateImage(yourMessage)
		endTime := time.Now()
		duration := endTime.Sub(startTime)

		if err != nil {
			loading.Hide()
			description.SetText("Error fetching image: " + err.Error())

			return
		}

		loading.Hide()

		imageUrl := chatGptResponse.Data[0].URL

		parsedUrl, err := storage.ParseURI(imageUrl)
		if err != nil {
			description.SetText("Error parsing URL: " + err.Error())
			return
		}

		newImage := canvas.NewImageFromURI(parsedUrl)
		newImage.FillMode = canvas.ImageFillContain
		newImage.SetMinSize(fyne.NewSize(float32(DESKTOP_WIDTH), float32(DESKTOP_HEIGHT))) // Set a minimum size for the image
		newImage.Refresh()

		tappableImage := NewTappableImage(newImage, func() {
			saveImage(imageUrl, w)
		})

		content := container.NewVBox(
			// topNav,
			label,
			input,
			loading,
			description,
			tappableImage,
		)

		w.SetContent(content)
		description.SetText(fmt.Sprintf("It took %v to create the image. Click the image to save it.", duration))
	}

	w.SetContent(container.NewVBox(
		// topNav,
		label,
		input,
		loading,
		description,
	))

	w.ShowAndRun()
}
```

In the code snippet, we include the parts we want to use for the app first.

```go
label := widget.NewLabel("Describe the image with details")
input := widget.NewEntry()
input.SetPlaceHolder("Enter text...")
description := widget.NewLabel("")
image := canvas.NewImageFromFile("")
image.FillMode = canvas.ImageFillContain // Use ImageFillContain to maintain aspect ratio and fit the window
loading := widget.NewProgressBarInfinite()
loading.Hide()
```

Then, when a user type a prompt to create an image every logic is handled by a callback for `input`.

```go
input.OnSubmitted = func(text string) {
	var yourMessage = strings.TrimSpace(text)

	if yourMessage == "" {
		description.SetText("Please enter some text.")
		return
	}

	if yourMessage == "!quit" {
		w.Close()
		return
	}

	description.SetText("")
	loading.Show()

	startTime := time.Now()
	chatGptResponse, err := CreateImage(yourMessage)
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	if err != nil {
		loading.Hide()
		description.SetText("Error fetching image: " + err.Error())

		return
	}

	loading.Hide()

	imageUrl := chatGptResponse.Data[0].URL

	parsedUrl, err := storage.ParseURI(imageUrl)
	if err != nil {
		description.SetText("Error parsing URL: " + err.Error())
		return
	}

	newImage := canvas.NewImageFromURI(parsedUrl)
	newImage.FillMode = canvas.ImageFillContain
	newImage.SetMinSize(fyne.NewSize(float32(DESKTOP_WIDTH), float32(DESKTOP_HEIGHT))) // Set a minimum size for the image
	newImage.Refresh()

	tappableImage := NewTappableImage(newImage, func() {
		saveImage(imageUrl, w)
	})

	content := container.NewVBox(
		// topNav,
		label,
		input,
		loading,
		description,
		tappableImage,
	)

	w.SetContent(content)
	description.SetText(fmt.Sprintf("It took %v to create the image. Click the image to save it.", duration))
}
```

Then, we can have a image.go to show a dialog to help you save an image from openai api.

```go
package main

import (
	"io"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type TappableImage struct {
	widget.BaseWidget
	image    *canvas.Image
	onTapped func()
}

func NewTappableImage(image *canvas.Image, onTapped func()) *TappableImage {
	t := &TappableImage{
		image:    image,
		onTapped: onTapped,
	}
	t.ExtendBaseWidget(t)
	return t
}

func (t *TappableImage) Tapped(*fyne.PointEvent) {
	if t.onTapped != nil {
		t.onTapped()
	}
}

func (t *TappableImage) TappedSecondary(*fyne.PointEvent) {}

func (t *TappableImage) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.image)
}

func saveImage(url string, w fyne.Window) {
	imageFolder := "/Users/steadylearner/Desktop/images"

	// Check if the directory exists, if not, create it
	if _, err := os.Stat(imageFolder); os.IsNotExist(err) {
		err := os.MkdirAll(imageFolder, os.ModePerm)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
	}

	fileDialog := dialog.NewFileSave(
		func(writer fyne.URIWriteCloser, _ error) {
			if writer == nil {
				return
			}
			defer writer.Close()

			response, err := http.Get(url)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			defer response.Body.Close()

			_, err = io.Copy(writer, response.Body)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
		}, w)

	fileDialog.SetFileName("image.png")
	fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))

	// Set the initial directory to the specified directory
	listableURI, err := storage.ListerForURI(storage.NewFileURI(imageFolder))
	if err == nil {
		fileDialog.SetLocation(listableURI)
	}

	fileDialog.Show()
}
```

When you use `$go run .` with the correct set up you will see a desktop app similar to the cover of this post.

## Packaging for distribution

With Go [Fyne], you can easily package a graphical app for distribution with a command similar to this.

```console
$fyne package -os darwin -icon icon.png --name ChatGPT
```

Then, it will create a ChatGPT.app folder and you will be able to find your app at Contents/MacOS.

[You can find more details from the official docs website.](https://docs.fyne.io/started/packaging.html)

## Conclusion

You have learned how to create a simple desktop app with [fyne] and [openai-go]. 

If you thought this post helpful, please share it and like the [ChatGPT repository][Steadylearner ChatGPT repository] and blog post.

You can also join [Go community](https://www.onlycoiners.com/community/go) with images created from the app described in this post. 

[You can also hire me if you need a full stack developer for your projects.][Hire me]