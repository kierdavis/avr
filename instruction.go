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
    
    NumInstructions
)

func (inst Instruction) String() string {
    switch inst {
    case ADC:      return "ADC"
    case ADD:      return "ADD"
    case ADIW:     return "ADIW"
    case AND:      return "AND"
    case ANDI:     return "ANDI"
    case ASR:      return "ASR"
    case BCLR:     return "BCLR"
    case BLD:      return "BLD"
    case BRBC:     return "BRBC"
    case BRBS:     return "BRBS"
    case BREAK:    return "BREAK"
    case BSET:     return "BSET"
    case BST:      return "BST"
    case CALL:     return "CALL"
    case CBI:      return "CBI"
    case COM:      return "COM"
    case CP:       return "CP"
    case CPC:      return "CPC"
    case CPI:      return "CPI"
    case CPSE:     return "CPSE"
    case DEC:      return "DEC"
    case DES:      return "DES"
    case EICALL:   return "EICALL"
    case EIJMP:    return "EIJMP"
    case ELPM_R0:  return "ELPM_R0"
    case ELPM:     return "ELPM"
    case ELPM_INC: return "ELPM_INC"
    case EOR:      return "EOR"
    case FMUL:     return "FMUL"
    case FMULS:    return "FMULS"
    case FMULSU:   return "FMULSU"
    case ICALL:    return "ICALL"
    case IJMP:     return "IJMP"
    case IN:       return "IN"
    case INC:      return "INC"
    case JMP:      return "JMP"
    case LAC:      return "LAC"
    case LAS:      return "LAS"
    case LAT:      return "LAT"
    case LD_X:     return "LD_X"
    case LD_X_INC: return "LD_X_INC"
    case LD_X_DEC: return "LD_X_DEC"
    case LD_Y_INC: return "LD_Y_INC"
    case LD_Y_DEC: return "LD_Y_DEC"
    case LDD_Y:    return "LDD_Y"
    case LD_Z_INC: return "LD_Z_INC"
    case LD_Z_DEC: return "LD_Z_DEC"
    case LDD_Z:    return "LDD_Z"
    case LDI:      return "LDI"
    case LDS:      return "LDS"
    default:       return ""
    }
}

func (inst Instruction) IsTwoWord() bool {
    switch inst {
    case CALL, JMP:
        return true
    default:
        return false
    }
}
