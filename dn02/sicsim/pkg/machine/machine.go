package machine

import (
	"fmt"
	"time"
)

const (
	MinInt24 = -8388608
	MaxInt24 = 8388607
)

const (
	LT = 0x40
	EQ = 0x00
	GT = 0x80
)

const NUM_OF_DEVICES = 256

type memChange struct {
	address int32
	value 	byte
}

type changeLog struct {
	regA, regX, regL, regB, regS, regT int32
	regF                               float64
	pc, sw							   int32
	changedMem 						   []memChange
}

type Machine struct {
	regA, regX, regL, regB, regS, regT int32
	regF                               float64
	pc, sw                             int32
	Memory
	devices 						   [NUM_OF_DEVICES]Device
	halted 							   bool
	paused 							   bool
	speed 						       time.Duration
	undoStack 						   []changeLog
}

/*
 *	MACHINE CONSTRUCTOR
 */
func NewMachine() *Machine {
	machine := &Machine{Memory: Memory{}}
	machine.devices[0] = &InputDevice{}
	machine.devices[1] = &OutputDevice{}
	machine.devices[2] = &ErrorDevice{}
	machine.speed = time.Millisecond * 1000
	machine.paused = true
	return machine
}

/*
 *	DEVICE FUNCTIONS
 */
func (m *Machine) GetDevice(num int) Device {
	checkDeviceNumber(num)
	return m.devices[num]
}

func (m *Machine) SetDevice(num int, device Device) {
	checkDeviceNumber(num)
	m.devices[num] = device
}

/*
 *	MAIN LOOP
 */
func (m *Machine) SetPaused(paused bool) {
	m.paused = paused
}

func (m *Machine) IsRunning() bool {
	return !m.halted && !m.paused
}

func (m *Machine) IsHalted() bool {
	return !m.halted
} 

func (m *Machine) Start() {
	m.paused = false;
	ticker := time.NewTicker(m.speed)
	defer ticker.Stop()
	for !m.halted && !m.paused {
		<- ticker.C
		m.PrintRegisters()
		m.PrintMEM(20)
		m.Step()
	}
}

func (m *Machine) Step() {
	if m.halted {
		return
	}
	m.execute()
}

func (m *Machine) Stop() {
	m.paused = true;
}

func (m *Machine) Undo() {
	if len(m.undoStack) == 0 {
		return
	}

	m.halted = false

	ix := len(m.undoStack)-1
	lastLog := m.undoStack[ix]
	m.undoStack = m.undoStack[:ix]

	m.regA = lastLog.regA
	m.regX = lastLog.regX
	m.regL = lastLog.regL
	m.regB = lastLog.regB
	m.regS = lastLog.regS
	m.regT = lastLog.regT
	m.regF = lastLog.regF
	m.pc = lastLog.pc
	m.sw = lastLog.sw

	for _, change := range lastLog.changedMem {
		m.SetByte(change.address, change.value)
	}
}

func (m *Machine) Reset() {
    m.regA, m.regX, m.regL, m.regB, m.regS, m.regT = 0, 0, 0, 0, 0, 0
    m.regF = 0.0
    m.pc, m.sw = 0, 0

    m.Memory.clear()

    for i := 0; i < NUM_OF_DEVICES; i++ {
        m.devices[i] = nil
    }
    m.devices[0] = &InputDevice{}
    m.devices[1] = &OutputDevice{}
    m.devices[2] = &ErrorDevice{}

    m.halted = false
    m.paused = true
    m.undoStack = nil
}

func (m *Machine) fetch() byte {
	ret := m.GetByte(m.GetPC())
	m.IncPC()
	return ret
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
	opcode = opcode & 0b11111100 // opcode without ni bits
	xbpe := byte(op >> 4)
	op = op << 8
	op |= int32(m.fetch())
	// Extended format
	if (xbpe & 1) == 1 {
		op = op << 8
		op += int32(m.fetch())
		if m.execF3F4(opcode, op, ni, true) {
			return
		}
	} else {
		if (opcode == J && ni == 3 && op == 0x2FFD) {
			m.SetPC(m.GetPC() - 3)
			m.halted = true;
			return;
		}
		if m.execF3F4(opcode, op, ni, false) {
			return
		}
	}
	opcodeNotValid(opcode)
}

func (m *Machine) execF1(opcode byte) bool {
	switch opcode {
	case FIX:
		notImplementedFloat()
	case FLOAT:
		notImplementedFloat()
	case HIO:
		notImplemented()
	case NORM:
		notImplementedFloat()
	case SIO:
		notImplemented()
	case TIO:
		notImplemented()
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
		m.saveStateToUndoStack(nil, nil, false, 2) // No memory changes for F2
		handler(op)
		return true
	}
	return false
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

	fetchByte := opcode == LDCH || opcode == STCH || opcode == RD || opcode == WD || opcode == TD
	isStore := opcode == STA || opcode == STB || opcode == STCH || opcode == STF ||
		opcode == STI || opcode == STL || opcode == STS || opcode == STSW ||
		opcode == STT || opcode == STX
	operand, old := m.getFullOperandAndCheckIfOld(op, ni, ex, fetchByte, isStore)

	if handler, ok := handlers[opcode]; ok {
		if opcode != SSK && opcode != STA && opcode != STB && opcode != STCH && opcode != STF &&
			opcode != STI && opcode != STL && opcode != STS && opcode != STSW && opcode != STT && opcode != STX {
				if ex {
					m.saveStateToUndoStack(nil, nil, false, 4) // Not stores instructions do not change memory
				} else {
					m.saveStateToUndoStack(nil, nil, false, 3) // Not stores instructions do not change memory
				}
		}
		handler(operand, ex, old)
		return true
	}
	return false
}

func (m *Machine) getFullOperandAndCheckIfOld(op int32, ni byte, ex bool, fetchByte bool, isStore bool) (int32, bool) {
	if ni == 0 {
		if !ex {
			invalidAddressing()
		}
		if (op & 0x8000) == 1 {
			return m.GetWord((op & 0x7FFF) + m.GetX()), true
		}
		return m.GetWord(op & 0x7FFF), true
	} else {
		offset, xbpe := getOffsetAndXBPE(op, ex)
		UN := m.getEffectiveAddress(xbpe, offset)
		fullOperand := m.getFullOperand(UN, ni, fetchByte, isStore)
		return fullOperand, false
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
	if (xbpe & 8) == 8 {
		x = m.GetX()
	}
	if (((xbpe & 2) == 1) && ((xbpe & 4) == 1) ) {
		invalidAddressing()
		panic("Not reachable!")
	} else if (xbpe & 2) == 2 {
		return m.GetPC() + offset + x
	} else if (xbpe & 4) == 4 {
		return m.GetB() + offset + x
	} else {
		return offset + x
	}
}

func (m *Machine) getFullOperand(UN int32, ni byte, fetchByte bool /* for commands like ldch */, isStore bool) int32 {
	if ni == 1 {
		if fetchByte {
			return UN & 0x0000FF
		}
		return UN
	} else if ni == 2 {
		if fetchByte {
			return int32(m.GetByte(m.GetWord(UN)))
		}
		return m.GetWord(m.GetWord(UN))
	} else {
		if fetchByte {
			if isStore {
				return UN
			}
			return int32(m.GetByte(UN))
		}
		if isStore {
			return UN
		}
		return m.GetWord(UN)
	}
}

/*
 *	PRINT STATE
 */
func (m *Machine) PrintRegisters() {
    fmt.Println("================== Registers ==================")
    fmt.Printf("A:  %08X  X: %08X  L: %08X\n", m.regA, m.regX, m.regL)
    fmt.Printf("B:  %08X  S: %08X  T: %08X\n", m.regB, m.regS, m.regT)
    fmt.Printf("F:  %X\n", m.regF)
    fmt.Printf("PC: %08X  SW: %08X\n", m.pc, m.sw)
    fmt.Println("===============================================")
}

func (m *Machine) PrintMEM(n int) {
    fmt.Println("================== Memory Dump ==================")
    for i := 0; i < n; i += 16 {
        fmt.Printf("0x%06X: ", i)
        for j := 0; j < 16 && i+j < n; j++ {
            fmt.Printf("%02X ", m.GetByte(int32(i+j)))
        }
        fmt.Println()
    }
    fmt.Println("================================================")
}

/*
 *	SPEED OF SIM
 */
func (m *Machine) GetSpeed() time.Duration {
	return m.speed
}

func (m *Machine) SetSpeed(speed time.Duration) {
	m.speed = speed
}

/*
 *	SAVE MACHINE STATE
 */
func (m *Machine) saveStateToUndoStack(address *int32, value *int32, wholeWord bool, minus int32 /* minus PC */) {
	// Save registers
	log := changeLog{}
	log.regA = m.GetA()
	log.regX = m.GetX()
	log.regL = m.GetL()
	log.regB = m.GetB()
	log.regS = m.GetT()
	log.regF = m.GetF()
	log.pc = m.GetPC() - minus
	log.sw = m.GetSW()

	if address != nil && value != nil {
		if !wholeWord {
			log.changedMem = append(log.changedMem, memChange{address: *address, value: byte(*value)})
		} else {
			log.changedMem = append(log.changedMem, memChange{address: *address, value: byte(*value >> 16)})
			log.changedMem = append(log.changedMem, memChange{address: *address+1, value: byte(*value >> 8)})
			log.changedMem = append(log.changedMem, memChange{address: *address+2, value: byte(*value)})
		}
	}

	m.undoStack = append(m.undoStack, log)
}