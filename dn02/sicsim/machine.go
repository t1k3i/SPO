package main

const (
	MinInt24 = -8388608
	MaxInt24 = 8388607
)

const NUM_OF_DEVICES = 256

type Machine struct {
	regA, regX, regL, regB, regS, regT int32
	regF                               float64
	pc, sw                             int32
	mem                                Memory
	devices                            [NUM_OF_DEVICES]Device
}

/*
 *	MACHINE CONSTRUCTOR
 */
func NewMachine() *Machine {
	machine := &Machine{}
	machine.mem = Memory{}
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
		panic("Float v intu")
	case 8:
		return m.GetPC()
	case 9:
		return m.GetSW()
	default:
		// TODO
		panic("Not valid register index")
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
		panic("Float v intu")
	case 8:
		m.SetPC(v)
	case 9:
		m.SetSW(v)
	default:
		// TODO
		panic("Not valid register index")
	}
}

func CheckValue(v int32) {
	if v < MinInt24 || v > MaxInt24 {
		// TODO
		panic("Registers are 24 bits")
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

func CheckDeviceNumber(num int) {
	if num < 0 || num >= NUM_OF_DEVICES {
		panic("Invalid device number")
	}
}