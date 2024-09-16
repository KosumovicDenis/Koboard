package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func selectFile(w fyne.Window) {
    file_dialog := dialog.NewFileOpen(func (r fyne.URIReadCloser, err error) {
        chk(err)
        if r != nil && r.URI() != nil {
            addAudio(r.URI().Path())
        } 
    }, w)
    file_dialog.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
    w.Resize(fyne.NewSize(700, 600))
    w.SetFixedSize(true)
    file_dialog.Resize(fyne.NewSize(700, 600))
    file_dialog.SetOnClosed(func() {
        w.Hide()
    })
    w.SetCloseIntercept(func() {/* Do nothig. Must select a file or cancel */})
    w.Show()
    w.RequestFocus()
    file_dialog.Show()
}

