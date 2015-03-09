package emulator

import (
	"github.com/kierdavis/avr"
)

// A Port encapsulates the interface to a particular I/O port.
type Port interface {
	Read() uint8
	Write(uint8)
}

// SregPort implements the SREG (status register) I/O port. It is automatically
// registered upon creation of an Emulator.
type SregPort struct {
	em *Emulator
}

func (p SregPort) Read() uint8 {
	return (p.em.flags[avr.FlagI] << 7) |
		   (p.em.flags[avr.FlagT] << 6) |
		   (p.em.flags[avr.FlagH] << 5) |
		   (p.em.flags[avr.FlagS] << 4) |
		   (p.em.flags[avr.FlagV] << 3) |
		   (p.em.flags[avr.FlagN] << 2) |
		   (p.em.flags[avr.FlagZ] << 1) |
		    p.em.flags[avr.FlagC]
}

func (p SregPort) Write(x uint8) {
	p.em.flags[avr.FlagI] = (x >> 7) & 1
	p.em.flags[avr.FlagT] = (x >> 6) & 1
	p.em.flags[avr.FlagH] = (x >> 5) & 1
	p.em.flags[avr.FlagS] = (x >> 4) & 1
	p.em.flags[avr.FlagV] = (x >> 3) & 1
	p.em.flags[avr.FlagN] = (x >> 2) & 1
	p.em.flags[avr.FlagZ] = (x >> 1) & 1
	p.em.flags[avr.FlagC] = x & 1
}

// SphPort implements the SPH (stack pointer high byte) I/O port. It is
// automatically registered upon creation of an Emulator.
type SphPort struct {
	em *Emulator
}

func (p SphPort) Read() uint8 {
	return uint8(p.em.sp >> 8)
}

func (p SphPort) Write(x uint8) {
	p.em.sp = (p.em.sp & 0x00FF) | (uint16(x) << 8)
}

// SplPort implements the SPL (stack pointer low byte) I/O port. It is
// automatially registered upon creation of an Emulator.
type SplPort struct {
	em *Emulator
}

func (p SplPort) Read() uint8 {
	return uint8(p.em.sp)
}

func (p SplPort) Write(x uint8) {
	p.em.sp = (p.em.sp & 0xFF00) | uint16(x)
}
