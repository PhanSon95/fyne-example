package desktop

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewMainWindow(app fyne.App) fyne.Window {
	mainWindow := app.NewWindow("Main Window")
	mainWindow.Resize(fyne.NewSize(900, 500))

	return mainWindow
}

func Run() {
	fyneApp := app.New()
	mainWindow := NewMainWindow(fyneApp)
	mainWindow.SetTitle("My app")

	if desk, ok := fyneApp.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				log.Println("Tapped show")
				mainWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	mainWindow.SetCloseIntercept(func() {
		mainWindow.Hide()
	})

	hexClock := NewHexClock()
	mainWindow.SetTitle("my app")
	mainWindow.SetContent(container.NewVBox(
		layout.NewSpacer(),
		hexClock,
		layout.NewSpacer(),
	))

	go func() {
		for {
			time.Sleep(1 * time.Second)
			hexClock.Refresh()
		}
	}()

	mainWindow.ShowAndRun()
}
