package main

import (
	"fmt"

	"github.com/KosumovicDenis/Koboard/internal/soundboard_data"
	"github.com/KosumovicDenis/Koboard/internal/widget"
	"github.com/KosumovicDenis/Koboard/pkg/audio"

	"fyne.io/fyne/v2/app"
)

func main() {    
    soundboard_data.InitSoundboard()
    soundboard_data.GetSoundBoard()
    fmt.Println(audio.GetVersion())

    a := app.New()
    widget.DrawThings(a)
    a.Run()
    fmt.Println("Exited")
}
