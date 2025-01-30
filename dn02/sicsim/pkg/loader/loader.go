package loader

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"sicsim/pkg/machine"
)

func readFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	return file, nil
}

func readByte(file *os.File, numOfBytes int) ([]byte, error) {
	b := make([]byte, numOfBytes)
	_, err := file.Read(b)
	if err != nil {
		return nil, errors.New("failed to read file")
	}
	return b, nil
}

func hexToInt(hexBytes []byte) (int32, error) {
	hexString := string(hexBytes)
	parsedInt, err := strconv.ParseInt(hexString, 16, 32)
	if err != nil {
		return 0, errors.New("invalid object file format")
	}
	return int32(parsedInt), nil
}

func readHeader(file *os.File) (int32, int32, error) {
	recordLetter, err := readByte(file, 1)
	if err != nil || recordLetter[0] != byte('H') {
		return 0, 0, errors.New("invalid object file: missing header")
	}
	if _, err := readByte(file, 6); err != nil {
		return 0, 0, err
	}

	loadAddressBytes, err := readByte(file, 6)
	if err != nil {
		return 0, 0, err
	}
	loadAddress, err := hexToInt(loadAddressBytes)
	if err != nil {
		return 0, 0, err
	}

	codeLengthBytes, err := readByte(file, 6)
	if err != nil {
		return 0, 0, err
	}
	codeLength, err := hexToInt(codeLengthBytes)
	if err != nil {
		return 0, 0, err
	}

	return loadAddress, codeLength, nil
}

func readEnd(file *os.File) (int32, error) {
	loadStartBytes, err := readByte(file, 6)
	if err != nil {
		return 0, err
	}
	return hexToInt(loadStartBytes)
}

// Gets the code bytes as a list of separate bytes
func getBytes(bytes []byte) ([]byte, error) {
	code := []byte{}
	l := len(bytes)
	for i := 0; i < l; i += 2 {
		if i+2 > l {
			return nil, errors.New("invalid object file format")
		}
		b, err := hexToInt(bytes[i : i+2])
		if err != nil {
			return nil, err
		}
		code = append(code, byte(b))
	}
	return code, nil
}

func readTrecord(file *os.File) (int32, int32, []byte, error) {
	codeAddressBytes, err := readByte(file, 6)
	if err != nil {
		return 0, 0, nil, err
	}
	codeAddress, err := hexToInt(codeAddressBytes)
	if err != nil {
		return 0, 0, nil, err
	}

	codeLenBytes, err := readByte(file, 2)
	if err != nil {
		return 0, 0, nil, err
	}
	codeLen, err := hexToInt(codeLenBytes)
	if err != nil {
		return 0, 0, nil, err
	}

	codeChars, err := readByte(file, int(codeLen)*2)
	if err != nil {
		return 0, 0, nil, err
	}
	code, err := getBytes(codeChars)
	if err != nil {
		return 0, 0, nil, err
	}

	if _, err := readByte(file, 1); err != nil {
		return 0, 0, nil, err
	}

	return codeAddress, codeLen, code, nil
}

func loadInMemory(codeAddress int32, codeLen int32, code []byte, m *machine.Machine) {
	endAddr := codeAddress + codeLen
	for i := codeAddress; i < endAddr; i++ {
		m.SetByte(i, code[i - codeAddress])
	}
}

func Load(filePath string, m *machine.Machine) error {
	file, err := readFile(filePath)
	if (err != nil) {
		return err
	}
	defer file.Close()

	loadAddress, codeLength, err := readHeader(file)
	if err != nil {
		return err
	}
	fmt.Println(loadAddress, "|", codeLength)

	if _, err := readByte(file, 1); err != nil {
		return err
	}

	for {
		recordLetter, err := readByte(file, 1)
		if err != nil {
			return err
		}
		if recordLetter[0] == byte('D') || recordLetter[0] == byte('R') {
			return errors.New("unsupported record type")
		}
		if recordLetter[0] == byte('E') {
			break
		}

		codeAddress, codeLen, code, err := readTrecord(file)
		if err != nil {
			return err
		}
		loadInMemory(codeAddress, codeLen, code, m)

		fmt.Print(codeAddress, " | ", codeLen, " | ")
		fmt.Printf("%X\n", code)
	}

	pcStart, err := readEnd(file)
	if err != nil {
		return err
	}
	fmt.Println(pcStart)
	m.SetPC(pcStart)

	return nil
}