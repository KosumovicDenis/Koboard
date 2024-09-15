package main

import (
	"fmt"

	"github.com/KosumovicDenis/Koboard/internal/soundboard"
	"github.com/KosumovicDenis/Koboard/internal/widget"

	"fyne.io/fyne/v2/app"
)

func main() {
    soundboard.StartSoundboard()

    a := app.New()
    widget.DrawThings(a)
    a.Run()
    fmt.Println("Exited")
}

