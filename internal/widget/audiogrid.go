package widget

import (
	"fmt"
	"strings"

	"github.com/KosumovicDenis/Koboard/internal/model"
	"github.com/KosumovicDenis/Koboard/internal/soundboard"
	"github.com/KosumovicDenis/Koboard/pkg/audio"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var audios = 0

func updateContent(c *fyne.Container, audios *[]*model.Audio) {
    for _, a := range *audios {        
        audio_button := widget.NewButton(a.Name , func() {
            go audio.PlayAudio(a.Path)
        })
        audio_button.Resize(fyne.NewSize(100, 100))
        c.Add(audio_button) 
    }
}

func addAudio(path string, c * fyne.Container) {
    profiles := []*model.Profile{}
    soundboard.GetProfiles(&profiles)
    fmt.Println(path, c)
    for _, p := range profiles {
        if p.Active {
            array_path := strings.Split(path, "/")
            p.Audios = []*model.Audio{
                {
                    Id: 1,
                    Name: array_path[len(array_path)-1],
                    Path: path,
                    Format: model.FileFormat_MP3,
                },
            }
            updateContent(c, &p.Audios)
        } 
    }
}

func selectFile(w fyne.Window, c *fyne.Container) {
    file_dialog := dialog.NewFileOpen(func (r fyne.URIReadCloser, err error) {
        chk(err)
        if r != nil && r.URI() != nil {
            addAudio(r.URI().Path(), c)
        } 
    }, w)
    file_dialog.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
    w.Resize(fyne.NewSize(700, 600))
    w.SetFixedSize(true)
    w.Show()
    file_dialog.Resize(fyne.NewSize(700, 600))
    file_dialog.Show()
    file_dialog.SetOnClosed(func() {
        w.Hide()
    })
}

func DrawThings(a fyne.App) {
    fmt.Println("Drawing app")
    w := a.NewWindow("Koboard")
    w.SetMaster()

    w.Resize(fyne.NewSize(800, 800))

    c := container.New(layout.NewGridWrapLayout(fyne.NewSize(300,100)))

    openFileButton := widget.NewButton("Open audio file to use", func () {
        w2 := a.NewWindow("Select mp3 file")
        selectFile(w2, c)
    })

    displayProfiles(c)

    c.Add(openFileButton)

    w.SetContent(c)

    w.Show()
}

func displayProfiles(c *fyne.Container) {
    profiles := []*model.Profile{}
    soundboard.GetProfiles(&profiles)
    
    list := widget.NewList(
            func() int {
                return len(profiles)
            },
            func() fyne.CanvasObject {
                return widget.NewLabel("test")
            },
            func(i widget.ListItemID, o fyne.CanvasObject) {
                o.(*widget.Label).SetText(profiles[i].GetName())
            },
        )
    c.Add(list)
}

func chk(err error) {
    if err != nil { panic(err) }
}
