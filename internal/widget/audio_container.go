package widget

import (
	"strings"

	"github.com/KosumovicDenis/Koboard/internal/model"
	"github.com/KosumovicDenis/Koboard/internal/soundboard"
	"github.com/KosumovicDenis/Koboard/pkg/audio"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var audios = 0
var audios_container *fyne.Container

func updateAudiosConatiner(p *model.Profile) {
    audios_container.RemoveAll()
    for _, a := range p.GetAudios() {        
        audio_button := widget.NewButton(a.Name , func() {
            go audio.PlayAudio(a.Path)
        })
        audio_button.Resize(fyne.NewSize(100, 100))
        audios_container.Add(audio_button) 
    }
}

func displayAudios() {
    soundboard.GetProfiles(&profiles)
    for _, p := range profiles {
        if p.Active {
            updateAudiosConatiner(p)
        } 
    }
}

func addAudio(path string) {
    soundboard.GetProfiles(&profiles)
    for _, p := range profiles {
        if p.Active {
            array_path := strings.Split(path, "/")
            audio := &model.Audio{
                    Id: 1,
                    Name: array_path[len(array_path)-1],
                    Path: path,
                    Format: model.FileFormat_MP3,
            }
            p.Audios = append(p.Audios, audio)
            updateAudiosConatiner(p)
            soundboard.SaveSoundboard()
        } 
    }
}

func chk(err error) {
    if err != nil { panic(err) }
}
