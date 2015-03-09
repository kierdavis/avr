package emulator

import (
	"github.com/kierdavis/avr/spec"
)

type Region interface {
	Spec() spec.RegionSpec
	Load(addr uint16) uint8
	Store(addr uint16, val uint8)
}

type RegsRegion struct {
	em         *Emulator
	regionSpec spec.RegsRegionSpec
}

func (r RegsRegion) Spec() spec.RegionSpec {
	return r.regionSpec
}

func (r RegsRegion) Load(addr uint16) uint8 {
	return r.em.regs[addr-r.regionSpec.Start()]
}

func (r RegsRegion) Store(addr uint16, val uint8) {
	r.em.regs[addr-r.regionSpec.Start()] = val
}

type IORegion struct {
	em         *Emulator
	regionSpec spec.IORegionSpec
}

func (r IORegion) Spec() spec.RegionSpec {
	return r.regionSpec
}

func (r IORegion) Load(addr uint16) uint8 {
	return r.em.readPort(r.regionSpec.BankNum(), addr-r.regionSpec.Start())
}

func (r IORegion) Store(addr uint16, val uint8) {
	r.em.writePort(r.regionSpec.BankNum(), addr-r.regionSpec.Start(), val)
}

type RAMRegion struct {
	em         *Emulator
	regionSpec spec.RAMRegionSpec
}

func (r RAMRegion) Spec() spec.RegionSpec {
	return r.regionSpec
}

func (r RAMRegion) Load(addr uint16) uint8 {
	return r.em.ram[addr-r.regionSpec.Start()]
}

func (r RAMRegion) Store(addr uint16, val uint8) {
	r.em.ram[addr-r.regionSpec.Start()] = val
}
