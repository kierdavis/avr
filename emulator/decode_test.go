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
}

func TestDecode(t *testing.T) {
    for _, test := range decodeTests {
        inst := Decode(test.word)
        if inst != test.inst {
            t.Errorf("Decode(0x%04x): expected '%s', got '%s'", test.word, test.inst, inst)
        }
    }
}
