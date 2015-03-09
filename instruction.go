package avr

type Instruction int

const (
    ADC Instruction = iota
    ADD
    ADIW
    AND
    ANDI
    ASR
    BCLR
    BLD
    BRBC
    BRBS
    BREAK
    BSET
    BST
    CALL
    CBI
    COM
    CP
    CPC
    CPI
    CPSE
    DEC
    DES
    EICALL
    EIJMP
    ELPM_R0
    ELPM
    ELPM_INC
    EOR
    FMUL
    FMULS
    FMULSU
    ICALL
    IJMP
    IN
    INC
    JMP
    LAC
    LAS
    LAT
    LD_X
    LD_X_INC
    LD_X_DEC
    LD_Y_INC
    LD_Y_DEC
    LDD_Y
    LD_Z_INC
    LD_Z_DEC
    LDD_Z
    LDI
    LDS
    LDS_SHORT
    LPM_R0
    LPM
    LPM_INC
    LSR
    MOV
    MOVW
    MUL
    MULS
    MULSU
    NEG
    NOP
    OR
    ORI
    OUT
    POP
    PUSH
    RCALL
    RET
    RETI
    RJMP
    ROR
    SBC
    SBCI
    SBI
    SBIC
    SBIS
    SBIW
    SBRC
    SBRS
    SLEEP
    
    NumInstructions
)

var instStrings = [...]string{
    "ADC",
    "ADD",
    "ADIW",
    "AND",
    "ANDI",
    "ASR",
    "BCLR",
    "BLD",
    "BRBC",
    "BRBS",
    "BREAK",
    "BSET",
    "BST",
    "CALL",
    "CBI",
    "COM",
    "CP",
    "CPC",
    "CPI",
    "CPSE",
    "DEC",
    "DES",
    "EICALL",
    "EIJMP",
    "ELPM_R0",
    "ELPM",
    "ELPM_INC",
    "EOR",
    "FMUL",
    "FMULS",
    "FMULSU",
    "ICALL",
    "IJMP",
    "IN",
    "INC",
    "JMP",
    "LAC",
    "LAS",
    "LAT",
    "LD_X",
    "LD_X_INC",
    "LD_X_DEC",
    "LD_Y_INC",
    "LD_Y_DEC",
    "LDD_Y",
    "LD_Z_INC",
    "LD_Z_DEC",
    "LDD_Z",
    "LDI",
    "LDS",
    "LDS_SHORT",
    "LPM_R0",
    "LPM",
    "LPM_INC",
    "LSR",
    "MOV",
    "MOVW",
    "MUL",
    "MULS",
    "MULSU",
    "NEG",
    "NOP",
    "OR",
    "ORI",
    "OUT",
    "POP",
    "PUSH",
    "RCALL",
    "RET",
    "RETI",
    "RJMP",
    "ROR",
    "SBC",
    "SBCI",
    "SBI",
    "SBIC",
    "SBIS",
    "SBIW",
    "SBRC",
    "SBRS",
    "SLEEP",
}

func init() {
    if len(instStrings) != int(NumInstructions) {
        panic("package avr: len(instStrings) != NumInstructions")
    }
}

func (inst Instruction) String() string {
    if inst >= 0 && inst < NumInstructions {
        return instStrings[inst]
    }
    return "<invalid>"
}

func (inst Instruction) IsTwoWord() bool {
    switch inst {
    case CALL, JMP:
        return true
    default:
        return false
    }
}
