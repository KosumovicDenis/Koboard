package widget

import (
	"github.com/KosumovicDenis/Koboard/internal/model"
	"github.com/KosumovicDenis/Koboard/internal/soundboard"
	"github.com/KosumovicDenis/Koboard/pkg/audio"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func addAllAudiosToConatiner(p *model.Profile, c *fyne.Container) {
    c.RemoveAll()
    for _, a := range p.GetAudios() {        
		name := truncateString(a.Name, 25)
        audioButton := widget.NewButton(name, func() {
            go audio.PlayAudio(a.Path)
        })
        c.Add(audioButton) 
    }
}
func addAudioToContainer(a *model.Audio,  c *fyne.Container) {
	name := truncateString(a.Name, 25)
    audioButton := widget.NewButton(name, func() {
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

func addAudio(path *string, name string, c *fyne.Container) {
    for _, p := range profiles {
        if p.Active {
            audio := &model.Audio{
                    Id: 1,
                    Name: name,
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

func truncateString(s string, max int) string {
    if len(s) > max {
        if max > 3 {
            return s[:max-3] + "..."
        }
        return s[:max] // Se max è troppo piccolo, non aggiungiamo "..."
    }
    return s
}
