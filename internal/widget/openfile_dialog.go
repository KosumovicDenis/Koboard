package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func selectFile(w fyne.Window) {
    fileDialog := dialog.NewFileOpen(func (r fyne.URIReadCloser, err error) {
        chk(err)
        if r != nil && r.URI() != nil {
            newAudioPath = r.URI().Path()
        } 
    }, w)
    fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
    w.Resize(fyne.NewSize(700, 600))
    w.SetFixedSize(true)
    fileDialog.Resize(fyne.NewSize(700, 600))
    fileDialog.SetOnClosed(func() {
        w.Hide()
    })
    fileDialog.Show()
    w.Show()
    w.SetCloseIntercept(func() {/* Do nothig. Must select a file or cancel */})
}
