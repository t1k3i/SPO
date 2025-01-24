package main

import (
	"os"

	"sicsim/pkg/loader"
	"sicsim/pkg/machine"
)

func main() {
	file := os.Args[1]
	machine := machine.NewMachine()
	loader.Load(file, machine)
	machine.PrintMEM(100)
	/*machine := NewMachine()
	loadMEM("0100171B20033F2FFD000017", machine)
	go machine.Start()
	time.Sleep(time.Millisecond * 1000)
	machine.Stop()
	fmt.Println("machine stoped .......................................................");
	time.Sleep(time.Millisecond * 3000)
	go machine.Start()
	fmt.Println("Press Enter to stop the machine...")
	fmt.Scanln()*/
}

/*func loadMEM(bytes string, machine *Machine) {
	for i := 0; i < (len(bytes) / 2); i++ {
		toNumber, _ := strconv.ParseInt(bytes[i*2 : i*2+2], 16, 16)
		singleByte := byte(toNumber)
		machine.SetByte(int32(i), singleByte)
	}
}*/