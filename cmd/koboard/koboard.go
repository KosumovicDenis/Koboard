package main

import (
	"fmt"

	"github.com/KosumovicDenis/Koboard/internal/soundboard_data"
	"github.com/KosumovicDenis/Koboard/internal/widget"
	"github.com/KosumovicDenis/Koboard/pkg/audio"

	"fyne.io/fyne/v2/app"
)

func main() {    
    soundboard_data.NewSoundBoard("Prima soundboard")
    soundboard_data.GetSoundBoard("Prima soundboard")
    fmt.Println(audio.GetVersion())

    a := app.New()
    widget.DrawThings(a)
    a.Run()
    fmt.Println("Exited")
}
