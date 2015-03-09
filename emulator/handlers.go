package emulator

import (
	"github.com/kierdavis/avr"
	"github.com/kierdavis/avr/spec"
)

type instHandler func(*Emulator, uint16) int

var handlers = [...]instHandler{
	doADC,
	doADD,
	doADIW,
	doAND,
	doANDI,
	doASR,
	doBCLR,
	doBLD,
	doBRBC,
	doBRBS,
	doBREAK,
	doBSET,
	doBST,
	doCALL,
	doCBI,
	doCOM,
	doCP,
	doCPC,
	doCPI,
	doCPSE,
	doDEC,
	doDES,
	doEICALL,
	doEIJMP,
	doELPM_R0,
	doELPM,
	doELPM_INC,
	doEOR,
	doFMUL,
	doFMULS,
	doFMULSU,
	doICALL,
	doIJMP,
	doIN,
	doINC,
	doJMP,
	doLAC,
	doLAS,
	doLAT,
	doLD_X,
	doLD_X_INC,
	doLD_X_DEC,
	doLD_Y_INC,
	doLD_Y_DEC,
	doLDD_Y,
	doLD_Z_INC,
	doLD_Z_DEC,
	doLDD_Z,
	doLDI,
}

func init() {
	if len(handlers) != int(avr.NumInstructions) {
		panic("package avr/emulator: len(handlers) != avr.NumInstructions")
	}
}

// add with carry
func doADC(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	r := ((word & 0x0200) >> 5) | (word & 0x000F)
	// get operands
	a := em.regs[d]
	b := em.regs[r]
	// compute result
	x := a + b + em.flags[avr.FlagC]
	// set flags
	c := (a & b) | (b & ^x) | (^x & a)
	v := (a & b & ^x) | (^a & ^b & x)
	em.flags[avr.FlagH] = (c & 0x08) >> 3
	em.flags[avr.FlagV] = (v & 0x80) >> 7
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] &= b2i(x == 0)
	em.flags[avr.FlagC] = (c & 0x80) >> 7
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	// store result
	em.regs[d] = x
	return 1
}

// add
func doADD(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	r := ((word & 0x0200) >> 5) | (word & 0x000F)
	// get operands
	a := em.regs[d]
	b := em.regs[r]
	// compute result
	x := a + b
	// set flags
	c := (a & b) | (b & ^x) | (^x & a)
	v := (a & b & ^x) | (^a & ^b & x)
	em.flags[avr.FlagH] = (c & 0x08) >> 3
	em.flags[avr.FlagV] = (v & 0x80) >> 7
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = (c & 0x80) >> 7
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	// store result
	em.regs[d] = x
	return 1
}

// add immediate to word
func doADIW(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := 24 + ((word & 0x0030) >> 3)
	k := ((word & 0x00C0) >> 2) | (word & 0x000F)
	// get operands
	a := (uint16(em.regs[d+1]) << 8) | uint16(em.regs[d])
	// compute result
	x := a + k
	// set flags
	v := ^a & x
	em.flags[avr.FlagV] = uint8((v & 0x8000) >> 15)
	em.flags[avr.FlagN] = uint8((x & 0x8000) >> 15)
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = uint8((^x & a & 0x8000) >> 15)
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	// store result
	em.regs[d+1] = uint8(x >> 8)
	em.regs[d] = uint8(x)
	return 2
}

// bitwise AND
func doAND(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	r := ((word & 0x0200) >> 5) | (word & 0x000F)
	// get operands
	a := em.regs[d]
	b := em.regs[r]
	// compute result
	x := a & b
	// set flags
	em.flags[avr.FlagV] = 0
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagS] = em.flags[avr.FlagN]
	// store result
	em.regs[d] = x
	return 1
}

// bitwise AND with immediate
func doANDI(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := 16 + ((word & 0x00F0) >> 4)
	k := uint8(((word & 0x0F00) >> 4) | (word & 0x000F))
	// get operands
	a := em.regs[d]
	// compute result
	x := a & k
	// set flags
	em.flags[avr.FlagV] = 0
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagS] = em.flags[avr.FlagN]
	// store result
	em.regs[d] = x
	return 1
}

// arithmetic shift right
func doASR(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	// get operands
	a := em.regs[d]
	// compute result
	x := uint8(int8(a) >> 1)
	// set flags
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = (a & 0x01) >> 7
	em.flags[avr.FlagV] = em.flags[avr.FlagN] ^ em.flags[avr.FlagC]
	em.flags[avr.FlagS] = em.flags[avr.FlagC]
	// store result
	em.regs[d] = x
	return 1
}

// clear flag bit
func doBCLR(em *Emulator, word uint16) (cycles int) {
	s := (word & 0x0070) >> 4
	em.flags[s] = 0
	return 1
}

// bit load (copy flag T to bit in register)
func doBLD(em *Emulator, word uint16) (cycles int) {
	d := (word & 0x01F0) >> 4
	b := word & 0x0007
	em.regs[d] = (em.regs[d] & ^(1 << b)) | (em.flags[avr.FlagT] << b)
	return 1
}

// branch if flag bit cleared
func doBRBC(em *Emulator, word uint16) (cycles int) {
	s := word & 0x0007
	if em.flags[s] == 0 {
		em.pc += uint32((int32(word&0x03F8) << 22) >> 25)
		return 2
	}
	return 1
}

// branch if flag bit set
func doBRBS(em *Emulator, word uint16) (cycles int) {
	s := word & 0x0007
	if em.flags[s] != 0 {
		em.pc += uint32((int32(word&0x03F8) << 22) >> 25)
		return 2
	}
	return 1
}

// breakpoint
func doBREAK(em *Emulator, word uint16) (cycles int) {
	panic("doBREAK: unimplemented")
	return 1
}

// clear flag bit
func doBSET(em *Emulator, word uint16) (cycles int) {
	s := (word & 0x0070) >> 4
	em.flags[s] = 1
	return 1
}

// bit store (copy bit in register to flag T)
func doBST(em *Emulator, word uint16) (cycles int) {
	d := (word & 0x01F0) >> 4
	b := word & 0x0007
	em.flags[avr.FlagT] = (em.regs[d] >> b) & 1
	return 1
}

// long call
func doCALL(em *Emulator, word uint16) (cycles int) {
	kh := ((word & 0x01F0) >> 3) | (word & 0x0001)
	kl := em.fetchProgWord()
	k := (uint32(kh) << 32) | uint32(kl)

	cycles = 4
	if em.Spec.Family == spec.XMEGA {
		cycles--
	}

	if em.Spec.LogProgMemSize > 16 { // pc is 3 bytes
		em.push(uint8(em.pc >> 16))
		cycles++
	}

	em.push(uint8(em.pc >> 8))
	em.push(uint8(em.pc))

	em.pc = k

	return cycles
}

// clear bit in I/O port
func doCBI(em *Emulator, word uint16) (cycles int) {
	a := (word & 0x00F8) >> 3
	b := word & 0x0007

	x := em.readPort(0, a)
	x = x & ^(1 << b)
	em.writePort(0, a, x)

	if em.Spec.Family == spec.XMEGA || em.Spec.Family == spec.ReducedCore {
		return 1
	}
	return 2
}

// complement bits
func doCOM(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	// get operand & compute result
	x := ^em.regs[d]
	// set flags
	em.flags[avr.FlagV] = 0
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = 1
	em.flags[avr.FlagS] = em.flags[avr.FlagN]
	// store result
	em.regs[d] = x
	return 1
}

// compare
func doCP(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	r := ((word & 0x0200) >> 5) | (word & 0x000F)
	// get operands
	a := em.regs[d]
	b := em.regs[r]
	// compute result
	x := a - b
	// set flags
	c := (^a & b) | (b & x) | (x & ^a)
	v := (a & ^b & ^x) | (^a & b & x)
	em.flags[avr.FlagH] = (c & 0x08) >> 3
	em.flags[avr.FlagV] = (v & 0x80) >> 7
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = (c & 0x80) >> 7
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	return 1
}

// compare with carry
func doCPC(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	r := ((word & 0x0200) >> 5) | (word & 0x000F)
	// get operands
	a := em.regs[d]
	b := em.regs[r]
	// compute result
	x := a - b - em.flags[avr.FlagC]
	// set flags
	c := (^a & b) | (b & x) | (x & ^a)
	v := (a & ^b & ^x) | (^a & b & x)
	em.flags[avr.FlagH] = (c & 0x08) >> 3
	em.flags[avr.FlagV] = (v & 0x80) >> 7
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] &= b2i(x == 0)
	em.flags[avr.FlagC] = (c & 0x80) >> 7
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	return 1
}

// compare with immedate
func doCPI(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := 16 + ((word & 0x00F0) >> 4)
	k := uint8(((word & 0x0F00) >> 4) | (word & 0x000F))
	// get operands
	a := em.regs[d]
	// compute result
	x := a - k
	// set flags
	c := (^a & k) | (k & x) | (x & ^a)
	v := (a & ^k & ^x) | (^a & k & x)
	em.flags[avr.FlagH] = (c & 0x08) >> 3
	em.flags[avr.FlagV] = (v & 0x80) >> 7
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = (c & 0x80) >> 7
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	return 1
}

// compare and skip if equal
func doCPSE(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	r := ((word & 0x0200) >> 5) | (word & 0x000F)
	// get operands
	a := em.regs[d]
	b := em.regs[r]
	// skip if equal
	if a == b {
		return em.skip() + 1
	}
	return 1 // no skip
}

// decrement
func doDEC(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := 16 + ((word & 0x00F0) >> 4)
	// get operands
	a := em.regs[d]
	// compute result
	x := a - 1
	// set flags
	em.flags[avr.FlagV] = b2i(a == 0x80)
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	// store result
	em.regs[d] = x
	return 1
}

// DES encryption
func doDES(em *Emulator, word uint16) (cycles int) {
	k := (word & 0x00F0) >> 4
	
	panic("doDES: unimplemented")
	_ = k
	
	return 1 // TODO: 2 if not preceeded by another DES
}

// extended indirect call
func doEICALL(em *Emulator, word uint16) (cycles int) {
	if em.Spec.LogProgMemSize > 16 { // pc is 3 bytes
		em.push(uint8(em.pc >> 16))
	}
	em.push(uint8(em.pc >> 8))
	em.push(uint8(em.pc))

	em.pc = (uint32(em.eind) << 16) | (uint32(em.regs[31]) << 8) | uint32(em.regs[30])

	if em.Spec.Family == spec.XMEGA {
		return 3
	} else {
		return 4
	}
}

// extended indirect jump
func doEIJMP(em *Emulator, word uint16) (cycles int) {
	em.pc = (uint32(em.eind) << 16) | (uint32(em.regs[31]) << 8) | uint32(em.regs[30])

	return 2
}

// extended load from program memory (destination implied to be r0)
func doELPM_R0(em *Emulator, word uint16) (cycles int) {
	addr := (uint32(em.rampz) << 16) | (uint32(em.regs[31]) << 8) | uint32(em.regs[30])
	x := em.prog[addr >> 1]
	
	// lowest bit of address is byte select
	if addr & 1 != 0 {
		em.regs[0] = uint8(x >> 8)
	} else {
		em.regs[0] = uint8(x)
	}
	
	return 3
}

// extended load from program memory
func doELPM(em *Emulator, word uint16) (cycles int) {
	d := (word & 0x01F0) >> 4
	
	addr := (uint32(em.rampz) << 16) | (uint32(em.regs[31]) << 8) | uint32(em.regs[30])
	x := em.prog[addr >> 1]
	
	// lowest bit of address is byte select
	if addr & 1 != 0 {
		em.regs[d] = uint8(x >> 8)
	} else {
		em.regs[d] = uint8(x)
	}
	
	return 3
}

// extended load from program memory (post-increment)
func doELPM_INC(em *Emulator, word uint16) (cycles int) {
	d := (word & 0x01F0) >> 4
	
	addr := (uint32(em.rampz) << 16) | (uint32(em.regs[31]) << 8) | uint32(em.regs[30])
	x := em.prog[addr >> 1]
	
	// lowest bit of address is byte select
	if addr & 1 != 0 {
		em.regs[d] = uint8(x >> 8)
	} else {
		em.regs[d] = uint8(x)
	}
	
	// post-increment
	addr++
	em.rampz = uint8(addr >> 16)
	em.regs[31] = uint8(addr >> 8)
	em.regs[30] = uint8(addr)
	
	return 3
}

// bitwise exclusive OR
func doEOR(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	r := ((word & 0x0200) >> 5) | (word & 0x000F)
	// get operands
	a := em.regs[d]
	b := em.regs[r]
	// compute result
	x := a ^ b
	// set flags
	em.flags[avr.FlagV] = 0
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagS] = em.flags[avr.FlagN]
	// store result
	em.regs[d] = x
	return 1
}

// unsigned fixed-point multiply
func doFMUL(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x0070) >> 4
	r := word & 0x0007
	// get operads
	a := em.regs[d]
	b := em.regs[r]
	// compute result
	x := uint16(a) * uint16(b)
	// set flags
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = uint8((x & 0x8000) >> 15)
	// perform corrective shift
	x = x << 1
	// store result
	em.regs[1] = uint8(x >> 8)
	em.regs[0] = uint8(x)
	return 2
}

// signed fixed-point multiply
func doFMULS(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x0070) >> 4
	r := word & 0x0007
	// get operads
	a := int8(em.regs[d])
	b := int8(em.regs[r])
	// compute result
	x := uint16(int16(a) * int16(b))
	// set flags
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = uint8((x & 0x8000) >> 15)
	// perform corrective shift
	x = x << 1
	// store result
	em.regs[1] = uint8(x >> 8)
	em.regs[0] = uint8(x)
	return 2
}

// fixed-point multiply with one signed operand & one unsigned operand
func doFMULSU(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x0070) >> 4
	r := word & 0x0007
	// get operads
	a := int8(em.regs[d])
	b := em.regs[r]
	// compute result
	x := uint16(int16(a) * int16(b))
	// set flags
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagC] = uint8((x & 0x8000) >> 15)
	// perform corrective shift
	x = x << 1
	// store result
	em.regs[1] = uint8(x >> 8)
	em.regs[0] = uint8(x)
	return 2
}

// indirect call
func doICALL(em *Emulator, word uint16) (cycles int) {
	cycles = 3
	
	if em.Spec.LogProgMemSize > 16 { // pc is 3 bytes
		em.push(uint8(em.pc >> 16))
		cycles++
	}
	em.push(uint8(em.pc >> 8))
	em.push(uint8(em.pc))

	em.pc = (uint32(em.regs[31]) << 8) | uint32(em.regs[30])

	if em.Spec.Family == spec.XMEGA {
		return cycles - 1
	} else {
		return cycles
	}
}

// indirect jump
func doIJMP(em *Emulator, word uint16) (cycles int) {
	em.pc = (uint32(em.regs[31]) << 8) | uint32(em.regs[30])
	return 2
}

// read I/O port
func doIN(em *Emulator, word uint16) (cycles int) {
	d := (word & 0x01F0) >> 4
	a := ((word & 0x0600) >> 5) | (word & 0x000F)
	em.regs[d] = em.readPort(0, a)
	return 1
}

// increment
func doINC(em *Emulator, word uint16) (cycles int) {
	// extract instruction fields
	d := (word & 0x01F0) >> 4
	// get operands
	a := em.regs[d]
	// compute result
	x := a - 1
	// set flags
	em.flags[avr.FlagV] = b2i(a == 0x7F)
	em.flags[avr.FlagN] = (x & 0x80) >> 7
	em.flags[avr.FlagZ] = b2i(x == 0)
	em.flags[avr.FlagS] = em.flags[avr.FlagN] ^ em.flags[avr.FlagV]
	// store result
	em.regs[d] = x
	return 1
}

// long jump
func doJMP(em *Emulator, word uint16) (cycles int) {
	kh := ((word & 0x01F0) >> 3) | (word & 0x0001)
	kl := em.fetchProgWord()
	em.pc = (uint32(kh) << 32) | uint32(kl)
	return 3
}

// load and clear
// Note: may be buggy as this instruction is not fully documented in the AVR spec.
func doLAC(em *Emulator, word uint16) (cycles int) {
	// Rd <- [Z]
	// [Z] <- [Z] & ~(old value of Rd)
	
	d := (word & 0x01F0) >> 4
	z := (uint16(em.regs[31]) << 8) | uint16(em.regs[30])
	
	x := em.regs[d]
	y := em.loadDataByte(z)
	
	em.regs[d] = y
	em.storeDataByte(z, y & ^x)
	
	return 1
}

// load and set
// Note: may be buggy as this instruction is not fully documented in the AVR spec.
func doLAS(em *Emulator, word uint16) (cycles int) {
	// Rd <- [Z]
	// [Z] <- [Z] | old value of Rd
	
	d := (word & 0x01F0) >> 4
	z := (uint16(em.regs[31]) << 8) | uint16(em.regs[30])
	
	x := em.regs[d]
	y := em.loadDataByte(z)
	
	em.regs[d] = y
	em.storeDataByte(z, y | x)
	
	return 1
}

// load and toggle
// Note: may be buggy as this instruction is not fully documented in the AVR spec.
func doLAT(em *Emulator, word uint16) (cycles int) {
	// Rd <- [Z]
	// [Z] <- [Z] ^ old value of Rd
	
	d := (word & 0x01F0) >> 4
	z := (uint16(em.regs[31]) << 8) | uint16(em.regs[30])
	
	x := em.regs[d]
	y := em.loadDataByte(z)
	
	em.regs[d] = y
	em.storeDataByte(z, y ^ x)
	
	return 1
}

// Generalisation across all LD/LDD implementations
// Params:
//   mode: one of ' ' -- unadorned LD
//                '+' -- post-increment LD
//                '-' -- pre-decrement LD
//                'd' -- additional displacement (LDD)
//   ptrLoReg: number of low register used for pointer (X => 26, Y => 28, Z => 30)
//             ptrHiReg is implied to be ptrLoReg+1 (X => 27, Y => 29, Z => 31)
//   ptrExt: reference to either em.rampx, em.rampy or em.rampz, depending on the pointer used.
func doGenericLoad(em *Emulator, word uint16, mode byte, ptrLoReg int, ptrExt *uint8) (cycles int) {
	ptrHiReg := ptrLoReg + 1
	
	d := (word & 0x01F0) >> 4
	
	var addr uint16
	
	// Get the addr
	if em.Spec.LogDataSpaceSize > 16 {
		// Address is RAMP?:Rh:Rl
		panic("doGenericLoad: devices with a data space size > 16 not yet fully implemented")
	} else if em.Spec.LogDataSpaceSize > 8 {
		// Address is Rh:Rl
		addr = (uint16(em.regs[ptrHiReg]) << 8) | uint16(em.regs[ptrLoReg])
	} else {
		// Address is Rl
		addr = uint16(em.regs[ptrLoReg])
	}
	
	// Handle additional displacement
	if mode == 'd' {
		d := ((word & 0x2000) >> 8) | ((word & 0x0C00) >> 7) | (word & 0x0007)
		addr += d
	}
	
	// Handle pre-decrement
	if mode == '-' {
		addr--
	}
	
	// Do the load
	em.regs[d] = em.loadDataByte(addr)
	
	// Handle post-increment
	if mode == '+' {
		addr++
	}
	
	// Write back the addr if needed
	if mode == '+' || mode == '-' {
		if em.Spec.LogDataSpaceSize > 16 {
			// Address is RAMP?:Rh:Rl
			panic("doGenericLoad: devices with a data space size > 16 not yet fully implemented")
		} else if em.Spec.LogDataSpaceSize > 8 {
			// Address is Rh:Rl
			em.regs[ptrHiReg] = uint8(addr >> 8)
			em.regs[ptrLoReg] = uint8(addr)
		} else {
			// Address is Rl
			em.regs[ptrLoReg] = uint8(addr)
		}
	}
	
	// Compute number of cycles
	// This is not fully compliant with the spec, but the spec has too many
	// special cases for full compliance to be worth it.
	if em.Spec.Family == spec.XMEGA {
		if mode == '-' || mode == 'd' {
			cycles = 2
		} else {
			cycles = 1
		}
		
		return cycles
	
	} else {
		if mode == '+' {
			cycles = 2
		} else if mode == '-' {
			cycles = 3
		} else {
			cycles = 1
		}
		
		return cycles
	}
}

func doLD_X(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, ' ', 26, &em.rampx)
}

func doLD_X_INC(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, '+', 26, &em.rampx)
}

func doLD_X_DEC(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, '-', 26, &em.rampx)
}

func doLD_Y_INC(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, '+', 28, &em.rampy)
}

func doLD_Y_DEC(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, '-', 28, &em.rampy)
}

func doLDD_Y(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, 'd', 28, &em.rampy)
}

func doLD_Z_INC(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, '+', 30, &em.rampz)
}

func doLD_Z_DEC(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, '-', 30, &em.rampz)
}

func doLDD_Z(em *Emulator, word uint16) (cycles int) {
	return doGenericLoad(em, word, 'd', 30, &em.rampz)
}

func doLDI(em *Emulator, word uint16) (cycles int) {
	k := uint8(((word & 0x0F00) >> 4) | (word & 0x000F))
	d := 16 + ((word & 0x00F0) >> 4)
	em.regs[d] = k
	return 1
}
