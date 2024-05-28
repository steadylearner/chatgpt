package main

import (
	"fmt"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// Replace this with CHATGPT things
// Function to fetch a random image URL
func fetchRandomImageURL() (string, error) {
	// Replace this with an actual API endpoint if needed
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
	w := a.NewWindow("Random Image Viewer")

	label := widget.NewLabel("Enter text and press Enter:")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")

	image := canvas.NewImageFromFile("")
	image.FillMode = canvas.ImageFillOriginal

	// Handle input change and fetch random image
	input.OnSubmitted = func(text string) {
		if strings.TrimSpace(text) == "" {
			label.SetText("Please enter some text.")
			return
		}
		url, err := fetchRandomImageURL()
		if err != nil {
			label.SetText("Error fetching image: " + err.Error())
			return
		}

		u, err := storage.ParseURI(url)
		if err != nil {
			label.SetText("Error parsing URL: " + err.Error())
			return
		}

		newImage := canvas.NewImageFromURI(u)
		newImage.FillMode = canvas.ImageFillOriginal
		newImage.Refresh()

		content := container.NewVBox(
			label,
			input,
			newImage,
		)
		w.SetContent(content)
		label.SetText("Showing random image for: " + text)
	}

	// Set initial content with the label, input bar, and placeholder image
	w.SetContent(container.NewVBox(
		label,
		input,
		image,
	))

	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}
