package emulator

import (
    "github.com/kierdavis/avr"
)

// On reduced core AVRs, the LDD and STD instructions are not present, and their
// bit patterns are used for alternate forms of LDS and STS.
func Decode(word uint16, reducedCore bool) avr.Instruction {
    switch {
    case word        == 0x0000:
        return avr.NOP
    case word&0xFF00 == 0x0100:
        return avr.MOVW
    case word&0xFF00 == 0x0200:
        return avr.MULS
    case word&0xFF88 == 0x0300:
        return avr.MULSU
    case word&0xFF88 == 0x0308:
        return avr.FMUL
    case word&0xFF88 == 0x0380:
        return avr.FMULS
    case word&0xFF88 == 0x0388:
        return avr.FMULSU
    case word&0xFC00 == 0x0400:
        return avr.CPC
    case word&0xFC00 == 0x0800:
        return avr.SBC
    case word&0xFC00 == 0x0C00:
        return avr.ADD
    case word&0xFC00 == 0x1000:
        return avr.CPSE
    case word&0xFC00 == 0x1400:
        return avr.CP
    case word&0xFC00 == 0x1C00:
        return avr.ADC
    case word&0xFC00 == 0x2000:
        return avr.AND
    case word&0xFC00 == 0x2400:
        return avr.EOR
    case word&0xFC00 == 0x2800:
        return avr.OR
    case word&0xFC00 == 0x2C00:
        return avr.MOV
    case word&0xF000 == 0x3000:
        return avr.CPI
    case word&0xF000 == 0x4000:
        return avr.SBCI
    case word&0xF000 == 0x6000:
        return avr.ORI
    case word&0xF000 == 0x7000:
        return avr.ANDI
    case word&0xD208 == 0x8000 && !reducedCore:
        return avr.LDD_Z
    case word&0xD208 == 0x8008 && !reducedCore:
        return avr.LDD_Y
    case word&0xFE0F == 0x9000 && !reducedCore:
        return avr.LDS
    case word&0xFE0F == 0x9001:
        return avr.LD_Z_INC
    case word&0xFE0F == 0x9002:
        return avr.LD_Z_DEC
    case word&0xFE0F == 0x9004:
        return avr.LPM
    case word&0xFE0F == 0x9005:
        return avr.LPM_INC
    case word&0xFE0F == 0x9006:
        return avr.ELPM
    case word&0xFE0F == 0x9007:
        return avr.ELPM_INC
    case word&0xFE0F == 0x9009:
        return avr.LD_Y_INC
    case word&0xFE0F == 0x900A:
        return avr.LD_Y_DEC
    case word&0xFE0F == 0x900C:
        return avr.LD_X
    case word&0xFE0F == 0x900D:
        return avr.LD_X_INC
    case word&0xFE0F == 0x900E:
        return avr.LD_X_DEC
    case word&0xFE0F == 0x900F:
        return avr.POP
    case word&0xFE0F == 0x920F:
        return avr.PUSH
    case word&0xFE0F == 0x9205:
        return avr.LAS
    case word&0xFE0F == 0x9206:
        return avr.LAC
    case word&0xFE0F == 0x9207:
        return avr.LAT
    case word&0xFE0F == 0x9400:
        return avr.COM
    case word&0xFE0F == 0x9401:
        return avr.NEG
    case word&0xFE0F == 0x9403:
        return avr.INC
    case word&0xFE0F == 0x9405:
        return avr.ASR
    case word&0xFE0F == 0x9406:
        return avr.LSR
    case word&0xFE0F == 0x9407:
        return avr.ROR
    case word&0xFF8F == 0x9408:
        return avr.BSET
    case word        == 0x9409:
        return avr.IJMP
    case word&0xFE0F == 0x940A:
        return avr.DEC
    case word&0xFF0F == 0x940B:
        return avr.DES
    case word&0xFE0E == 0x940C:
        return avr.JMP
    case word&0xFE0E == 0x940E:
        return avr.CALL
    case word        == 0x9419:
        return avr.EIJMP
    case word&0xFF8F == 0x9488:
        return avr.BCLR
    case word        == 0x9508:
        return avr.RET
    case word        == 0x9518:
        return avr.RETI
    case word        == 0x9509:
        return avr.ICALL
    case word        == 0x9519:
        return avr.EICALL
    case word        == 0x9598:
        return avr.BREAK
    case word        == 0x95C8:
        return avr.LPM_R0
    case word        == 0x95D8:
        return avr.ELPM_R0
    case word&0xFF00 == 0x9600:
        return avr.ADIW
    case word&0xFF00 == 0x9700:
        return avr.SBIW
    case word&0xFF00 == 0x9800:
        return avr.CBI
    case word&0xFF00 == 0x9900:
        return avr.SBIC
    case word&0xFF00 == 0x9A00:
        return avr.SBI
    case word&0xFF00 == 0x9B00:
        return avr.SBIS
    case word&0xFC00 == 0x9C00:
        return avr.MUL
    case word&0xF800 == 0xA000 && reducedCore:
        return avr.LDS_SHORT
    case word&0xF800 == 0xB000:
        return avr.IN
    case word&0xF800 == 0xB800:
        return avr.OUT
    case word&0xF000 == 0xC000:
        return avr.RJMP
    case word&0xF000 == 0xD000:
        return avr.RCALL
    case word&0xF000 == 0xE000:
        return avr.LDI
    case word&0xFC00 == 0xF000:
        return avr.BRBS
    case word&0xFC00 == 0xF400:
        return avr.BRBC
    case word&0xFE08 == 0xF800:
        return avr.BLD
    case word&0xFE08 == 0xFA00:
        return avr.BST
    case word&0xFE08 == 0xFC00:
        return avr.SBRC
    case word&0xFE08 == 0xFE00:
        return avr.SBRS
    default:
        return -1
    }
}
