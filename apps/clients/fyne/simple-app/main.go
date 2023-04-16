package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(Hello("clients-fyne-simple-app"))
	a := app.New()

	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
