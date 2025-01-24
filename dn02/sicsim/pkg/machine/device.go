package machine

import (
	"fmt"
	"os"
)

/*
 *	DEVICE INTERFACE
 */
type Device interface {
	Test() bool
	Read() byte
	Write(value byte)
}

/*
 * COMMON DEVICE THAT HAS A GENERIC IMPLEMENTATION OF DEVICE FUNCTIONS
 */
type CommonDevice struct{}

func (d *CommonDevice) Test() bool {
	return true
}

func (d *CommonDevice) Read() byte {
	return 0
}

func (d *CommonDevice) Write(value byte) {}

/*
 * INPUT DEVICE
 */
type InputDevice struct {
	*CommonDevice
}

func (d *InputDevice) Read() byte {
	var buffer [1]byte
	n, err := os.Stdin.Read(buffer[:])
	if err != nil || n == 0 {
		panic("Error in stdin")
	}
	return buffer[0]
}

/*
 * OUTPUT DEVICE
 */
type OutputDevice struct {
	*CommonDevice
}

func (d *OutputDevice) Write(value byte) {
	fmt.Printf("%c", value)
}

/*
 * ERROR DEVICE
 */
type ErrorDevice struct{
	*CommonDevice
}

func (d *ErrorDevice) Write(value byte) {
	fmt.Printf("%c", value)
}

/*
 * FILE DEVICE
 */
type FileDevice struct {
	file *os.File
}

func (d *FileDevice) Test() bool {
	if d.file == nil {
		return false
	}
	info, err := os.Stat(d.file.Name())
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (d *FileDevice) Read() byte {
	checkIfFileExists(d.file)
	buffer := make([]byte, 1)
	n, err := d.file.Read(buffer)
	if err != nil || n == 0 {
		return 0
	}
	return buffer[0]
}

func (d *FileDevice) Write(value byte) {
	checkIfFileExists(d.file)
	_, err := d.file.Write([]byte{value})
	if err != nil {
		panic("Failed to write to file: " + err.Error())
	}
}

func checkIfFileExists(file *os.File) {
	if file == nil {
		fileDoesNotExist()
	}
}