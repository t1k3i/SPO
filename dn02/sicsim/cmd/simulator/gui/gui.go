package gui

import (
	"fmt"
	"image/color"
	"sicsim/pkg/loader"
	"sicsim/pkg/machine"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func StartGUI(m *machine.Machine) {
	myApp := app.New()
	myWindow := myApp.NewWindow("SIC/XE Simulator")

	topLeft := widget.NewLabel("CPE:")
	middleLeft := widget.NewLabel("ACTIONS:")
	rightSection := widget.NewLabel("MEM:")
	bottomLeft := widget.NewLabel("SETTINGS:")

	// Separator
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1

	// Separator2
	line2 := canvas.NewLine(color.White)
	line2.StrokeWidth = 1

	// Vertical separator
	separator := canvas.NewLine(color.White)
	separator.StrokeWidth = 1

	// Empty line
	emptyLine := widget.NewLabel(" ")

	registerRow := container.NewGridWithColumns(3,
		container.NewHBox(widget.NewLabel("A:     "), widget.NewLabel("000000")),
		container.NewHBox(widget.NewLabel("X:   "), widget.NewLabel("000000")),
		container.NewHBox(widget.NewLabel("L:   "), widget.NewLabel("000000")),
		container.NewHBox(widget.NewLabel("S:     "), widget.NewLabel("000000")),
		container.NewHBox(widget.NewLabel("T:   "), widget.NewLabel("000000")),
		container.NewHBox(widget.NewLabel("B:   "), widget.NewLabel("000000")),
		container.NewHBox(widget.NewLabel("SW:  "), widget.NewLabel("000000")),
		container.NewHBox(widget.NewLabel("F:   "), widget.NewLabel("000000000000")),
	)
	registerPC := container.NewHBox(widget.NewLabel("PC:   "), widget.NewLabel("000000"))

	actionButtonsGrid := container.NewGridWithColumns(2,
		widget.NewButton("START", func() {}),
		widget.NewButton("STOP", func() {}),
		widget.NewButton("STEP", func() {}),
		widget.NewButton("UNDO", func() {}),
	)

	objectFile := "None"
	objectFileLabel := widget.NewLabel(fmt.Sprintf("Object file: %s", objectFile))
	objectFileLabel.Wrapping = fyne.TextWrapWord

	resetButton := widget.NewButton("RESET", nil)
	resetButton.OnTapped = func() {
		if objectFile != "None" {
			m.Reset()
			err := loader.Load(objectFile, m)
			if err != nil {
				objectFile = "None"
				objectFileLabel.SetText(fmt.Sprintf("Object file: %s", objectFile))
				resetButton.Disable()
				return
			}
			fmt.Println("Machine reset and file reloaded:", objectFile)
		}
	}
	resetButton.Disable()

	actionResetGrid := container.NewGridWithColumns(3,
		widget.NewLabel(" "),
		resetButton,
		widget.NewLabel(" "),
	)

	loadButton := widget.NewButton("LOAD", func() {
		dialog := dialog.NewFileOpen(func(fileURI fyne.URIReadCloser, err error) {
			if err != nil || fileURI == nil {
				return
			}
			filePath := fileURI.URI().Path()
			loadErr := loader.Load(filePath, m)
			if loadErr != nil {
				objectFile = "None"
				objectFileLabel.SetText(fmt.Sprintf("Object file: %s", objectFile))
				resetButton.Disable()
				return
			}
			objectFile = filePath
			objectFileLabel.SetText(fmt.Sprintf("Object file: %s", objectFile))
			resetButton.Enable()
		}, myWindow)
		dialog.Show()
	})

	objectFileGrid := container.NewGridWithColumns(2,
		objectFileLabel, 
		loadButton,
	)

	frequencyLabel := widget.NewLabel("Frequency:")
	frequencyEntry := widget.NewEntry()
	frequencyEntry.SetPlaceHolder("Set Frequency")
	frequencyEntry.SetText("100")

	frequencyRow := container.NewGridWithColumns(4, frequencyLabel, frequencyEntry, 
		widget.NewLabel(" "), widget.NewLabel(" "))

	leftContainer := container.NewVBox(topLeft, registerRow, registerPC,
		emptyLine, line, middleLeft, actionButtonsGrid, actionResetGrid,
		emptyLine, line2, bottomLeft, objectFileGrid, emptyLine, frequencyRow)

	totalAddresses := 1 << 24
	memoryList := widget.NewList(
		func() int {
			return (totalAddresses + 16) / 16 // +16 only for better visualization
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("000000: 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			label := object.(*widget.Label)
			address := id * 16
			label.SetText(fmt.Sprintf("0x%06X: %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X",
				address, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0))
		},
	)

	scrollableMemory := container.NewVScroll(memoryList)
	scrollableMemory.SetMinSize(fyne.NewSize(415, 600))

	rightContainer := container.NewVBox(
		rightSection,
		scrollableMemory,
	)

	gridLayout := container.NewHBox(leftContainer, separator, rightContainer)

	myWindow.SetContent(gridLayout)
	myWindow.Resize(fyne.NewSize(850, 655))
	myWindow.ShowAndRun()
}