package gui

import (
	"fmt"
	"image/color"
	"sicsim/pkg/loader"
	"sicsim/pkg/machine"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func StartGUI(m *machine.Machine) {
	myApp := app.NewWithID("sicxe.simulator")
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

	// Registers
	registerRow := container.NewGridWithColumns(3,
		widget.NewLabel("A:        000000"),
		widget.NewLabel("X:     000000"),
		widget.NewLabel("L:     000000"),
		widget.NewLabel("S:        000000"),
		widget.NewLabel("T:     000000"),
		widget.NewLabel("B:     000000"),
		widget.NewLabel("SW:    000000"),
		widget.NewLabel("F:     000000000000"),
	)
	registerPC := widget.NewLabel("PC:     000000")

	// Memory
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
				address, m.GetByte(int32(address)), m.GetByte(int32(address) + 1), m.GetByte(int32(address) + 2), 
				m.GetByte(int32(address) + 3), m.GetByte(int32(address) + 4), m.GetByte(int32(address) + 5), 
				m.GetByte(int32(address) + 6), m.GetByte(int32(address) + 7), m.GetByte(int32(address) + 8), 
				m.GetByte(int32(address) + 9), m.GetByte(int32(address) + 10), m.GetByte(int32(address) + 11), 
				m.GetByte(int32(address) + 12), m.GetByte(int32(address) + 13), m.GetByte(int32(address) + 14), 
				m.GetByte(int32(address) + 15)))
		},
	)

	// Action buttons
	start := widget.NewButton("START", nil)
	stop := widget.NewButton("STOP", nil)
	step := widget.NewButton("STEP", nil)
	undo := widget.NewButton("UNDO", nil)
	resetButton := widget.NewButton("RESET", nil)
	start.Disable()
	stop.Disable()
	step.Disable()
	undo.Disable()
	resetButton.Disable()

	// Frequency entry
	frequencyEntry := widget.NewEntry()
	frequencyLabel := widget.NewLabel("Speed (ms):")
	frequencyEntry.SetPlaceHolder("Speed in ms")
	frequencyEntry.SetText(fmt.Sprintf("%d", m.GetSpeed().Milliseconds()))

	// Load file
	objectFile := "None"
	objectFileLabel := widget.NewLabel(fmt.Sprintf("Object file: %s", objectFile))
	objectFileLabel.Wrapping = fyne.TextWrapWord

	// Action buttons actions
	step.OnTapped = func() {
		m.Step()
		updateRegisterValues(m, registerRow, registerPC)
		updateAllMemory(m, memoryList)
	}
	undo.OnTapped = func() {
		m.Undo()
		updateRegisterValues(m, registerRow, registerPC)
		updateAllMemory(m, memoryList)
	}
	stop.OnTapped = func() {
		m.Stop()

		frequencyEntry.Enable()
    	resetButton.Enable()
    	step.Enable()
    	undo.Enable()
	}
	start.OnTapped = func() {
		frequencyEntry.Disable()
    	resetButton.Disable()
    	step.Disable()
    	undo.Disable()

		go startInGui(m, registerRow, registerPC, memoryList, step, undo, resetButton, frequencyEntry)
	}
	resetButton.OnTapped = func() {
		if objectFile != "None" {
			m.Reset()
			err := loader.Load(objectFile, m)
			if err != nil {
				objectFile = "None"
				objectFileLabel.SetText(fmt.Sprintf("Object file: %s", objectFile))
				resetButton.Disable()
				step.Disable()
				undo.Disable()
				start.Disable()
				stop.Disable()
				return
			}
			updateAllMemory(m, memoryList)
			updateRegisterValues(m, registerRow, registerPC)
			step.Enable()
			undo.Enable()
			start.Enable()
			stop.Enable()
		}
	}

	actionButtonsGrid := container.NewGridWithColumns(2,
		start, stop, step, undo,
	)

	actionResetGrid := container.NewGridWithColumns(3,
		widget.NewLabel(" "),
		resetButton,
		widget.NewLabel(" "),
	)

	// Load
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
				step.Disable()
				undo.Disable()
				start.Disable()
				stop.Disable()
				return
			}
			updateAllMemory(m, memoryList)
			updateRegisterValues(m, registerRow, registerPC)
			step.Enable()
			undo.Enable()
			start.Enable()
			stop.Enable()
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

	// Frequency action
	frequencyEntry.OnSubmitted = func(value string) {
		newFreq, err := strconv.Atoi(value)
		if err != nil || newFreq <= 0 {
			frequencyEntry.SetText(fmt.Sprintf("%d", m.GetSpeed().Milliseconds()))
			return
		}
	
		m.SetSpeed(time.Duration(newFreq) * time.Millisecond)
	}

	frequencyRow := container.NewGridWithColumns(4, frequencyLabel, frequencyEntry, 
		widget.NewLabel(" "), widget.NewLabel(" "))

	leftContainer := container.NewVBox(topLeft, registerRow, registerPC,
		emptyLine, line, middleLeft, actionButtonsGrid, actionResetGrid,
		emptyLine, line2, bottomLeft, objectFileGrid, emptyLine, frequencyRow)

	scrollableMemory := container.NewVScroll(memoryList)
	scrollableMemory.SetMinSize(fyne.NewSize(415, 600))

	rightContainer := container.NewVBox(
		rightSection,
		scrollableMemory,
	)

	gridLayout := container.NewHBox(leftContainer, separator, rightContainer)

	myWindow.SetContent(container.NewCenter(gridLayout))
	myWindow.Resize(fyne.NewSize(850, 655))
	myWindow.ShowAndRun()
}

func updateRegisterValues(m *machine.Machine, registerRow *fyne.Container, registerPC *widget.Label) {
    registerRow.Objects[0].(*widget.Label).SetText(fmt.Sprintf("A:        %06X", m.GetA()))
    registerRow.Objects[1].(*widget.Label).SetText(fmt.Sprintf("X:     %06X", m.GetX()))
    registerRow.Objects[2].(*widget.Label).SetText(fmt.Sprintf("L:     %06X", m.GetL()))
    registerRow.Objects[3].(*widget.Label).SetText(fmt.Sprintf("S:        %06X", m.GetB()))
    registerRow.Objects[4].(*widget.Label).SetText(fmt.Sprintf("T:     %06X", m.GetT()))
    registerRow.Objects[5].(*widget.Label).SetText(fmt.Sprintf("B:     %06X", m.GetB()))
    registerRow.Objects[6].(*widget.Label).SetText(fmt.Sprintf("SW:    %06X", m.GetSW()))
    registerRow.Objects[7].(*widget.Label).SetText(fmt.Sprintf("F:     %012X", m.GetF()))
	registerPC.SetText(fmt.Sprintf("PC:     %06X", m.GetPC()))
}

func updateAllMemory(m *machine.Machine, memoryList *widget.List) {
	memoryItems := make([]string, machine.NUM_OF_ADDRESES/16)

	for i := 0; i < machine.NUM_OF_ADDRESES; i += 16 {
		memoryItems[i/16] = fmt.Sprintf("0x%06X: %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X %02X", 
			i, m.GetByte(int32(i)), m.GetByte(int32(i+1)), m.GetByte(int32(i+2)), m.GetByte(int32(i+3)), m.GetByte(int32(i+4)), 
			m.GetByte(int32(i+5)), m.GetByte(int32(i+6)), m.GetByte(int32(i+7)), m.GetByte(int32(i+8)), m.GetByte(int32(i+9)),
			m.GetByte(int32(i+10)), m.GetByte(int32(i+11)), m.GetByte(int32(i+12)), m.GetByte(int32(i+13)), 
			m.GetByte(int32(i+14)), m.GetByte(int32(i+15)),
		)
	}

	memoryList.Length = func() int {
		return len(memoryItems)
	}
	memoryList.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
		label := o.(*widget.Label)
		label.SetText(memoryItems[i])
	}
	
	memoryList.Refresh()
}

func startInGui(m *machine.Machine, registerRow *fyne.Container, registerPC *widget.Label, memoryList *widget.List,
				step *widget.Button, undo *widget.Button, reset *widget.Button, freq *widget.Entry) {
	m.SetPaused(false)
	
	ticker := time.NewTicker(m.GetSpeed())
	defer ticker.Stop()
	for m.IsRunning() {
		<- ticker.C
		if m.IsRunning() {
			m.Step()
			updateRegisterValues(m, registerRow, registerPC)
			updateAllMemory(m, memoryList)
		}
	}

	step.Enable()
	undo.Enable()
	freq.Enable()
	reset.Enable()
}