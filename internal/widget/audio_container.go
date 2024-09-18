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

func addAllAudiosToConatiner(p *model.Profile, c *fyne.Container) {
    c.RemoveAll()
    for _, a := range p.GetAudios() {        
        audioButton := widget.NewButton(a.Name , func() {
            go audio.PlayAudio(a.Path)
        })
        c.Add(audioButton) 
    }
}
func addAudioToContainer(a *model.Audio,  c *fyne.Container) {
    audioButton := widget.NewButton(a.Name , func() {
        go audio.PlayAudio(a.Path)
    })
    c.Add(audioButton) 
}

func displayAudios(c *fyne.Container) {
    soundboard.GetProfiles(&profiles)
    for _, p := range profiles {
        if p.Active {
            addAllAudiosToConatiner(p, c)
        } 
    }
}

func addAudio(path *string, c *fyne.Container) {
    for _, p := range profiles {
        if p.Active {
            arrayPath := strings.Split(*path, "/")
            audio := &model.Audio{
                    Id: 1,
                    Name: arrayPath[len(arrayPath)-1],
                    Path: *path,
                    Format: model.FileFormat_MP3,
            }
            p.Audios = append(p.Audios, audio) 
            soundboard.SaveSoundboard()
            addAudioToContainer(audio, c)
        } 
    }
    *path = ""
}

func chk(err error) {
    if err != nil { panic(err) }
}
