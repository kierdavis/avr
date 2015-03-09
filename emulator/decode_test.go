package emulator

import (
    "github.com/kierdavis/avr"
    "testing"
)

type decodeTest struct {
    word uint16
    inst avr.Instruction
    reducedCore bool
}

var decodeTests = []decodeTest{
    decodeTest{word: 0x1ef9, inst: avr.ADC},
    decodeTest{word: 0x1c23, inst: avr.ADC},
    decodeTest{word: 0x0efc, inst: avr.ADD},
    decodeTest{word: 0x0f6d, inst: avr.ADD},
    decodeTest{word: 0x96c7, inst: avr.ADIW},
    decodeTest{word: 0x96b4, inst: avr.ADIW},
    decodeTest{word: 0x2013, inst: avr.AND},
    decodeTest{word: 0x23ed, inst: avr.AND},
    decodeTest{word: 0x7dfc, inst: avr.ANDI},
    decodeTest{word: 0x70f6, inst: avr.ANDI},
    decodeTest{word: 0x9415, inst: avr.ASR},
    decodeTest{word: 0x95b5, inst: avr.ASR},
    decodeTest{word: 0x9498, inst: avr.BCLR},
    decodeTest{word: 0x94a8, inst: avr.BCLR},
    decodeTest{word: 0xf802, inst: avr.BLD},
    decodeTest{word: 0xf8f6, inst: avr.BLD},
    decodeTest{word: 0xf456, inst: avr.BRBC},
    decodeTest{word: 0xf5e3, inst: avr.BRBC},
    decodeTest{word: 0xf254, inst: avr.BRBS},
    decodeTest{word: 0xf179, inst: avr.BRBS},
    decodeTest{word: 0x9598, inst: avr.BREAK},
    decodeTest{word: 0x9418, inst: avr.BSET},
    decodeTest{word: 0x9468, inst: avr.BSET},
    decodeTest{word: 0xfb14, inst: avr.BST},
    decodeTest{word: 0xfbe7, inst: avr.BST},
    decodeTest{word: 0x94de, inst: avr.CALL},
    decodeTest{word: 0x95ef, inst: avr.CALL},
    decodeTest{word: 0x981c, inst: avr.CBI},
    decodeTest{word: 0x98b4, inst: avr.CBI},
    decodeTest{word: 0x9400, inst: avr.COM},
    decodeTest{word: 0x9530, inst: avr.COM},
    decodeTest{word: 0x170a, inst: avr.CP},
    decodeTest{word: 0x146f, inst: avr.CP},
    decodeTest{word: 0x060d, inst: avr.CPC},
    decodeTest{word: 0x0796, inst: avr.CPC},
    decodeTest{word: 0x3e57, inst: avr.CPI},
    decodeTest{word: 0x3ead, inst: avr.CPI},
    decodeTest{word: 0x1310, inst: avr.CPSE},
    decodeTest{word: 0x121e, inst: avr.CPSE},
    decodeTest{word: 0x944a, inst: avr.DEC},
    decodeTest{word: 0x94fa, inst: avr.DEC},
    decodeTest{word: 0x941b, inst: avr.DES},
    decodeTest{word: 0x94ab, inst: avr.DES},
    decodeTest{word: 0x9519, inst: avr.EICALL},
    decodeTest{word: 0x9419, inst: avr.EIJMP},
    decodeTest{word: 0x95D8, inst: avr.ELPM_R0},
    decodeTest{word: 0x9076, inst: avr.ELPM},
    decodeTest{word: 0x91a6, inst: avr.ELPM},
    decodeTest{word: 0x91f7, inst: avr.ELPM_INC},
    decodeTest{word: 0x9147, inst: avr.ELPM_INC},
    decodeTest{word: 0x242b, inst: avr.EOR},
    decodeTest{word: 0x24c1, inst: avr.EOR},
    decodeTest{word: 0x035d, inst: avr.FMUL},
    decodeTest{word: 0x0339, inst: avr.FMUL},
    decodeTest{word: 0x03f2, inst: avr.FMULS},
    decodeTest{word: 0x0380, inst: avr.FMULS},
    decodeTest{word: 0x03ce, inst: avr.FMULSU},
    decodeTest{word: 0x03ea, inst: avr.FMULSU},
    decodeTest{word: 0x9509, inst: avr.ICALL},
    decodeTest{word: 0x9409, inst: avr.IJMP},
    decodeTest{word: 0xb41d, inst: avr.IN},
    decodeTest{word: 0xb5c9, inst: avr.IN},
    decodeTest{word: 0x95f3, inst: avr.INC},
    decodeTest{word: 0x95d3, inst: avr.INC},
    decodeTest{word: 0x951d, inst: avr.JMP},
    decodeTest{word: 0x94ed, inst: avr.JMP},
    decodeTest{word: 0x9226, inst: avr.LAC},
    decodeTest{word: 0x93c6, inst: avr.LAC},
    decodeTest{word: 0x92d5, inst: avr.LAS},
    decodeTest{word: 0x9325, inst: avr.LAS},
    decodeTest{word: 0x9227, inst: avr.LAT},
    decodeTest{word: 0x93e7, inst: avr.LAT},
    decodeTest{word: 0x90cc, inst: avr.LD_X},
    decodeTest{word: 0x908c, inst: avr.LD_X},
    decodeTest{word: 0x904d, inst: avr.LD_X_INC},
    decodeTest{word: 0x91ad, inst: avr.LD_X_INC},
    decodeTest{word: 0x905e, inst: avr.LD_X_DEC},
    decodeTest{word: 0x90fe, inst: avr.LD_X_DEC},
    decodeTest{word: 0x80f8, inst: avr.LDD_Y},
    decodeTest{word: 0x8128, inst: avr.LDD_Y},
    decodeTest{word: 0x9119, inst: avr.LD_Y_INC},
    decodeTest{word: 0x9089, inst: avr.LD_Y_INC},
    decodeTest{word: 0x91da, inst: avr.LD_Y_DEC},
    decodeTest{word: 0x916a, inst: avr.LD_Y_DEC},
    decodeTest{word: 0xa938, inst: avr.LDD_Y},
    decodeTest{word: 0x80d9, inst: avr.LDD_Y},
    decodeTest{word: 0x8060, inst: avr.LDD_Z},
    decodeTest{word: 0x8130, inst: avr.LDD_Z},
    decodeTest{word: 0x9011, inst: avr.LD_Z_INC},
    decodeTest{word: 0x9031, inst: avr.LD_Z_INC},
    decodeTest{word: 0x90a2, inst: avr.LD_Z_DEC},
    decodeTest{word: 0x9102, inst: avr.LD_Z_DEC},
    decodeTest{word: 0xa9a7, inst: avr.LDD_Z},
    decodeTest{word: 0x8964, inst: avr.LDD_Z},
    decodeTest{word: 0xeb93, inst: avr.LDI},
    decodeTest{word: 0xeadd, inst: avr.LDI},
    decodeTest{word: 0x9170, inst: avr.LDS},
    decodeTest{word: 0x9190, inst: avr.LDS},
    decodeTest{word: 0xa4fd, inst: avr.LDS_SHORT, reducedCore: true},
    decodeTest{word: 0xa688, inst: avr.LDS_SHORT, reducedCore: true},
    decodeTest{word: 0x95C8, inst: avr.LPM_R0},
    decodeTest{word: 0x9174, inst: avr.LPM},
    decodeTest{word: 0x9084, inst: avr.LPM},
    decodeTest{word: 0x90b5, inst: avr.LPM_INC},
    decodeTest{word: 0x91f5, inst: avr.LPM_INC},
    decodeTest{word: 0x95f6, inst: avr.LSR},
    decodeTest{word: 0x9416, inst: avr.LSR},
    decodeTest{word: 0x2c63, inst: avr.MOV},
    decodeTest{word: 0x2e77, inst: avr.MOV},
    decodeTest{word: 0x0105, inst: avr.MOVW},
    decodeTest{word: 0x01de, inst: avr.MOVW},
    decodeTest{word: 0x9cf2, inst: avr.MUL},
    decodeTest{word: 0x9f63, inst: avr.MUL},
    decodeTest{word: 0x0214, inst: avr.MULS},
    decodeTest{word: 0x022d, inst: avr.MULS},
    decodeTest{word: 0x0327, inst: avr.MULSU},
    decodeTest{word: 0x0341, inst: avr.MULSU},
    decodeTest{word: 0x94a1, inst: avr.NEG},
    decodeTest{word: 0x95c1, inst: avr.NEG},
    decodeTest{word: 0x0000, inst: avr.NOP},
    decodeTest{word: 0x2a07, inst: avr.OR},
    decodeTest{word: 0x2978, inst: avr.OR},
    decodeTest{word: 0x6f51, inst: avr.ORI},
    decodeTest{word: 0x6783, inst: avr.ORI},
    decodeTest{word: 0xb890, inst: avr.OUT},
    decodeTest{word: 0xbf07, inst: avr.OUT},
    decodeTest{word: 0x91cf, inst: avr.POP},
    decodeTest{word: 0x90ef, inst: avr.POP},
    decodeTest{word: 0x93af, inst: avr.PUSH},
    decodeTest{word: 0x928f, inst: avr.PUSH},
    decodeTest{word: 0xdeec, inst: avr.RCALL},
    decodeTest{word: 0xdf7e, inst: avr.RCALL},
    // reducedCore also true for STS_SHORT
    
    // TODO: add tests for opcodes that can resolve to two different instructions
    // depending on the status of the reducedCore flag.
}

func TestDecode(t *testing.T) {
    for _, test := range decodeTests {
        inst := Decode(test.word, test.reducedCore)
        if inst != test.inst {
            t.Errorf("Decode(0x%04x): expected '%s', got '%s'", test.word, test.inst, inst)
        }
    }
}

func BenchmarkDecode951d(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Decode(0x951d, false)
    }
}

func BenchmarkDecodeRandom(b *testing.B) {
    // implement an xorshift RNG for speed
    x := uint16(0xabcd)
    for i := 0; i < b.N; i++ {
        x ^= x << 13
        x ^= x >> 9
        x ^= x << 7
        Decode(x, false)
    }
}
