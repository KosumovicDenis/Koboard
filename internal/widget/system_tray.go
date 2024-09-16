package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func setSystemTrayMenu(a fyne.App, w fyne.Window) {
    if desk, ok := a.(desktop.App); ok {
        m := fyne.NewMenu("Koboard",
            fyne.NewMenuItem("Show", func() {
                w.Show()
            }))
        desk.SetSystemTrayMenu(m)
    }
}
