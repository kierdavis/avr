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

type PortRef struct {
    BankNum uint
    Index   uint16
}

