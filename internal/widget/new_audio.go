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

    formItem := widget.NewFormItem("Sound path", selectFileButton)
    soundName := widget.NewEntry()
    formItem2 := widget.NewFormItem("Sound Name", soundName)

    items := make([]*widget.FormItem, 2)
    items[0] = formItem
    items[1] = formItem2

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
