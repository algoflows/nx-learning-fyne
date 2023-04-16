package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type App struct {
    output *widget.Label
    dark   bool
}

var myApp App

func main() {
    a := app.New()
    w := a.NewWindow("Hello World")

    output, entry, buttons := myApp.makeUI()

    w.SetContent(container.NewVBox(output, entry, buttons))
    w.Resize(fyne.NewSize(400, 200))
    w.ShowAndRun()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *fyne.Container) {
    output := widget.NewLabel("Hello World!")
    entry := widget.NewEntry()

    enterBtn := widget.NewButton("Enter", func() {
        app.output.SetText(entry.Text)
    })

    enterBtn.Importance = widget.HighImportance

    selectOptions := []string{"Light Theme", "Dark Theme"}
    selectWidget := widget.NewSelect(selectOptions, func(selected string) {
        if selected == "Dark Theme" {
            app.dark = true
            fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
        } else {
            app.dark = false
            fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
        }
    })
	selectWidget.PlaceHolder = "Select Theme"

    buttons := container.NewHBox(enterBtn, selectWidget)

    app.output = output

    return output, entry, buttons
}