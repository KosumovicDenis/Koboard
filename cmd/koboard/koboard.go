package main

import (
	"fmt"
    
    "github.com/KosumovicDenis/Koboard/internal/widget"
    "github.com/KosumovicDenis/Koboard/pkg/audio"

    "fyne.io/fyne/v2/app"
)

func main() {
    fmt.Print("Calling internal function -> ")
    widget.CallFromInternalDir("Funziona!!!")
    
    fmt.Println("---------------------------------")
    fmt.Println(audio.GetVersion())
    audio.PlayAudio()

    fmt.Println("Koboard")
    a := app.New()
    widget.DrawThings(a)
    a.Run()
    fmt.Println("Exited")
}
