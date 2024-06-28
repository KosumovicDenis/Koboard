package main

import (
	"fmt"
    "time"

    "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateContent(clock *widget.Label) {
    formatted := time.Now().Format("Time: 03:04:05")
    clock.SetText(formatted)
}

func main() {
    fmt.Println("Koboard")
    a := app.New()
    w := a.NewWindow("Koboard")

    clock := widget.NewLabel("")
    updateContent(clock)

    w.SetContent(clock)
    w.Resize(fyne.NewSize(400, 80))

    go func () {
        for range time.Tick(time.Second) {
            updateContent(clock)
        }
    } () 

    w.Show()
    a.Run()
    fmt.Println("Exited")
}
