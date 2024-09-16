package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartApp() {
    a := app.NewWithID("Koboard_1")
    mainWindow(a)
    a.Run()
}

func mainWindow(a fyne.App) { 
    w := a.NewWindow("Koboard")
    w.SetMaster()
    setSystemTrayMenu(a, w)

    vc := container.New(layout.NewVBoxLayout())

    w2 := a.NewWindow("Select mp3 file")
    openFileButton := widget.NewButton("Open mp3 file", func () {
        selectFile(w2)
    })

    vc.Add(openFileButton)

    audios_container = container.New(layout.NewGridLayoutWithColumns(4))
    displayAudios()

    vc.Add(audios_container)

    w.SetContent(vc)
    
    w.SetCloseIntercept(func() {
        w.Hide()
    })
    w.Show()
}
