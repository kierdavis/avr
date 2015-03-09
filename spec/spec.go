package spec

import (
	"github.com/kierdavis/avr"
)

// An MCUFamily identifies the capabilities of an 8-bit MCU, with respect to
// which instructions are supported.
// See http://en.wikipedia.org/wiki/Atmel_AVR_instruction_set#Instruction_set_inheritance
type MCUFamily int
const (
    ReducedCore MCUFamily = iota
    MinimalCore
    ClassicCore8K
    ClassicCore128K
    EnhancedCore8K
    EnhancedCore128K
    EnhancedCore4M
    XMEGA
)

// An MCUSpec is a specification of a particular AVR variant.
type MCUSpec struct {
	Label          string
	Family         MCUFamily
	NumRegs        uint
	LogProgMemSize uint
	LogRAMSize     uint
	LogEEPROMSize  uint
	IOBankSizes    []uint
	Regions        []RegionSpec
	Ports          map[string]avr.PortRef
	Available      [avr.NumInstructions]bool
}

// A RegionSpec is a specification of a region of data memory pertaining to a
// particular AVR variant.
type RegionSpec interface {
	Start() uint16
	Size() uint16
}

// A RegsRegionSpec is a specification of a mapping of the register file to the
// data memory space.
type RegsRegionSpec struct {
	mcuSpec *MCUSpec
	start   uint16
}

func (r RegsRegionSpec) Start() uint16 {
	return r.start
}

func (r RegsRegionSpec) Size() uint16 {
	return uint16(r.mcuSpec.NumRegs)
}

// An IORegionSpec is a specification of a mapping of an IO bank to the data
// memory space.
type IORegionSpec struct {
	mcuSpec *MCUSpec
	start   uint16
	bankNum uint
}

func (r IORegionSpec) BankNum() uint {
	return r.bankNum
}

func (r IORegionSpec) Start() uint16 {
	return r.start
}

func (r IORegionSpec) Size() uint16 {
	return uint16(r.mcuSpec.IOBankSizes[r.bankNum])
}

// A RAMRegionSpec is a specification of a mapping of the RAM to the data memory
// space.
type RAMRegionSpec struct {
	mcuSpec *MCUSpec
	start   uint16
}

func (r RAMRegionSpec) Start() uint16 {
	return r.start
}

func (r RAMRegionSpec) Size() uint16 {
	return 1 << r.mcuSpec.LogRAMSize
}

// Utility function used in definitions of MCUSpecs. Assigns the mcuSpec field
// of all RegionSpecs.
func linkRegions(s *MCUSpec) *MCUSpec {
	for i, r := range s.Regions {
		switch r_ := r.(type) {
		case RegsRegionSpec:
			r_.mcuSpec = s
		case IORegionSpec:
			r_.mcuSpec = s
		case RAMRegionSpec:
			r_.mcuSpec = s
		}
		s.Regions[i] = r
	}

	return s
}
