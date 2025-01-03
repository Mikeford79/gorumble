package main

import (
    "github.com/andlabs/ui"
)

func main() {
    err := ui.Main(func() {
        window := ui.NewWindow("Test GUI", 300, 100, true)
        label := ui.NewLabel("Hello, world!")
        window.SetChild(label)
        window.OnClosing(func(*ui.Window) bool {
            ui.Quit()
            return true
        })
        window.Show()
    })
    if err != nil {
        panic(err)
    }
}
