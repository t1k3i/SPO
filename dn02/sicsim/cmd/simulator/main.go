package main

import (
	"sicsim/cmd/simulator/gui"
	"sicsim/pkg/machine"
)

func main() {
	machine := machine.NewMachine()

	gui.StartGUI(machine)
}