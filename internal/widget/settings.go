package widget

import (
	"fmt"

	"github.com/KosumovicDenis/Koboard/internal/model"
	"github.com/KosumovicDenis/Koboard/internal/soundboard"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var profiles []*model.Profile

func OpenSettings(a fyne.App) {
    w := a.NewWindow("Settings")

    w.Resize(fyne.NewSquareSize(250))

    c := container.NewVBox()
    label := widget.NewLabel("Profiles:")
    c.Add(label)
    displayProfiles(c)
    
    changeProfileButton := widget.NewButtonWithIcon("Change profile", theme.ConfirmIcon(), func () {
        fmt.Println("TODO: Change profile")
    })

    c.Add(changeProfileButton)

    w.SetContent(container.NewVScroll(c))

    w.Show()
}

func displayProfiles(c *fyne.Container) {
    profiles := []*model.Profile{}
    soundboard.GetProfiles(&profiles)
    
    list := widget.NewList(
            func() int {
                return len(profiles)
            },
            func() fyne.CanvasObject {
                return widget.NewLabel("test")
            },
            func(i widget.ListItemID, o fyne.CanvasObject) {
                o.(*widget.Label).SetText(profiles[i].GetName())
            },
        )
    c.Add(list)
}
