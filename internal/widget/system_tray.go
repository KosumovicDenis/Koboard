package widget

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func setSystemTrayMenu(a fyne.App, w fyne.Window) {
    if desk, ok := a.(desktop.App); ok {
        m := fyne.NewMenu("Koboard",
            fyne.NewMenuItem("Show", func() {
                w.Show()
            }),
            fyne.NewMenuItem("Quit", func() {
                a.Quit()
            }))
        data, err := os.ReadFile("logo.png")
        chk(err)
        desk.SetSystemTrayIcon(fyne.NewStaticResource("logo.png", data))
        desk.SetSystemTrayMenu(m)
    }
}
