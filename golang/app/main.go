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

// go:embed assets/icon.png
// var iconData []byte

func main() {
	a := app.New()
	w := a.NewWindow("ChatGPT Image Generator")

	w.Resize(fyne.NewSize(float32(DESKTOP_WIDTH), float32(DESKTOP_HEIGHT))) // Set initial window size

	// Set the icon for the application
	// icon := fyne.NewStaticResource("icon.png", iconData)
	// w.SetIcon(icon)

	label := widget.NewLabel("Describe the image with details")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	description := widget.NewLabel("")

	image := canvas.NewImageFromFile("")
	image.FillMode = canvas.ImageFillContain // Use ImageFillContain to maintain aspect ratio and fit the window

	// resetButton := widget.NewButton("Reset", func() {
	// 	input.SetText("")
	// 	description.SetText("")
	// 	image.File = ""
	// 	image.Refresh()
	// })

	// var topNav = container.NewHBox(
	// 	label,
	// 	// resetButton,
	// )

	loading := widget.NewProgressBarInfinite()
	loading.Hide()

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
