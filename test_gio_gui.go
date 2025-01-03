package main

import (
    "gioui.org/app"
    "gioui.org/io/system"
    "gioui.org/layout"
    "gioui.org/widget/material"
    "gioui.org/unit"
    "log"
)

func main() {
    go func() {
        // Create a new window.
        w := app.NewWindow()
        th := material.NewTheme()
        for e := range w.Events() {
            if e, ok := e.(system.FrameEvent); ok {
                gtx := layout.NewContext(&e.Queue, e)
                layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
                    return material.H1(th, "Hello, world!").Layout(gtx)
                })
                e.Frame(gtx.Ops)
            }
        }
    }()
    // Run the Gio main event loop.
    app.Main()
}
