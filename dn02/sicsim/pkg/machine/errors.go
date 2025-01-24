package machine

import "fmt"

/*
 *	ERRORS
 */
func checkValue(v int32) {
	if v < MinInt24 || v > MaxInt24 {
		panic("Values in SIC can be up to 24 bits!")
	}
}

func checkDeviceNumber(num int) {
	if num < 0 || num >= NUM_OF_DEVICES {
		panic("Invalid device number!")
	}
}

func notValidRegisterIndex() {
	panic("Not valid register index!")
}

func notImplemented() {
	panic("Not implemented!")
}

func notImplementedFloat() {
	panic("Does not support floats yet!")
}

func opcodeNotValid(opcode byte) {
	str := fmt.Sprintf("Operation code %d is not valid!", opcode)
	panic(str)
}

func invalidAddressing() {
	panic("Invalid addressing!")
}

func checkAddress(addr int32) {
	if addr < MIN_ADDRESS || addr > MAX_ADDRESS {
		panic("Addres not valid!")
	}
}

func fileDoesNotExist() {
	panic("File does not exist!")
}