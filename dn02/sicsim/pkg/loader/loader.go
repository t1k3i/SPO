package loader

import (
	"fmt"
	"os"
	"strconv"

	"sicsim/pkg/machine"
)

func readFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		fileErr()
	}
	return file
}

func readByte(file *os.File, numOfBytes int) []byte {
	b := make([]byte, numOfBytes)
	_, err := file.Read(b)
	if err != nil {
		panic("Something wrong")
	}
	return b
}

func hexToInt(hexBytes []byte) int32 {
	hexString := string(hexBytes)
	parsedInt, err := strconv.ParseInt(hexString, 16, 32)
	if err != nil {
		objectFileErr()
	}
	return int32(parsedInt)
}

func readHeader(file *os.File) (int32, int32) {
	recordLetter := readByte(file, 1)[0]
	if recordLetter != byte('H') {
		objectFileErr()
	}
	readByte(file, 6)

	loadAddressBytes := readByte(file, 6)
	loadAddress := hexToInt(loadAddressBytes)

	codeLengthBytes := readByte(file, 6)
	codeLength := hexToInt(codeLengthBytes)

	return loadAddress, codeLength
}

func readEnd(file *os.File) int32 {
	loadStartBytes := readByte(file, 6)
	loadStartAddress := hexToInt(loadStartBytes)
	return loadStartAddress
}

func getBytes(bytes []byte) []byte {
	code := []byte{}
	l := len(bytes)
	for i := 0; i < l; i += 2 {
		if i+2 > l {
            objectFileErr()
        }
        b := hexToInt(bytes[i : i+2])
		code = append(code, byte(b))
	}
	return code
}

func readTrecord(file *os.File) (int32, int32, []byte) {
	codeAddressBytes := readByte(file, 6)
	codeAddress := hexToInt(codeAddressBytes)

	codeLenBytes := readByte(file,2)
	codeLen := hexToInt(codeLenBytes)

	codeChars := readByte(file, int(codeLen) * 2)
	code := getBytes(codeChars)

	readByte(file, 1)
	
	return codeAddress, codeLen, code
}

func loadInMemory(codeAddress int32, codeLen int32, code []byte, m *machine.Machine) {
	endAddr := codeAddress + codeLen
	for i := codeAddress; i < endAddr; i++ {
		m.SetByte(i, code[i - codeAddress])
	}
}

func Load(filePath string, m *machine.Machine) {
	file := readFile(filePath)

	loadAddress, codeLength := readHeader(file)
	fmt.Println(loadAddress, "|", codeLength)

	readByte(file, 1)

	for {
		recordLetter := readByte(file, 1)[0]
		if recordLetter == byte('D') || recordLetter == byte('R') {
			loadingFailed()	
		}
		if recordLetter == byte('E') {
			break
		}

		codeAddress, codeLen, code := readTrecord(file)
		loadInMemory(codeAddress, codeLen, code, m)
		fmt.Print(codeAddress, " | ", codeLen, " | ")
		fmt.Printf("%X\n", code)
	}

	pcStart := readEnd(file)
	fmt.Println(pcStart)
	m.SetPC(pcStart)
}

func loadingFailed() {
	panic("Loader failed!")
}

func fileErr() {
	panic("Not able to read file!")
}

func objectFileErr() {
	panic("Object file error")
}