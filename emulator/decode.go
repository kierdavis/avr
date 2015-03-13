package emulator

import (
    "github.com/kierdavis/avr"
)

// Contains 5 similar implementations of instruction decoder:
//   DecodeFull: reference implementation using a single switch statement (~46 ns/op)
//   DecodeTree: DecodeFull, but with the switch split into many smaller switches (~21 ns/op)
//   DecodeLut9: DecodeFull, but with all instructions in the range 0x9000 - 0x9FFF decoded using a 4096-entry lookup table (~30 ns/op)
//   DecodeLut9Tree: DecodeLut9, but with the switch split into many smaller switches (~18 ns/op)
//   DecodeLut: all instructions decoded using a 65536-entry (256 kB!) lookup table (~2 ns/op)

// Non-reduced core AVRs:
//   Normal (2-word) LDS/STS
//   LDD/STD with any displacement
// Reduced core AVRs:
//   Short (1-word) LDS/STS - overlaps with LDD/STD
//   LD/ST equivalent to LDD/STD with zero displacement
func DecodeFull(word uint16, reducedCore bool) avr.Instruction {
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
    case word&0xFC00 == 0x1800:
        return avr.SUB
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
    case word&0xF000 == 0x5000:
        return avr.SUBI
    case word&0xF000 == 0x6000:
        return avr.ORI
    case word&0xF000 == 0x7000:
        return avr.ANDI
    case word&0xFE0F == 0x8000 && reducedCore:
        return avr.LD_Z
    case word&0xD208 == 0x8000 && !reducedCore:
        return avr.LDD_Z
    case word&0xFE0F == 0x8008 && reducedCore:
        return avr.LD_Y
    case word&0xD208 == 0x8008 && !reducedCore:
        return avr.LDD_Y
    case word&0xFE0F == 0x8200 && reducedCore:
        return avr.ST_Z
    case word&0xD208 == 0x8200 && !reducedCore:
        return avr.STD_Z
    case word&0xFE0F == 0x8208 && reducedCore:
        return avr.ST_Y
    case word&0xD208 == 0x8208 && !reducedCore:
        return avr.STD_Y
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
    case word&0xFE0F == 0x9200 && !reducedCore:
        return avr.STS
    case word&0xFE0F == 0x920F:
        return avr.PUSH
    case word&0xFE0F == 0x9201:
        return avr.ST_Z_INC
    case word&0xFE0F == 0x9202:
        return avr.ST_Z_DEC
    case word&0xFE0F == 0x9204:
        return avr.XCH
    case word&0xFE0F == 0x9205:
        return avr.LAS
    case word&0xFE0F == 0x9206:
        return avr.LAC
    case word&0xFE0F == 0x9207:
        return avr.LAT
    case word&0xFE0F == 0x9209:
        return avr.ST_Y_INC
    case word&0xFE0F == 0x920A:
        return avr.ST_Y_DEC
    case word&0xFE0F == 0x920C:
        return avr.ST_X
    case word&0xFE0F == 0x920D:
        return avr.ST_X_INC
    case word&0xFE0F == 0x920E:
        return avr.ST_X_DEC
    case word&0xFE0F == 0x9400:
        return avr.COM
    case word&0xFE0F == 0x9401:
        return avr.NEG
    case word&0xFE0F == 0x9402:
        return avr.SWAP
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
    case word        == 0x9509:
        return avr.ICALL
    case word        == 0x9518:
        return avr.RETI
    case word        == 0x9519:
        return avr.EICALL
    case word        == 0x9588:
        return avr.SLEEP
    case word        == 0x9598:
        return avr.BREAK
    case word        == 0x95A8:
        return avr.WDR
    case word        == 0x95C8:
        return avr.LPM_R0
    case word        == 0x95D8:
        return avr.ELPM_R0
    case word        == 0x95E8:
        return avr.SPM
    case word        == 0x95F8:
        return avr.SPM_2
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
    case word&0xF800 == 0xA800 && reducedCore:
        return avr.STS_SHORT
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

func DecodeTree(word uint16, reducedCore bool) avr.Instruction {
    var masked uint16
    
    masked = word & 0xF000
    switch {
    case masked == 0x3000:
        return avr.CPI
    case masked == 0x4000:
        return avr.SBCI
    case masked == 0x5000:
        return avr.SUBI
    case masked == 0x6000:
        return avr.ORI
    case masked == 0x7000:
        return avr.ANDI
    case masked == 0xC000:
        return avr.RJMP
    case masked == 0xD000:
        return avr.RCALL
    case masked == 0xE000:
        return avr.LDI
    }
    
    masked = word & 0xF800
    switch {
    case masked == 0xA000 && reducedCore:
        return avr.LDS_SHORT
    case masked == 0xA800 && reducedCore:
        return avr.STS_SHORT
    case masked == 0xB000:
        return avr.IN
    case masked == 0xB800:
        return avr.OUT
    }
    
    if !reducedCore {
        masked = word & 0xD208
        switch {
        case masked == 0x8000:
            return avr.LDD_Z
        case masked == 0x8008:
            return avr.LDD_Y
        case masked == 0x8200:
            return avr.STD_Z
        case masked == 0x8208:
            return avr.STD_Y
        }
    }
    
    masked = word & 0xFC00
    switch {
    case masked == 0x0400:
        return avr.CPC
    case masked == 0x0800:
        return avr.SBC
    case masked == 0x0C00:
        return avr.ADD
    case masked == 0x1000:
        return avr.CPSE
    case masked == 0x1400:
        return avr.CP
    case masked == 0x1800:
        return avr.SUB
    case masked == 0x1C00:
        return avr.ADC
    case masked == 0x2000:
        return avr.AND
    case masked == 0x2400:
        return avr.EOR
    case masked == 0x2800:
        return avr.OR
    case masked == 0x2C00:
        return avr.MOV
    case masked == 0x9C00:
        return avr.MUL
    case masked == 0xF000:
        return avr.BRBS
    case masked == 0xF400:
        return avr.BRBC
    }
    
    masked = word & 0xFE08
    switch {
    case masked == 0xF800:
        return avr.BLD
    case masked == 0xFA00:
        return avr.BST
    case masked == 0xFC00:
        return avr.SBRC
    case masked == 0xFE00:
        return avr.SBRS
    }
    
    masked = word & 0xFF00
    switch {
    case masked == 0x0100:
        return avr.MOVW
    case masked == 0x0200:
        return avr.MULS
    case masked == 0x9600:
        return avr.ADIW
    case masked == 0x9700:
        return avr.SBIW
    case masked == 0x9800:
        return avr.CBI
    case masked == 0x9900:
        return avr.SBIC
    case masked == 0x9A00:
        return avr.SBI
    case masked == 0x9B00:
        return avr.SBIS
    }
    
    masked = word & 0xFE0E
    switch {
    case masked == 0x940C:
        return avr.JMP
    case masked == 0x940E:
        return avr.CALL
    }
    
    masked = word & 0xFF88
    switch {
    case masked == 0x0300:
        return avr.MULSU
    case masked == 0x0308:
        return avr.FMUL
    case masked == 0x0380:
        return avr.FMULS
    case masked == 0x0388:
        return avr.FMULSU
    }
    
    masked = word & 0xFE0F
    switch {
    case masked == 0x9000 && !reducedCore:
        return avr.LDS
    case masked == 0x9001:
        return avr.LD_Z_INC
    case masked == 0x9002:
        return avr.LD_Z_DEC
    case masked == 0x9004:
        return avr.LPM
    case masked == 0x9005:
        return avr.LPM_INC
    case masked == 0x9006:
        return avr.ELPM
    case masked == 0x9007:
        return avr.ELPM_INC
    case masked == 0x9009:
        return avr.LD_Y_INC
    case masked == 0x900A:
        return avr.LD_Y_DEC
    case masked == 0x900C:
        return avr.LD_X
    case masked == 0x900D:
        return avr.LD_X_INC
    case masked == 0x900E:
        return avr.LD_X_DEC
    case masked == 0x900F:
        return avr.POP
    case masked == 0x9200 && !reducedCore:
        return avr.STS
    case masked == 0x920F:
        return avr.PUSH
    case masked == 0x9201:
        return avr.ST_Z_INC
    case masked == 0x9202:
        return avr.ST_Z_DEC
    case masked == 0x9204:
        return avr.XCH
    case masked == 0x9205:
        return avr.LAS
    case masked == 0x9206:
        return avr.LAC
    case masked == 0x9207:
        return avr.LAT
    case masked == 0x9209:
        return avr.ST_Y_INC
    case masked == 0x920A:
        return avr.ST_Y_DEC
    case masked == 0x920C:
        return avr.ST_X
    case masked == 0x920D:
        return avr.ST_X_INC
    case masked == 0x920E:
        return avr.ST_X_DEC
    case masked == 0x9400:
        return avr.COM
    case masked == 0x9401:
        return avr.NEG
    case masked == 0x9402:
        return avr.SWAP
    case masked == 0x9403:
        return avr.INC
    case masked == 0x9405:
        return avr.ASR
    case masked == 0x9406:
        return avr.LSR
    case masked == 0x9407:
        return avr.ROR
    case masked == 0x940A:
        return avr.DEC
    }
    
    if reducedCore {
        switch {
        case masked == 0x8000:
            return avr.LD_Z
        case masked == 0x8008:
            return avr.LD_Y
        case masked == 0x8200:
            return avr.ST_Z
        case masked == 0x8208:
            return avr.ST_Y
        }
    }
    
    if word&0xFF0F == 0x940B {
        return avr.DES
    }
    
    masked = word & 0xFF8F
    switch {
    case masked == 0x9408:
        return avr.BSET
    case masked == 0x9488:
        return avr.BCLR
    }
    
    switch {
    case word == 0x0000:
        return avr.NOP
    case word == 0x9409:
        return avr.IJMP
    case word == 0x9419:
        return avr.EIJMP
    case word == 0x9508:
        return avr.RET
    case word == 0x9509:
        return avr.ICALL
    case word == 0x9518:
        return avr.RETI
    case word == 0x9519:
        return avr.EICALL
    case word == 0x9588:
        return avr.SLEEP
    case word == 0x9598:
        return avr.BREAK
    case word == 0x95A8:
        return avr.WDR
    case word == 0x95C8:
        return avr.LPM_R0
    case word == 0x95D8:
        return avr.ELPM_R0
    case word == 0x95E8:
        return avr.SPM
    case word == 0x95F8:
        return avr.SPM_2
    }
    
    return -1
}

func DecodeLut9(word uint16, reducedCore bool) avr.Instruction {
    if word & 0xF000 == 0x9000 && !reducedCore {
        // delegate to lookup table
        return lut9[word & 0x0FFF]
    }
    
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
    case word&0xFC00 == 0x1800:
        return avr.SUB
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
    case word&0xF000 == 0x5000:
        return avr.SUBI
    case word&0xF000 == 0x6000:
        return avr.ORI
    case word&0xF000 == 0x7000:
        return avr.ANDI
    case word&0xFE0F == 0x8000 && reducedCore:
        return avr.LD_Z
    case word&0xD208 == 0x8000 && !reducedCore:
        return avr.LDD_Z
    case word&0xFE0F == 0x8008 && reducedCore:
        return avr.LD_Y
    case word&0xD208 == 0x8008 && !reducedCore:
        return avr.LDD_Y
    case word&0xFE0F == 0x8200 && reducedCore:
        return avr.ST_Z
    case word&0xD208 == 0x8200 && !reducedCore:
        return avr.STD_Z
    case word&0xFE0F == 0x8208 && reducedCore:
        return avr.ST_Y
    case word&0xD208 == 0x8208 && !reducedCore:
        return avr.STD_Y
    case word&0xF800 == 0xA000 && reducedCore:
        return avr.LDS_SHORT
    case word&0xF800 == 0xA800 && reducedCore:
        return avr.STS_SHORT
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

func DecodeLut9Tree(word uint16, reducedCore bool) avr.Instruction {
    if word & 0xF000 == 0x9000 && !reducedCore {
        // delegate to lookup table
        return lut9[word & 0x0FFF]
    }
    
    var masked uint16
    
    masked = word & 0xF000
    switch {
    case masked == 0x3000:
        return avr.CPI
    case masked == 0x4000:
        return avr.SBCI
    case masked == 0x5000:
        return avr.SUBI
    case masked == 0x6000:
        return avr.ORI
    case masked == 0x7000:
        return avr.ANDI
    case masked == 0xC000:
        return avr.RJMP
    case masked == 0xD000:
        return avr.RCALL
    case masked == 0xE000:
        return avr.LDI
    }
    
    if !reducedCore {
        masked = word & 0xD208
        switch {
        case masked == 0x8000:
            return avr.LDD_Z
        case masked == 0x8008:
            return avr.LDD_Y
        case masked == 0x8200:
            return avr.STD_Z
        case masked == 0x8208:
            return avr.STD_Y
        }
    }
    
    masked = word & 0xF800
    switch {
    case masked == 0xA000 && reducedCore:
        return avr.LDS_SHORT
    case masked == 0xA800 && reducedCore:
        return avr.STS_SHORT
    case masked == 0xB000:
        return avr.IN
    case masked == 0xB800:
        return avr.OUT
    }
    
    masked = word & 0xFC00
    switch {
    case masked == 0x0400:
        return avr.CPC
    case masked == 0x0800:
        return avr.SBC
    case masked == 0x0C00:
        return avr.ADD
    case masked == 0x1000:
        return avr.CPSE
    case masked == 0x1400:
        return avr.CP
    case masked == 0x1800:
        return avr.SUB
    case masked == 0x1C00:
        return avr.ADC
    case masked == 0x2000:
        return avr.AND
    case masked == 0x2400:
        return avr.EOR
    case masked == 0x2800:
        return avr.OR
    case masked == 0x2C00:
        return avr.MOV
    case masked == 0xF000:
        return avr.BRBS
    case masked == 0xF400:
        return avr.BRBC
    }
    
    masked = word & 0xFF00
    switch {
    case masked == 0x0100:
        return avr.MOVW
    case masked == 0x0200:
        return avr.MULS
    }
    
    masked = word & 0xFE08
    switch {
    case masked == 0xF800:
        return avr.BLD
    case masked == 0xFA00:
        return avr.BST
    case masked == 0xFC00:
        return avr.SBRC
    case masked == 0xFE00:
        return avr.SBRS
    }
    
    masked = word & 0xFF88
    switch {
    case masked == 0x0300:
        return avr.MULSU
    case masked == 0x0308:
        return avr.FMUL
    case masked == 0x0380:
        return avr.FMULS
    case masked == 0x0388:
        return avr.FMULSU
    }
    
    if reducedCore {
        masked = word & 0xFE0F
        switch {
        case masked == 0x8000:
            return avr.LD_Z
        case masked == 0x8008:
            return avr.LD_Y
        case masked == 0x8200:
            return avr.ST_Z
        case masked == 0x8208:
            return avr.ST_Y
        }
    }
    
    if word == 0x0000 {
        return avr.NOP
    }
    
    return -1
}

func DecodeLut(word uint16, reducedCore bool) avr.Instruction {
    if reducedCore {
        return lutRC[word]
    } else {
        return lutNonRC[word]
    }
}

var Decode = DecodeLut9Tree

// decoder lookup table for instructions whose high nibble is 9
// reducedCore assumed to be false
var lut9 [4096]avr.Instruction
func init() {
    max := uint16(len(lut9))
    for i := uint16(0); i < max; i++ {
        lut9[i] = DecodeFull(0x9000 | i, false)
    }
}

var lutRC [65536]avr.Instruction
var lutNonRC [65536]avr.Instruction
func init() {
    max := len(lutRC)
    for i := 0; i < max; i++ {
        lutRC[i] = DecodeFull(uint16(i), true)
        lutNonRC[i] = DecodeFull(uint16(i), false)
    }
}
