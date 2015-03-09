package avr

// A CPU status flag.
type Flag int

const (
    // Carry
    FlagC Flag = iota
    // Zero
    FlagZ
    // Negative
    FlagN
    // Overflow
    FlagV
    // Sign
    FlagS
    // Half-carry
    FlagH
    // Stored bit
    FlagT
    // Interrupts enabled
    FlagI
)

// A reference to an I/O port. It consists of a bank number (referring to a bank
// defined in the relevant MCUSpec) and an index within the bank.
type PortRef struct {
    BankNum uint
    Index   uint16
}

