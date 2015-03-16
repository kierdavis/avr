package avr

import (
    "testing"
)

func TestInstructionString(t *testing.T) {
    if ADC.String() != "ADC" {
        t.Error("ADC.String(): expected \"ADC\", got %q", ADC.String())
    }

    if LD_Z_DEC.String() != "LD_Z_DEC" {
        t.Error("LD_Z_DEC.String(): expected \"LD_Z_DEC\", got %q", LD_Z_DEC.String())
    }

    if SPM_2.String() != "SPM_2" {
        t.Error("SPM_2.String(): expected \"SPM_2\", got %q", SPM_2.String())
    }

    if XCH.String() != "XCH" {
        t.Error("XCH.String(): expected \"XCH\", got %q", XCH.String())
    }
}
