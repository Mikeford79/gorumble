package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Test GUI")

    label := widget.NewLabel("Hello, world!")
    myWindow.SetContent(container.NewVBox(
        label,
    ))

    myWindow.ShowAndRun()
}
