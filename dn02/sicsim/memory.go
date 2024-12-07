package main

const NUM_OF_ADDRESES = 1048576
const MAX_ADDRESS = 0xFFFFFF
const MIN_ADDRESS = 0

type Memory struct {
	memory [NUM_OF_ADDRESES]byte
}

func (m *Memory) GetByte(addr int32) byte {
	CheckAddress(addr)
	return m.memory[addr]
}

func (m *Memory) SetByte(addr int32, v byte) {
	CheckAddress(addr)
	m.memory[addr] = v
}

func (m *Memory) GetWord(addr int32) int32 {
	CheckAddress(addr)
	word := int32(m.memory[addr])<<16 | int32(m.memory[addr+1])<<8 | int32(m.memory[addr+2])
	return word
}

func (m *Memory) SetWord(addr int32, v int32) {
	CheckAddress(addr)
	CheckValue(v)
	m.memory[addr+2] = byte(v & 0xFF)
	v >>= 8
	m.memory[addr+2] = byte(v & 0xFF)
	v >>= 8
	m.memory[addr+2] = byte(v & 0xFF)
}

/*
 *	ERRORS
 */
func CheckAddress(addr int32) {
	if addr < MIN_ADDRESS || addr > MAX_ADDRESS {
		panic("Addres not valid!")
	}
}