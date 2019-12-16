package ui

import (
	"testing"
	"github.com/andlabs/ui"
)

func TestUI(t *testing.T) {
	err := ui.Main(func() {
		keyInput := ui.NewEntry()
		expDate := ui.NewDatePicker()
		greeting := ui.NewLabel("key长度错误(16字节)")
		button := ui.NewButton("生成")
		button.Disable()

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("key: "), false)
		box.Append(keyInput, false)
		box.Append(ui.NewLabel("过期时间: "), false)
		box.Append(expDate, false)
		box.Append(button, false)
		box.Append(greeting, false)

		window := ui.NewWindow("License", 200, 100, false)
		window.SetMargined(true)
		window.SetChild(box)

		keyInput.OnChanged(func(entry *ui.Entry) {
			if len(entry.Text()) != 16 {
				button.Disable()
				greeting.SetText("key长度错误(16字节)")
			} else {
				button.Enable()
				greeting.SetText("")
			}
		})

		button.OnClicked(func(*ui.Button) {
			expDate.Handle()
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
