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
    defer portaudio.Terminate() // by using defer portaudio will be terminated only at the end of the function

    PrintDevices()
    
    PrintDefaultHostApi()

    PrintAllHostApi()
}

func PrintDevices() error {
    devicesInfo, err := portaudio.Devices()
    if (err != nil) {
        fmt.Print("Errore nella lettura dei dispositivi.\n")
        return err
    }
    fmt.Println("----- Printing Devices -----\n", "Numero dispositivi:", len(devicesInfo))
    for i := range devicesInfo {
        fmt.Print(i + 1, ". ", devicesInfo[i], "\n")
    }
    return nil
}

func PrintDefaultHostApi() error {
    hostAPi, err := portaudio.DefaultHostApi()
    if (err != nil) {
        fmt.Println("Errore nella lettura del default Host Api")
        return err
    }
    fmt.Println("Default host api: ", hostAPi)
    return nil
}

func PrintAllHostApi() error {
    allHostApi, err := portaudio.HostApis()
    if (err != nil) {
        fmt.Println("Errore nella lettura dei HostApi")
        return err
    }
    fmt.Println("----- Printing HostApis -----\n", "Numero HostAPi:", len(allHostApi))
    for i := range allHostApi {
        fmt.Print(i + 1, ". ", allHostApi[i], "\n")
    }
    return nil
}
