package main

import (
	"fmt"
    
    "github.com/KosumovicDenis/Koboard/internal/widget"
    "github.com/KosumovicDenis/Koboard/pkg/audio"

    "fyne.io/fyne/v2/app"
)

func main() {    
    fmt.Println(audio.GetVersion())
    audio.PlayAudio()

    a := app.New()
    widget.DrawThings(a)
    a.Run()
    fmt.Println("Exited")
}
