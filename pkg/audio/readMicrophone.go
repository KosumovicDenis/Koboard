package audio

import (
    "fmt"

    "github.com/gordonklaus/portaudio"
)

func GetVersion() int {
    return portaudio.Version()
}

func PlayAudio() {
    fmt.Println("- Playing audio -")
    err := portaudio.Initialize()
    if (err != nil) {
        fmt.Println("Errore nell'inizializzazione di PortAudio:", err)
        return
    }
    fmt.Println("PortAudio Initialized")
    deviceInfo, err := portaudio.Devices()
    if (err != nil) {
        fmt.Println("Errore nella lettura dei dispositivi")
        return
    }
    fmt.Println(len(deviceInfo))
    for i := range deviceInfo {
        fmt.Println(deviceInfo[i])
    }
    defer portaudio.Terminate()
}
