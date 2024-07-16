package widget

import (
    "fmt"
    "strconv"
    "time"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

var audios = 0
var cps = 0

func updateContent(c *fyne.Container, audioNum int) {
   c.Add(widget.NewLabel(strconv.Itoa(audioNum))) 
}

func DrawThings(a fyne.App) {
    fmt.Println("Drawing app")
    w := a.NewWindow("Koboard")
    w.SetMaster()

    w.Resize(fyne.NewSize(400, 80))

    c := container.New(layout.NewGridWrapLayout(fyne.NewSize(100,100)))

    button := widget.NewButton("Click mew", func() {
        updateContent(c, audios)
        audios++;
    })

    cpsLabel := widget.NewLabel("CPS -> 0")
    c.Add(cpsLabel)
    cpsButton := widget.NewButton("Cps counter", func() {
        cps++;
    })
    c.Add(cpsButton)
    
    start := time.Now()

    go func() {
        for range time.Tick(time.Second) {
            t := time.Now()
            elasped := t.Sub(start)
            cpsLabel.SetText(strconv.Itoa(cps / int(elasped.Seconds())));
        }
    } ()

    c.Add(button)

    w.SetContent(c)

    w.Show()
}
