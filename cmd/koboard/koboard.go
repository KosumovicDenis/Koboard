package main

import (
	"github.com/KosumovicDenis/Koboard/internal/soundboard"
	"github.com/KosumovicDenis/Koboard/internal/widget"
	"github.com/gordonklaus/portaudio"
)

func main() {
    portaudio.Initialize()
    defer portaudio.Terminate()

    soundboard.StartSoundboard()    

    widget.StartApp()
}

