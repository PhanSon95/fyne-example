package desktop

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type HexClock struct {
	widget.BaseWidget
	label *widget.Label
}

func NewHexClock() *HexClock {
	hexClock := &HexClock{}
	hexClock.label = widget.NewLabel("")
	hexClock.ExtendBaseWidget(hexClock)

	return hexClock
}

func (h *HexClock) CreateRenderer() fyne.WidgetRenderer {
	return &hexClockRenderer{hexClock: h}
}

func (h *HexClock) Refresh() {
	h.label.SetText(time.Now().Format("15:04:05"))
	h.Refresh()
}

type hexClockRenderer struct {
	hexClock *HexClock
}

func (r *hexClockRenderer) MinSize() fyne.Size {
	return fyne.Size{Width: 100, Height: 40}
}

func (r *hexClockRenderer) Layout(size fyne.Size) {
	r.hexClock.label.Resize(size)
}

func (r *hexClockRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *hexClockRenderer) Destroy() {
	// Cleanup resources if needed
}

func (r *hexClockRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.hexClock.label}
}

func (r *hexClockRenderer) Refresh() {
	// No special handling for refresh
}
