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
