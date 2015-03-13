package emulator

import (
    "github.com/kierdavis/avr/spec"
)

type Region interface {
    Contains(addr uint16) bool
    Load(addr uint16) uint8
    Store(addr uint16, val uint8)
}

type RegsRegion struct {
    em         *Emulator
    regionSpec spec.RegsRegionSpec
}

func (r RegsRegion) Contains(addr uint16) bool {
    return (addr - r.regionSpec.Start()) < r.regionSpec.Size()
}

func (r RegsRegion) Load(addr uint16) uint8 {
    return r.em.regs[addr - r.regionSpec.Start()]
}

func (r RegsRegion) Store(addr uint16, val uint8) {
    r.em.regs[addr - r.regionSpec.Start()] = val
}

type IORegion struct {
    em         *Emulator
    regionSpec spec.IORegionSpec
}

func (r IORegion) Contains(addr uint16) bool {
    return (addr - r.regionSpec.Start()) < r.regionSpec.Size()
}

func (r IORegion) Load(addr uint16) uint8 {
    return r.em.readPort(r.regionSpec.BankNum(), addr - r.regionSpec.Start())
}

func (r IORegion) Store(addr uint16, val uint8) {
    r.em.writePort(r.regionSpec.BankNum(), addr - r.regionSpec.Start(), val)
}

type RAMRegion struct {
    em         *Emulator
    regionSpec spec.RAMRegionSpec
}

func (r RAMRegion) Contains(addr uint16) bool {
    return (addr - r.regionSpec.Start()) < r.regionSpec.Size()
}

func (r RAMRegion) Load(addr uint16) uint8 {
    return r.em.ram[addr - r.regionSpec.Start()]
}

func (r RAMRegion) Store(addr uint16, val uint8) {
    r.em.ram[addr - r.regionSpec.Start()] = val
}
