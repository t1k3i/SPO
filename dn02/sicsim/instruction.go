package main

/*
 *	F2 COMMANDS
 */
func (m *Machine) addr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r1)+m.GetReg(r2))
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
	m.SetReg(r2, m.GetReg(r1)/m.GetReg(r2))
}

func (m *Machine) mulr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r1)*m.GetReg(r2))
}

func (m *Machine) rmo(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r1))
}

func (m *Machine) shiftl(op byte) {
	r1 := op & 0b11110000
	v := op & 0b00001111
	m.SetReg(r1, m.GetReg(r1)<<v)
}

func (m *Machine) shiftr(op byte) {
	r1 := op & 0b11110000
	v := op & 0b00001111
	m.SetReg(r1, m.GetReg(r1)>>v)
}

func (m *Machine) subr(op byte) {
	r1 := op & 0b11110000
	r2 := op & 0b00001111
	m.SetReg(r2, m.GetReg(r2)-m.GetReg(r1))
}

func (m *Machine) svc(op byte) {
	NotImplemented()
}

func (m *Machine) tixr(op byte) {
	m.SetX(m.GetX() + 1)
	r1 := 0b00010000 // register x
	r2 := op         // register from operand
	op2 := byte(r1) + r2
	m.compr(op2)
}

/*
 *	F3F4 COMMANDS
 */
func (m *Machine) add(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() + op)
}

func (m *Machine) addf(op int32, ex bool, oldSic bool) {
	NotImplementedFloat()
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
	NotImplementedFloat()
}

func (m *Machine) div(op int32, ex bool, oldSic bool) {
	m.SetA(m.GetA() / op)
}

func (m *Machine) divf(op int32, ex bool, oldSic bool) {
	NotImplementedFloat()
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
	m.SetA((m.GetA() & 0xFFFF00) | (op & 0xFF))
}

func (m *Machine) ldf(op int32, ex bool, oldSic bool) {
	NotImplementedFloat()
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
	NotImplementedFloat()
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
	m.SetByte(op, byte(m.GetA() & 0xFF))
}

func (m *Machine) stf(op int32, ex bool, oldSic bool) {
	NotImplementedFloat()
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
	NotImplementedFloat()
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