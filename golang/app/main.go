package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func fetchRandomImageURL() (string, error) {
	const imageURL = "https://avatars.githubusercontent.com/u/32325099?v=4"
	resp, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch random image, status code: %d", resp.StatusCode)
	}

	return resp.Request.URL.String(), nil
}

func main() {
	a := app.New()
	w := a.NewWindow("ChatGPT Image Generator") // Title
	// Unable to quit
	// w.SetFullScreen(true)                       // Make the window full screen
	w.Resize(fyne.NewSize(1440, 900)) // Use your desktop size here.

	label := widget.NewLabel("Describe the image you want with details")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	description := widget.NewLabel("")

	image := canvas.NewImageFromFile("")
	image.FillMode = canvas.ImageFillOriginal

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

		startTime := time.Now()
		url, err := fetchRandomImageURL()
		endTime := time.Now()
		duration := endTime.Sub(startTime)

		if err != nil {
			description.SetText("Error fetching image: " + err.Error())
			return
		}

		parsedUrl, err := storage.ParseURI(url)
		if err != nil {
			description.SetText("Error parsing URL: " + err.Error())
			return
		}

		newImage := canvas.NewImageFromURI(parsedUrl)
		newImage.FillMode = canvas.ImageFillOriginal
		newImage.Refresh()

		content := container.NewVBox(
			label,
			input,
			description,
			newImage,
		)

		w.SetContent(content)
		description.SetText(fmt.Sprintf("It took %v to create the image.", duration))
	}

	// Set initial content with the label, input bar, and placeholder image
	w.SetContent(container.NewVBox(
		label,
		input,
		image,
	))

	w.ShowAndRun()
}
