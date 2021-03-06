package avr

//go:generate stringer -type=Instruction

type Instruction int8

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
    LD_Y
    LD_Y_INC
    LD_Y_DEC
    LDD_Y
    LD_Z
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
    SPM
    SPM_2
    ST_X
    ST_X_INC
    ST_X_DEC
    ST_Y
    ST_Y_INC
    ST_Y_DEC
    STD_Y
    ST_Z
    ST_Z_INC
    ST_Z_DEC
    STD_Z
    STS
    STS_SHORT
    SUB
    SUBI
    SWAP
    WDR
    XCH
)

const NumInstructions = int(XCH) + 1

func (inst Instruction) IsTwoWord() bool {
    switch inst {
    case CALL, JMP, LDS, STS:
        return true
    default:
        return false
    }
}
