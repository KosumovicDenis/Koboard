package widget

import (
    "fmt"
    "strconv"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

var audios = 0

func CallFromInternalDir(str string) {
    fmt.Println(str)
}

func updateContent(c *fyne.Container, audioNum int) {
   c.Add(widget.NewLabel(strconv.Itoa(audioNum))) 
}

func DrawThings(a fyne.App) {
    w := a.NewWindow("Koboard")
    w.SetMaster()

    w.Resize(fyne.NewSize(400, 80))

    c := container.New(layout.NewGridWrapLayout(fyne.NewSize(100,100)))

    button := widget.NewButton("Click mew", func() {
        updateContent(c, audios)
        audios++;
    })

    c.Add(button)

    w.SetContent(c)

    w.Show()
}
