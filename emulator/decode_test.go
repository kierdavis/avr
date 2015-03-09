package emulator

import (
    "github.com/kierdavis/avr"
    "testing"
)

type decodeTest struct {
    word uint16
    inst avr.Instruction
}

var decodeTests = []decodeTest{
    decodeTest{0x1ef9, avr.ADC},
    decodeTest{0x1c23, avr.ADC},
    decodeTest{0x0efc, avr.ADD},
    decodeTest{0x0f6d, avr.ADD},
    decodeTest{0x96c7, avr.ADIW},
    decodeTest{0x96b4, avr.ADIW},
    decodeTest{0x2013, avr.AND},
    decodeTest{0x23ed, avr.AND},
    decodeTest{0x7dfc, avr.ANDI},
    decodeTest{0x70f6, avr.ANDI},
    decodeTest{0x9415, avr.ASR},
    decodeTest{0x95b5, avr.ASR},
    decodeTest{0x9498, avr.BCLR},
    decodeTest{0x94a8, avr.BCLR},
    decodeTest{0xf802, avr.BLD},
    decodeTest{0xf8f6, avr.BLD},
    decodeTest{0xf456, avr.BRBC},
    decodeTest{0xf5e3, avr.BRBC},
    decodeTest{0xf254, avr.BRBS},
    decodeTest{0xf179, avr.BRBS},
    decodeTest{0x9598, avr.BREAK},
    decodeTest{0x9418, avr.BSET},
    decodeTest{0x9468, avr.BSET},
    decodeTest{0xfb14, avr.BST},
    decodeTest{0xfbe7, avr.BST},
    decodeTest{0x94de, avr.CALL},
    decodeTest{0x95ef, avr.CALL},
    decodeTest{0x981c, avr.CBI},
    decodeTest{0x98b4, avr.CBI},
    decodeTest{0x9400, avr.COM},
    decodeTest{0x9530, avr.COM},
    decodeTest{0x170a, avr.CP},
    decodeTest{0x146f, avr.CP},
    decodeTest{0x060d, avr.CPC},
    decodeTest{0x0796, avr.CPC},
    decodeTest{0x3e57, avr.CPI},
    decodeTest{0x3ead, avr.CPI},
    decodeTest{0x1310, avr.CPSE},
    decodeTest{0x121e, avr.CPSE},
    decodeTest{0x944a, avr.DEC},
    decodeTest{0x94fa, avr.DEC},
    decodeTest{0x941b, avr.DES},
    decodeTest{0x94ab, avr.DES},
    decodeTest{0x9519, avr.EICALL},
    decodeTest{0x9419, avr.EIJMP},
    decodeTest{0x95D8, avr.ELPM_R0},
    decodeTest{0x9076, avr.ELPM},
    decodeTest{0x91a6, avr.ELPM},
    decodeTest{0x91f7, avr.ELPM_INC},
    decodeTest{0x9147, avr.ELPM_INC},
    decodeTest{0x242b, avr.EOR},
    decodeTest{0x24c1, avr.EOR},
    decodeTest{0x035d, avr.FMUL},
    decodeTest{0x0339, avr.FMUL},
    decodeTest{0x03f2, avr.FMULS},
    decodeTest{0x0380, avr.FMULS},
    decodeTest{0x03ce, avr.FMULSU},
    decodeTest{0x03ea, avr.FMULSU},
    decodeTest{0x9509, avr.ICALL},
    decodeTest{0x9409, avr.IJMP},
    decodeTest{0xb41d, avr.IN},
    decodeTest{0xb5c9, avr.IN},
    decodeTest{0x95f3, avr.INC},
    decodeTest{0x95d3, avr.INC},
    decodeTest{0x951d, avr.JMP},
    decodeTest{0x94ed, avr.JMP},
    decodeTest{0x9226, avr.LAC},
    decodeTest{0x93c6, avr.LAC},
    decodeTest{0x92d5, avr.LAS},
    decodeTest{0x9325, avr.LAS},
    decodeTest{0x9227, avr.LAT},
    decodeTest{0x93e7, avr.LAT},
    decodeTest{0x90cc, avr.LD_X},
    decodeTest{0x908c, avr.LD_X},
    decodeTest{0x904d, avr.LD_X_INC},
    decodeTest{0x91ad, avr.LD_X_INC},
    decodeTest{0x905e, avr.LD_X_DEC},
    decodeTest{0x90fe, avr.LD_X_DEC},
    decodeTest{0x80f8, avr.LDD_Y}, 
    decodeTest{0x8128, avr.LDD_Y},
    decodeTest{0x9119, avr.LD_Y_INC},
    decodeTest{0x9089, avr.LD_Y_INC},
    decodeTest{0x91da, avr.LD_Y_DEC},
    decodeTest{0x916a, avr.LD_Y_DEC},
    decodeTest{0xa938, avr.LDD_Y},
    decodeTest{0x80d9, avr.LDD_Y},
    decodeTest{0x8060, avr.LDD_Z},
    decodeTest{0x8130, avr.LDD_Z},
    decodeTest{0x9011, avr.LD_Z_INC},
    decodeTest{0x9031, avr.LD_Z_INC},
    decodeTest{0x90a2, avr.LD_Z_DEC},
    decodeTest{0x9102, avr.LD_Z_DEC},
    decodeTest{0xa9a7, avr.LDD_Z},
    decodeTest{0x8964, avr.LDD_Z},
    decodeTest{0xeb93, avr.LDI},
    decodeTest{0xeadd, avr.LDI},
    decodeTest{0x9170, avr.LDS},
    decodeTest{0x9190, avr.LDS},
}

func TestDecode(t *testing.T) {
    for _, test := range decodeTests {
        inst := Decode(test.word)
        if inst != test.inst {
            t.Errorf("Decode(0x%04x): expected '%s', got '%s'", test.word, test.inst, inst)
        }
    }
}

func BenchmarkDecode951d(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Decode(0x951d)
    }
}

func BenchmarkDecodeRandom(b *testing.B) {
    // implement an xorshift RNG for speed
    x := uint16(0xabcd)
    for i := 0; i < b.N; i++ {
        x ^= x << 13
        x ^= x >> 9
        x ^= x << 7
        Decode(x)
    }
}
