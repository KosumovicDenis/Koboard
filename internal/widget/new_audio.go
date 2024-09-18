package widget

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var newAudioPath string

func newAudioWindow(a fyne.App, c *fyne.Container) {
    w := a.NewWindow("Add new sound")
    w.Resize(fyne.NewSize(300,250))
    w.SetFixedSize(true)
 
    w2 := a.NewWindow("Select mp3 file")
    selectFileButton := widget.NewButton("Select file", func() { 
        selectFile(w2)
    })

    formItem := widget.NewFormItem("Audio path", selectFileButton)

    items := make([]*widget.FormItem, 1)
    items[0] = formItem

    form := dialog.NewForm("", "Enter", "Cancel", items, func(b bool) {
        if b {
            if strings.Compare(newAudioPath, "") != 0 {
                addAudio(&newAudioPath, c)
            }
        }
    }, w)
    
    w.Show()
    form.Resize(fyne.NewSize(300, 250))
    form.Show()
    form.SetOnClosed(func() {
        w.Hide()
    })
}
