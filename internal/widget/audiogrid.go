package widget

import (
    "fmt"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

func CallFromInternalDir(str string) {
    fmt.Println(str)
}

func updateContent(clock *widget.Label) {
    formatted := time.Now().Format("Time: 03:04:05")
    clock.SetText(formatted)
}

func DrawThings(a fyne.App) {
    w := a.NewWindow("Koboard")
    w.SetMaster()

    clock := widget.NewLabel("")
    updateContent(clock)

    w.Resize(fyne.NewSize(400, 80))

    go func () {
        for range time.Tick(time.Second) {
            updateContent(clock)
        }
    } () 

    button := widget.NewButton("Click mew", func() {
        w2 := a.NewWindow("Hola chico")
        w2.SetContent(widget.NewLabel("Yhep"))
        w2.Show()
    })

    w.SetContent(container.New(layout.NewVBoxLayout(), clock, button))

    w.Show()
}
