package main

import "fmt"

const (
	MinInt24 = -8388608
	MaxInt24 = 8388607
)

const NUM_OF_DEVICES = 256

type Machine struct {
	regA, regX, regL, regB, regS, regT int32
	regF                               float64
	pc, sw                             int32
	Memory
	devices [NUM_OF_DEVICES]Device
}

/*
 *	MACHINE CONSTRUCTOR
 */
func NewMachine() *Machine {
	machine := &Machine{Memory: Memory{}}
	machine.devices[0] = &InputDevice{}
	machine.devices[1] = &OutputDevice{}
	machine.devices[2] = &ErrorDevice{}
	return machine
}

/*
 *	GET REGISTER VALUES
 */
func (m *Machine) GetA() int32 {
	return m.regA
}

func (m *Machine) GetX() int32 {
	return m.regX
}

func (m *Machine) GetL() int32 {
	return m.regL
}

func (m *Machine) GetB() int32 {
	return m.regB
}

func (m *Machine) GetS() int32 {
	return m.regS
}

func (m *Machine) GetT() int32 {
	return m.regT
}

func (m *Machine) GetF() float64 {
	return m.regF
}

func (m *Machine) GetPC() int32 {
	return m.pc
}

func (m *Machine) GetSW() int32 {
	return m.sw
}

/*
 *	SET REGISTER VALUES
 */
func (m *Machine) SetA(v int32) {
	CheckValue(v)
	m.regA = v
}

func (m *Machine) SetX(v int32) {
	CheckValue(v)
	m.regX = v
}

func (m *Machine) SetL(v int32) {
	CheckValue(v)
	m.regL = v
}

func (m *Machine) SetB(v int32) {
	CheckValue(v)
	m.regB = v
}

func (m *Machine) SetS(v int32) {
	CheckValue(v)
	m.regS = v
}

func (m *Machine) SetT(v int32) {
	CheckValue(v)
	m.regT = v
}

func (m *Machine) SetF(v float64) {
	m.regF = v
}

func (m *Machine) SetPC(v int32) {
	CheckValue(v)
	m.pc = v
}

func (m *Machine) IncPC() {
	m.pc++
	CheckValue(m.pc)
}

func (m *Machine) SetSW(v int32) {
	CheckValue(v)
	m.sw = v
}

/*
 *	GET REGISTER VALUES BY INDEX
 */
func (m *Machine) GetReg(reg int) int32 {
	switch reg {
	case 0:
		return m.GetA()
	case 1:
		return m.GetX()
	case 2:
		return m.GetL()
	case 3:
		return m.GetB()
	case 4:
		return m.GetS()
	case 5:
		return m.GetT()
	case 6:
		// TODO
		panic("Float v intu")
	case 8:
		return m.GetPC()
	case 9:
		return m.GetSW()
	default:
		NotValidRegisterIndex()
		panic("Unreachable code after panic")
	}
}

/*
 *	SET REGISTER VALUES BY INDEX
 */
func (m *Machine) SetReg(reg int, v int32) {
	CheckValue(v)
	switch reg {
	case 0:
		m.SetA(v)
	case 1:
		m.SetX(v)
	case 2:
		m.SetL(v)
	case 3:
		m.SetB(v)
	case 4:
		m.SetS(v)
	case 5:
		m.SetT(v)
	case 6:
		// TODO
		panic("Float v intu")
	case 8:
		m.SetPC(v)
	case 9:
		m.SetSW(v)
	default:
		NotValidRegisterIndex()
	}
}

/*
 *	DEVICE FUNCTIONS
 */
func (m *Machine) GetDevice(num int) Device {
	CheckDeviceNumber(num)
	return m.devices[num]
}

func (m *Machine) SetDevice(num int, device Device) {
	CheckDeviceNumber(num)
	m.devices[num] = device
}

/*
 *	MAIN LOOP
 */
func (m *Machine) Start() {
	for {
		opcode := m.fetch()

		switch opcode {
		case ADD:
			NotImplemented()
		case ADDF:
			NotImplemented()
		case ADDR:
			NotImplemented()
		case AND:
			NotImplemented()
		case CLEAR:
			NotImplemented()
		case COMP:
			NotImplemented()
		case COMPF:
			NotImplemented()
		case COMPR:
			NotImplemented()
		case DIV:
			NotImplemented()
		case DIVF:
			NotImplemented()
		case DIVR:
			NotImplemented()
		case FIX:
			NotImplemented()
		case FLOAT:
			NotImplemented()
		case HIO:
			NotImplemented()
		case J:
			NotImplemented()
		case JEQ:
			NotImplemented()
		case JGT:
			NotImplemented()
		case JLT:
			NotImplemented()
		case JSUB:
			NotImplemented()
		case LDA:
			NotImplemented()
		case LDB:
			NotImplemented()
		case LDCH:
			NotImplemented()
		case LDF:
			NotImplemented()
		case LDL:
			NotImplemented()
		case LDS:
			NotImplemented()
		case LDT:
			NotImplemented()
		case LDX:
			NotImplemented()
		case LPS:
			NotImplemented()
		case MUL:
			NotImplemented()
		case MULF:
			NotImplemented()
		case MULR:
			NotImplemented()
		case NORM:
			NotImplemented()
		case OR:
			NotImplemented()
		case RD:
			NotImplemented()
		case RMO:
			NotImplemented()
		case RSUB:
			NotImplemented()
		case SHIFTL:
			NotImplemented()
		case SHIFTR:
			NotImplemented()
		case SIO:
			NotImplemented()
		case SSK:
			NotImplemented()
		case STA:
			NotImplemented()
		case STB:
			NotImplemented()
		case STCH:
			NotImplemented()
		case STF:
			NotImplemented()
		case STI:
			NotImplemented()
		case STL:
			NotImplemented()
		case STS:
			NotImplemented()
		case STSW:
			NotImplemented()
		case STT:
			NotImplemented()
		case STX:
			NotImplemented()
		case SUB:
			NotImplemented()
		case SUBF:
			NotImplemented()
		case SUBR:
			NotImplemented()
		case SVC:
			NotImplemented()
		case TD:
			NotImplemented()
		case TIO:
			NotImplemented()
		case TIX:
			NotImplemented()
		case TIXR:
			NotImplemented()
		case WD:
			NotImplemented()
		default:
			OpcodeNotValid(opcode)
		}
	}
}

func (m *Machine) execute() {
	opcode := m.fetch()
	if m.execF1(opcode) {
		return
	}
	op := int(m.fetch())
	if m.execF2(opcode, byte(op)) {
		return
	}
	op = op << 8
	op += int(m.fetch())
	ni := opcode & 3
	if m.execF3F4(opcode, op, ni) {
		return
	}
	OpcodeNotValid(opcode)
}

func (m *Machine) execF1(opcode byte) bool {
	return false
}

func (m *Machine) execF2(opcode byte, op byte) bool {
	return false
}

func (m *Machine) execF3F4(opcode byte, op int, ni byte) bool {
	return false
}

func (m *Machine) fetch() byte {
	ret := m.GetByte(m.GetPC())
	m.IncPC()
	return ret
}

/*
 *	ERRORS
 */
func CheckValue(v int32) {
	if v < MinInt24 || v > MaxInt24 {
		panic("Values in SIC can be up to 24 bits!")
	}
}

func CheckDeviceNumber(num int) {
	if num < 0 || num >= NUM_OF_DEVICES {
		panic("Invalid device number!")
	}
}

func NotValidRegisterIndex() {
	panic("Not valid register index!")
}

func NotImplemented() {
	panic("Not implemented!")
}

func OpcodeNotValid(opcode byte) {
	str := fmt.Sprintf("Operation code %d is not valid!", opcode)
	panic(str)
}

func InvalidAddressing() {
	panic("Invalid addressing!")
}