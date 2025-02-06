package machine

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
	v = v & 0xFFFFFF
	m.regA = v
}

func (m *Machine) SetX(v int32) {
	v = v & 0xFFFFFF
	m.regX = v
}

func (m *Machine) SetL(v int32) {
	v = v & 0xFFFFFF
	m.regL = v
}

func (m *Machine) SetB(v int32) {
	v = v & 0xFFFFFF
	m.regB = v
}

func (m *Machine) SetS(v int32) {
	v = v & 0xFFFFFF
	m.regS = v
}

func (m *Machine) SetT(v int32) {
	v = v & 0xFFFFFF
	m.regT = v
}

func (m *Machine) SetF(v float64) {
	notImplementedFloat()
}

func (m *Machine) SetPC(v int32) {
	checkValue(v)
	m.pc = v
}

func (m *Machine) IncPC() {
	m.pc++
	checkValue(m.pc)
}

func (m *Machine) SetSW(v int32) {
	v = v & 0xFFFFFF
	m.sw = v
}

/*
 *	GET REGISTER VALUES BY INDEX
 */
func (m *Machine) GetReg(reg byte) int32 {
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
		notImplementedFloat()
		panic("Unreachable code after panic")
	case 8:
		return m.GetPC()
	case 9:
		return m.GetSW()
	default:
		notValidRegisterIndex()
		panic("Unreachable code after panic")
	}
}

/*
 *	SET REGISTER VALUES BY INDEX
 */
func (m *Machine) SetReg(reg byte, v int32) {
	v = v & 0xFFFFFF
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
		notImplementedFloat()
	case 8:
		m.SetPC(v)
	case 9:
		m.SetSW(v)
	default:
		notValidRegisterIndex()
	}
}