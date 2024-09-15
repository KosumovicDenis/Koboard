package audio

import (
    "bytes"
    "os"
    "os/signal"
    "encoding/binary"
	"fmt"

	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/gordonklaus/portaudio"
)

func PlayAudio(audio_file string) { 
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	// create mpg123 decoder instance
	decoder, err := mpg123.NewDecoder("")
	chk(err)

	fileName := audio_file
	chk(decoder.Open(fileName))
	defer decoder.Close()

	// get audio format information
	rate, channels, _ := decoder.GetFormat()

	// make sure output format does not change
	decoder.FormatNone()
	decoder.Format(rate, channels, mpg123.ENC_SIGNED_16)

	portaudio.Initialize()
	defer portaudio.Terminate()
	out := make([]int16, 8192)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate), len(out), &out)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	defer stream.Stop()
	for {
		audio := make([]byte, 2*len(out))
		_, err = decoder.Read(audio)
		if err == mpg123.EOF {
			break
		}
		chk(err)

		chk(binary.Read(bytes.NewBuffer(audio), binary.LittleEndian, out))
		chk(stream.Write())
		select {
		case <-sig:
			return
		default:
		}
	}
}

func GetVersion() int {
    return portaudio.Version()
}

func PrintDefaultOutputDevice() error {
    defaultOutput , err := portaudio.DefaultOutputDevice()
    chk(err)
    fmt.Println("--- Printing Default Output ---\n", defaultOutput)
    return nil
}

func PrintDevices() error {
    devicesInfo, err := portaudio.Devices()
    chk(err)
    fmt.Println("----- Printing Devices -----\n", "Numero dispositivi:", len(devicesInfo))
    for i := range devicesInfo {
        fmt.Print(i + 1, ". ", devicesInfo[i], "\n")
    }
    return nil
}

func PrintDefaultHostApi() error {
    hostAPi, err := portaudio.DefaultHostApi()
    chk(err)
    fmt.Println("Default host api: ", hostAPi)
    return nil
}

func PrintAllHostApi() error {
    allHostApi, err := portaudio.HostApis()
    chk(err)
    fmt.Println("----- Printing HostApis -----\n", "Numero HostAPi:", len(allHostApi))
    for i := range allHostApi {
        fmt.Print(i + 1, ". ", allHostApi[i], "\n")
    }
    return nil
}

func chk (err error) {
    if err != nil { panic(err) } 
}
