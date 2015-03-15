package main

import (
    "github.com/kierdavis/avr"
)

type RCMode uint8

const (
	RC RCMode = iota
	NotRC
	Either
)

type InstDef struct {
	Inst avr.Instruction
	Mask uint16
	Match uint16
	RCMode RCMode
}

var InstDefs = []InstDef{
	InstDef{avr.ADC,       0xFC00, 0x1C00, Either},
	InstDef{avr.ADD,       0xFC00, 0x0C00, Either},
	InstDef{avr.ADIW,      0xFF00, 0x9600, Either},
	InstDef{avr.AND,       0xFC00, 0x2000, Either},
	InstDef{avr.ANDI,      0xF000, 0x7000, Either},
	InstDef{avr.ASR,       0xFE0F, 0x9405, Either},
	InstDef{avr.BCLR,      0xFF8F, 0x9488, Either},
	InstDef{avr.BLD,       0xFE08, 0xF800, Either},
	InstDef{avr.BRBC,      0xFC00, 0xF400, Either},
	InstDef{avr.BRBS,      0xFC00, 0xF000, Either},
	InstDef{avr.BREAK,     0xFFFF, 0x9598, Either},
	InstDef{avr.BSET,      0xFF8F, 0x9408, Either},
	InstDef{avr.BST,       0xFE08, 0xFA00, Either},
	InstDef{avr.CALL,      0xFE0E, 0x940E, Either},
	InstDef{avr.CBI,       0xFF00, 0x9800, Either},
	InstDef{avr.COM,       0xFE0F, 0x9400, Either},
	InstDef{avr.CP,        0xFC00, 0x1400, Either},
	InstDef{avr.CPC,       0xFC00, 0x0400, Either},
	InstDef{avr.CPI,       0xF000, 0x3000, Either},
	InstDef{avr.CPSE,      0xFC00, 0x1000, Either},
	InstDef{avr.DEC,       0xFE0F, 0x940A, Either},
	InstDef{avr.DES,       0xFF0F, 0x940B, Either},
	InstDef{avr.EICALL,    0xFFFF, 0x9519, Either},
	InstDef{avr.EIJMP,     0xFFFF, 0x9419, Either},
	InstDef{avr.ELPM,      0xFE0F, 0x9006, Either},
	InstDef{avr.ELPM_INC,  0xFE0F, 0x9007, Either},
	InstDef{avr.ELPM_R0,   0xFFFF, 0x95D8, Either},
	InstDef{avr.EOR,       0xFC00, 0x2400, Either},
	InstDef{avr.FMUL,      0xFF88, 0x0308, Either},
	InstDef{avr.FMULS,     0xFF88, 0x0380, Either},
	InstDef{avr.FMULSU,    0xFF88, 0x0388, Either},
	InstDef{avr.ICALL,     0xFFFF, 0x9509, Either},
	InstDef{avr.IJMP,      0xFFFF, 0x9409, Either},
	InstDef{avr.IN,        0xF800, 0xB000, Either},
	InstDef{avr.INC,       0xFE0F, 0x9403, Either},
	InstDef{avr.JMP,       0xFE0E, 0x940C, Either},
	InstDef{avr.LAC,       0xFE0F, 0x9206, Either},
	InstDef{avr.LAS,       0xFE0F, 0x9205, Either},
	InstDef{avr.LAT,       0xFE0F, 0x9207, Either},
	InstDef{avr.LD_X,      0xFE0F, 0x900C, Either},
	InstDef{avr.LD_X_DEC,  0xFE0F, 0x900E, Either},
	InstDef{avr.LD_X_INC,  0xFE0F, 0x900D, Either},
	InstDef{avr.LD_Y,      0xFE0F, 0x8008, RC},
	InstDef{avr.LD_Y_DEC,  0xFE0F, 0x900A, Either},
	InstDef{avr.LD_Y_INC,  0xFE0F, 0x9009, Either},
	InstDef{avr.LD_Z,      0xFE0F, 0x8000, RC},
	InstDef{avr.LD_Z_DEC,  0xFE0F, 0x9002, Either},
	InstDef{avr.LD_Z_INC,  0xFE0F, 0x9001, Either},
	InstDef{avr.LDD_Y,     0xD208, 0x8008, NotRC},
	InstDef{avr.LDD_Z,     0xD208, 0x8000, NotRC},
	InstDef{avr.LDI,       0xF000, 0xE000, Either},
	InstDef{avr.LDS,       0xFE0F, 0x9000, NotRC},
	InstDef{avr.LDS_SHORT, 0xF800, 0xA000, RC},
	InstDef{avr.LPM,       0xFE0F, 0x9004, Either},
	InstDef{avr.LPM_INC,   0xFE0F, 0x9005, Either},
	InstDef{avr.LPM_R0,    0xFFFF, 0x95C8, Either},
	InstDef{avr.LSR,       0xFE0F, 0x9406, Either},
	InstDef{avr.MOV,       0xFC00, 0x2C00, Either},
	InstDef{avr.MOVW,      0xFF00, 0x0100, Either},
	InstDef{avr.MUL,       0xFC00, 0x9C00, Either},
	InstDef{avr.MULS,      0xFF00, 0x0200, Either},
	InstDef{avr.MULSU,     0xFF88, 0x0300, Either},
	InstDef{avr.NEG,       0xFE0F, 0x9401, Either},
	InstDef{avr.NOP,       0xFFFF, 0x0000, Either},
	InstDef{avr.OR,        0xFC00, 0x2800, Either},
	InstDef{avr.ORI,       0xF000, 0x6000, Either},
	InstDef{avr.OUT,       0xF800, 0xB800, Either},
	InstDef{avr.POP,       0xFE0F, 0x900F, Either},
	InstDef{avr.PUSH,      0xFE0F, 0x920F, Either},
	InstDef{avr.RCALL,     0xF000, 0xD000, Either},
	InstDef{avr.RET,       0xFFFF, 0x9508, Either},
	InstDef{avr.RETI,      0xFFFF, 0x9518, Either},
	InstDef{avr.RJMP,      0xF000, 0xC000, Either},
	InstDef{avr.ROR,       0xFE0F, 0x9407, Either},
	InstDef{avr.SBC,       0xFC00, 0x0800, Either},
	InstDef{avr.SBCI,      0xF000, 0x4000, Either},
	InstDef{avr.SBI,       0xFF00, 0x9A00, Either},
	InstDef{avr.SBIC,      0xFF00, 0x9900, Either},
	InstDef{avr.SBIS,      0xFF00, 0x9B00, Either},
	InstDef{avr.SBIW,      0xFF00, 0x9700, Either},
	InstDef{avr.SBRC,      0xFE08, 0xFC00, Either},
	InstDef{avr.SBRS,      0xFE08, 0xFE00, Either},
	InstDef{avr.SLEEP,     0xFFFF, 0x9588, Either},
	InstDef{avr.SPM,       0xFFFF, 0x95E8, Either},
	InstDef{avr.SPM_2,     0xFFFF, 0x95F8, Either},
	InstDef{avr.ST_X,      0xFE0F, 0x920C, Either},
	InstDef{avr.ST_X_DEC,  0xFE0F, 0x920E, Either},
	InstDef{avr.ST_X_INC,  0xFE0F, 0x920D, Either},
	InstDef{avr.ST_Y,      0xFE0F, 0x8208, RC},
	InstDef{avr.ST_Y_DEC,  0xFE0F, 0x920A, Either},
	InstDef{avr.ST_Y_INC,  0xFE0F, 0x9209, Either},
	InstDef{avr.ST_Z,      0xFE0F, 0x8200, RC},
	InstDef{avr.ST_Z_DEC,  0xFE0F, 0x9202, Either},
	InstDef{avr.ST_Z_INC,  0xFE0F, 0x9201, Either},
	InstDef{avr.STD_Y,     0xD208, 0x8208, NotRC},
	InstDef{avr.STD_Z,     0xD208, 0x8200, NotRC},
	InstDef{avr.STS,       0xFE0F, 0x9200, NotRC},
	InstDef{avr.STS_SHORT, 0xF800, 0xA800, RC},
	InstDef{avr.SUB,       0xFC00, 0x1800, Either},
	InstDef{avr.SUBI,      0xF000, 0x5000, Either},
	InstDef{avr.SWAP,      0xFE0F, 0x9402, Either},
	InstDef{avr.WDR,       0xFFFF, 0x95A8, Either},
	InstDef{avr.XCH,       0xFE0F, 0x9204, Either},
}

func Decode(word uint16, reducedCore bool) (inst avr.Instruction) {
	var r RCMode
	if reducedCore {
		r = RC
	} else {
		r = NotRC
	}
	
	for _, instDef := range InstDefs {
		if word & instDef.Mask == instDef.Match && (instDef.RCMode == r || instDef.RCMode == Either) {
			return instDef.Inst
		}
	}
	
	return -1
}
