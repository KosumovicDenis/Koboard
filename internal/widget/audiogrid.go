package widget

import (
	"fmt"
	"strings"

	"github.com/KosumovicDenis/Koboard/pkg/audio"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var audios = 0

func updateContent(c *fyne.Container, file_path string) {
    path := strings.Split(file_path, "/")
    audio_button := widget.NewButton(path[len(path) - 1] , func() {
        go audio.PlayAudio(file_path)
    })
    audio_button.Resize(fyne.NewSize(100, 100))
    c.Add(audio_button) 
}

func selectFile(w fyne.Window, c *fyne.Container) {
    file_dialog := dialog.NewFileOpen(func (r fyne.URIReadCloser, err error) {
        chk(err)
        if r.URI() != nil {
            updateContent(c, r.URI().Path())
        } 
    }, w)
    file_dialog.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
    file_dialog.Show()
    file_dialog.SetOnClosed(func() {
        w.Hide()
    })
    w.Show()
    file_dialog.Resize(fyne.NewSize(700, 600))
}

func DrawThings(a fyne.App) {
    fmt.Println("Drawing app")
    w := a.NewWindow("Koboard")
    w.SetMaster()

    w.Resize(fyne.NewSize(800, 800))

    c := container.New(layout.NewGridWrapLayout(fyne.NewSize(300,100)))

    openFileButton := widget.NewButton("Open audio file to use", func () {
        w2 := a.NewWindow("Select mp3 file")
        w2.Resize(fyne.NewSize(700, 600))
        w2.SetFixedSize(true)
        selectFile(w2, c)
    })

    c.Add(openFileButton)

    w.SetContent(c)

    w.Show()
}

func chk(err error) {
    if err != nil { panic(err) }
}
