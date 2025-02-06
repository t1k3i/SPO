package machine

const NUM_OF_ADDRESES = 1048576
const MAX_ADDRESS = 0xFFFFFF
const MIN_ADDRESS = 0

type Memory struct {
	memory [NUM_OF_ADDRESES]byte
}

func (m *Memory) GetByte(addr int32) byte {
	checkAddress(addr)
	return m.memory[addr]
}

func (m *Memory) SetByte(addr int32, v byte) {
	checkAddress(addr)
	m.memory[addr] = v
}

func (m *Memory) GetWord(addr int32) int32 {
	checkAddress(addr)
	word := int32(m.memory[addr])<<16 | int32(m.memory[addr+1])<<8 | int32(m.memory[addr+2])
	return word
}

func (m *Memory) SetWord(addr int32, v int32) {
	checkAddress(addr)
	m.memory[addr+2] = byte(v)
	v >>= 8
	m.memory[addr+1] = byte(v)
	v >>= 8
	m.memory[addr] = byte(v)
}

func (m *Memory) clear() {
	zeroes := [NUM_OF_ADDRESES]byte{}
	copy(m.memory[:], zeroes[:])
}