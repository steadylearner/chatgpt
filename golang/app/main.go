package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

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

func main() {
	a := app.New()
	w := a.NewWindow("ChatGPT Image Generator")

	w.Resize(fyne.NewSize(float32(DESKTOP_WIDTH), float32(DESKTOP_HEIGHT))) // Set initial window size

	label := widget.NewLabel("Describe the image with details")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	description := widget.NewLabel("")

	image := canvas.NewImageFromFile("")
	image.FillMode = canvas.ImageFillContain // Use ImageFillContain to maintain aspect ratio and fit the window

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
		chatGptResponse, err := CreateImage(yourMessage)
		endTime := time.Now()
		duration := endTime.Sub(startTime)

		if err != nil {
			description.SetText("Error fetching image: " + err.Error())
			return
		}

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
			label,
			input,
			description,
			tappableImage,
		)

		w.SetContent(content)
		description.SetText(fmt.Sprintf("It took %v to create the image. Click the image to save it.", duration))
	}

	w.SetContent(container.NewVBox(
		label,
		input,
		image,
	))

	w.ShowAndRun()
}
