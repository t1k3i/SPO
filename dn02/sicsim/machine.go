package main

import "fmt"

const (
	MinInt24 = -8388608
	MaxInt24 = 8388607
)

const (
	LT = -1
	EQ = 0
	GT = 1
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
func (m *Machine) SetReg(reg byte, v int32) {
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
		m.execute()
	}
}

func (m *Machine) execute() {
	opcode := m.fetch()
	if m.execF1(opcode) {
		return
	}
	op := int32(m.fetch())
	if m.execF2(opcode, byte(op)) {
		return
	}
	ni := opcode & 3
	xbpe := byte(op >> 4)
	op = op << 8
	op += int32(m.fetch())
	// Extended format
	if (xbpe & 1) == 1 {
		op = op << 8
		op += int32(m.fetch())
		if m.execF3F4(opcode, op, ni, true) {
			return
		}
	} else {
		if m.execF3F4(opcode, op, ni, false) {
			return
		}
	}
	
	OpcodeNotValid(opcode)
}

func (m *Machine) execF1(opcode byte) bool {
	switch opcode {
	case FIX:
		NotImplemented()
	case FLOAT:
		NotImplemented()
	case HIO:
		NotImplemented()
	case NORM:
		NotImplemented()
	case SIO:
		NotImplemented()
	case TIO:
		NotImplemented()
	default:
		return false
	}
	return true
}

/*
 *	F2 COMMANDS
 */
 func (m *Machine) execF2(opcode byte, op byte) bool {
	handlers := map[byte]func(byte){
		ADDR:   m.addr,
		CLEAR:  m.clear,
		COMPR:  m.compr,
		DIVR:   m.divr,
		MULR:   m.mulr,
		RMO:    m.rmo,
		SHIFTL: m.shiftl,
		SHIFTR: m.shiftr,
		SUBR:   m.subr,
		SVC:    m.svc,
		TIXR:   m.tixr,
	}

	if handler, ok := handlers[opcode]; ok {
		handler(op)
		return true
	}
	return false
}

func (m *Machine) addr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r1) + m.GetReg(r2))
}

func (m *Machine) clear(op byte) {
	m.SetReg(op, 0)
}

func (m *Machine) compr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	r1Value := m.GetReg(r1)
	r2Value := m.GetReg(r2)
	if r1Value < r2Value {
		m.SetSW(LT)
	} else if r1Value == r2Value {
		m.SetSW(EQ)
	} else {
		m.SetSW(GT)
	}
}

func (m *Machine) divr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r1) / m.GetReg(r2))
}

func (m *Machine) mulr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r1) * m.GetReg(r2))
}

func (m *Machine) rmo(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r1))
}

func (m *Machine) shiftl(op byte) {
	r1 := op & 0b11110000
	v := op & 0b00001111
	m.SetReg(r1, m.GetReg(r1) << v)
}

func (m *Machine) shiftr(op byte) {
	r1 := op & 0b11110000
	v := op & 0b00001111
	m.SetReg(r1, m.GetReg(r1) >> v)
}

func (m *Machine) subr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r2) - m.GetReg(r1))
}

func (m *Machine) svc(op byte) {
	NotImplemented()
}

func (m *Machine) tixr(op byte) {
	m.SetX(m.GetX() + 1)
	r1 := 0b00010000 // register x
	r2 := op // register from operand
	op2 := byte(r1) + r2
	m.compr(op2)
}

/*
 *	F3F4 COMMANDS
 */
 func (m *Machine) execF3F4(opcode byte, op int32, ni byte, ex bool) bool {
	handlers := map[byte]func(int32, bool, bool){
		ADD:   m.add,
		ADDF:  m.addf,
		AND:   m.and,
		COMP:  m.comp,
		COMPF: m.compf,
		DIV:   m.div,
		DIVF:  m.divf,
		J:     m.j,
		JEQ:   m.jeq,
		JGT:   m.jgt,
		JLT:   m.jlt,
		JSUB:  m.jsub,
		LDA:   m.lda,
		LDB:   m.ldb,
		LDCH:  m.ldch,
		LDF:   m.ldf,
		LDL:   m.ldl,
		LDS:   m.lds,
		LDT:   m.ldt,
		LDX:   m.ldx,
		LPS:   m.lps,
		MUL:   m.mul,
		MULF:  m.mulf,
		OR:    m.or,
		RD:    m.rd,
		RSUB:  m.rsub,
		SSK:   m.ssk,
		STA:   m.sta,
		STB:   m.stb,
		STCH:  m.stch,
		STF:   m.stf,
		STI:   m.sti,
		STL:   m.stl,
		STS:   m.sts,
		STSW:  m.stsw,
		STT:   m.stt,
		STX:   m.stx,
		SUB:   m.sub,
		SUBF:  m.subf,
		TD:    m.td,
		TIX:   m.tix,
		WD:    m.wd,
	}

	operand, old := m.getFullOperandAndCheckIfOld(op, ni, ex)

	if handler, ok := handlers[opcode]; ok {
		handler(operand, ex, old)
		return true
	}
	return false
}

func (m *Machine) add(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() + op)
}

func (m *Machine) addf(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) and(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() & op)
}

func (m *Machine) comp(op int32, ex bool, oldSic bool) {
	valueA := m.GetA()
	if valueA < op {
		m.SetSW(LT)
	} else if valueA == op {
		m.SetSW(EQ)
	} else {
		m.SetSW(GT)
	}
}

func (m *Machine) compf(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) div(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() / op)
}

func (m *Machine) divf(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) j(op int32, ex bool, oldSic bool) {
	m.SetPC(op)
}

func (m *Machine) jeq(op int32, ex bool, oldSic bool) {
	if m.GetSW() == EQ {
		m.SetPC(op)
	}
}

func (m *Machine) jgt(op int32, ex bool, oldSic bool) {
	if m.GetSW() == GT {
		m.SetPC(op)
	}
}

func (m *Machine) jlt(op int32, ex bool, oldSic bool) {
	if m.GetSW() == LT {
		m.SetPC(op)
	}
}

func (m *Machine) jsub(op int32, ex bool, oldSic bool) {
	m.SetL(m.GetPC())
	m.SetPC(op)
}

func (m *Machine) lda(op int32, ex bool, oldSic bool) {
	m.SetA(op)
}

func (m *Machine) ldb(op int32, ex bool, oldSic bool) {
	m.SetB(op)
}

func (m *Machine) ldch(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) ldf(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) ldl(op int32, ex bool, oldSic bool) {
	m.SetL(op)
}

func (m *Machine) lds(op int32, ex bool, oldSic bool) {
	m.SetS(op)
}

func (m *Machine) ldt(op int32, ex bool, oldSic bool) {
	m.SetT(op)
}

func (m *Machine) ldx(op int32, ex bool, oldSic bool) {
	m.SetX(op)
}

func (m *Machine) lps(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) mul(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() * op)
}

func (m *Machine) mulf(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) or(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() | op)
}

func (m *Machine) rd(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) rsub(op int32, ex bool, oldSic bool) {
	m.SetPC(m.GetL())
}

func (m *Machine) ssk(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) sta(op int32, ex bool, oldSic bool) {
	m.SetWord(op, m.GetA())
}

func (m *Machine) stb(op int32, ex bool, oldSic bool) {
	m.SetWord(op, m.GetB())
}

func (m *Machine) stch(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) stf(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) sti(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) stl(op int32, ex bool, oldSic bool) {
	m.SetWord(op, m.GetL())
}

func (m *Machine) sts(op int32, ex bool, oldSic bool) {
	m.SetWord(op, m.GetS())
}

func (m *Machine) stsw(op int32, ex bool, oldSic bool) {
	m.SetWord(op, m.GetSW())
}

func (m *Machine) stt(op int32, ex bool, oldSic bool) {
	m.SetWord(op, m.GetT())
}

func (m *Machine) stx(op int32, ex bool, oldSic bool) {
	m.SetWord(op, m.GetX())
}

func (m *Machine) sub(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() - op)
}

func (m *Machine) subf(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) td(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) tix(op int32, ex bool, oldSic bool) {
	m.SetX(m.GetX() + 1)
	xValue := m.GetX()
	if xValue < op {
		m.SetSW(LT)
	} else if xValue == op {
		m.SetSW(EQ)
	} else {
		m.SetSW(GT)
	}
}

func (m *Machine) wd(op int32, ex bool, oldSic bool) {
	NotImplemented()
}

func (m *Machine) getFullOperandAndCheckIfOld(op int32, ni byte, ex bool) (int32, bool) {
	if ni == 0 {
		if !ex {
			panic("Old SIC and extended!")
		}
		if (op & 0x8000) == 1 {
			return m.GetWord((op & 0x7FFF) + m.GetX()), true
		}
		return m.GetWord(op & 0x7FFF), true
	} else {
		offset, xbpe := getOffsetAndXBPE(op, ex)
		UN := m.getEffectiveAddress(xbpe, offset)
		return m.getFullOperand(UN, ni), false
	}
}

func getOffsetAndXBPE(op int32, ex bool) (int32, byte) {
	if ex {
		return op & 0x0FFFFF, byte((op & 0xF00000) >> 20)
	} else {
		return op & 0x0FFF, byte((op & 0xF000) >> 12)
	}
}

func (m *Machine) getEffectiveAddress(xbpe byte, offset int32) int32 {
	var x int32 = 0
	if (xbpe & 8) == 1 {
		x = m.GetX()
	}
	if (xbpe & 2) == 1 {
		return m.GetPC() + offset + x
	} else if (xbpe & 4) == 1 {
		return m.GetB() + offset + x
	} else if ((xbpe & 2) == 0) && ((xbpe & 4) == 0) {
		return offset + x
	} else {
		panic("This type of addressing is not supported!")
	}
}

func (m *Machine) getFullOperand(UN int32, ni byte) int32 {
	if ni == 1 {
		return UN
	} else if ni == 2 {
		return m.GetWord(m.GetWord(UN))
	} else {
		return m.GetWord(UN)
	}
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