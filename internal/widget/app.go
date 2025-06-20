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

	w.SetFixedSize(true)

	w.SetMaster()
	setSystemTrayMenu(a, w)

	vc := container.New(layout.NewVBoxLayout())
	audioContainer := container.New(layout.NewAdaptiveGridLayout(3))

	newAudioButton := widget.NewButton("New sound", func() {
		newAudioWindow(a, audioContainer)
	})

	settingsButton := widget.NewButton("Settings", func() {
		OpenSettings(a)
	})

	displayAudios(audioContainer)

	vc.Add(newAudioButton)
	vc.Add(settingsButton)
	vc.Add(audioContainer)

	w.SetContent(vc)

	w.SetCloseIntercept(func() {
		w.Hide()
	})
	w.Show()
}
