package emulator

import (
    "fmt"
    "github.com/kierdavis/avr"
    "github.com/kierdavis/avr/spec"
)

// Returns 1 if x is true, else 0.
func b2i(x bool) uint8 {
    if x {
        return 1
    }
    return 0
}

func setFlagIfEqual(dest *uint8, x, y uint8) {
    if x == y {
        *dest = 1
    } else {
        *dest = 0
    }
}

func setFlagIfEqual16(dest *uint8, x, y uint16) {
    if x == y {
        *dest = 1
    } else {
        *dest = 0
    }
}

func andFlagIfEqual(dest *uint8, x, y uint8) {
    if x != y {
        *dest = 0
    }
}

// Warnings include events such as
// * invalid instructions
// * instructions not available on this particular MCU
// * accesses to an unmapped data memory address
// which are ignored by a real MCU but are often indicative of software errors.
type Warning interface {
    String() string
}

// An attempt to load from/store to an address in data memory outside of the
// ranges specified in the MCUSpec.
type UnmappedAddressWarning struct {
    PC      uint32
    Address uint16
}

func (w UnmappedAddressWarning) String() string {
    return fmt.Sprintf("access of unmapped data memory address $%04X (at PC $%06X)", w.Address, w.PC)
}

// An attempt to load from/store to an I/O port address that is not mapped by
// the MCUSpec.
type UnmappedPortWarning struct {
    PC       uint32
    BankNum  uint
    PortAddr uint16
}

func (w UnmappedPortWarning) String() string {
    return fmt.Sprintf("access of unmapped I/O port at address $%04X in I/O bank %d (at PC $%06X)", w.PortAddr, w.BankNum, w.PC)
}

// An invalid instruction word.
type InvalidInstructionWarning struct {
    PC          uint32
    Word uint16
}

func (w InvalidInstructionWarning) String() string {
    return fmt.Sprintf("invalid instruction word $%04X (at PC $%06X)", w.Word, w.PC)
}

// An instruction that is unsupported by this MCU.
type UnavailableInstructionWarning struct {
    PC          uint32
    Instruction avr.Instruction
    MCUSpec     *spec.MCUSpec
}

func (w UnavailableInstructionWarning) String() string {
    return fmt.Sprintf("instruction %s is not available on %s (at PC $%06X)", w.Instruction, w.MCUSpec.Label, w.PC)
}
