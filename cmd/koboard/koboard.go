package main

import (
	"fmt"

	"github.com/KosumovicDenis/Koboard/internal/soundboard"
	"github.com/KosumovicDenis/Koboard/internal/widget"
	"github.com/gordonklaus/portaudio"

	"fyne.io/fyne/v2/app"
)

func main() {
    portaudio.Initialize()
    defer portaudio.Terminate()
    soundboard.StartSoundboard()

    a := app.New()
    widget.DrawThings(a)
    a.Run()
    fmt.Println("Exited")
}

