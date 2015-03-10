package spec

import (
    "fmt"
    "github.com/kierdavis/avr"
)

func tiny4_5_9_10(v int) *MCUSpec {
    ports := map[string]avr.PortRef{
        "PINB":   avr.PortRef{0, 0x00},
        "DDRB":   avr.PortRef{0, 0x01},
        "PORTB":  avr.PortRef{0, 0x02},
        "PUEB":   avr.PortRef{0, 0x03},
        "PORTCR": avr.PortRef{0, 0x0C},
        "PCMSK":  avr.PortRef{0, 0x10},
        "PCIFR":  avr.PortRef{0, 0x11},
        "PCICR":  avr.PortRef{0, 0x12},
        "EIMSK":  avr.PortRef{0, 0x13},
        "EIFR":   avr.PortRef{0, 0x14},
        "EICRA":  avr.PortRef{0, 0x15},
        "ACSR":   avr.PortRef{0, 0x1F},
        "ICR0L":  avr.PortRef{0, 0x22},
        "ICR0H":  avr.PortRef{0, 0x23},
        "OCR0BL": avr.PortRef{0, 0x24},
        "OCR0BH": avr.PortRef{0, 0x25},
        "OCR0AL": avr.PortRef{0, 0x26},
        "OCR0AH": avr.PortRef{0, 0x27},
        "TCNT0L": avr.PortRef{0, 0x28},
        "TCNT0H": avr.PortRef{0, 0x29},
        "TIFR0":  avr.PortRef{0, 0x2A},
        "TIMSK0": avr.PortRef{0, 0x2B},
        "TCCR0C": avr.PortRef{0, 0x2C},
        "TCCR0B": avr.PortRef{0, 0x2D},
        "TCCR0A": avr.PortRef{0, 0x2E},
        "GTCCR":  avr.PortRef{0, 0x2F},
        "WDTCSR": avr.PortRef{0, 0x31},
        "NVMCSR": avr.PortRef{0, 0x32},
        "NVMCMD": avr.PortRef{0, 0x33},
        "VLMCSR": avr.PortRef{0, 0x34},
        "PRR":    avr.PortRef{0, 0x35},
        "CLKPSR": avr.PortRef{0, 0x36},
        "CLKMSR": avr.PortRef{0, 0x37},
        "OSCCAL": avr.PortRef{0, 0x39},
        "SMCR":   avr.PortRef{0, 0x3A},
        "RSTFLR": avr.PortRef{0, 0x3B},
        "CCP":    avr.PortRef{0, 0x3C},
        "SPL":    avr.PortRef{0, 0x3D},
        "SPH":    avr.PortRef{0, 0x3E},
        "SREG":   avr.PortRef{0, 0x3F},
    }
    
    // ADC
    if v == 5 || v == 10 {
        ports["DIDR0"] = avr.PortRef{0, 0x17}
        ports["ADCL"] = avr.PortRef{0, 0x19}
        ports["ADMUX"] = avr.PortRef{0, 0x1B}
        ports["ADCSRA"] = avr.PortRef{0, 0x1C}
        ports["ADCSRB"] = avr.PortRef{0, 0x1D}
    }

    var logProgMemSize uint
    switch v {
    case 4, 5:
        logProgMemSize = 8 // 256 W (512 B)
    case 9, 10:
        logProgMemSize = 9 // 512 W (1024 B)
    }

    return linkRegions(&MCUSpec{
        Label:            fmt.Sprintf("ATtiny%d", v),
        Family:           ReducedCore,
        NumRegs:          32, // technically only 16, but the 16 that are implemented are r16-r31
        LogProgMemSize:   logProgMemSize,
        LogDataSpaceSize: 7, // data memory address width
        LogRAMSize:       5, // 32 B
        LogEEPROMSize:    0, // none
        IOBankSizes:      []uint{64},
        Regions: []RegionSpec{
            IORegionSpec{start: 0x0000, bankNum: 0},
            RAMRegionSpec{start: 0x0040},
        },
        Ports: ports,
        Available: [avr.NumInstructions]bool{
            /* ADC */       true,
            /* ADD */       true,
            /* ADIW */      false,
            /* AND */       true,
            /* ANDI */      true,
            /* ASR */       true,
            /* BCLR */      true,
            /* BLD */       true,
            /* BRBC */      true,
            /* BRBS */      true,
            /* BREAK */     true,
            /* BSET */      true,
            /* BST */       true,
            /* CALL */      false,
            /* CBI */       true,
            /* COM */       true,
            /* CP */        true,
            /* CPC */       true,
            /* CPI */       true,
            /* CPSE */      true,
            /* DEC */       true,
            /* DES */       false,
            /* EICALL */    false,
            /* EIJMP */     false,
            /* ELPM_R0 */   false,
            /* ELPM */      false,
            /* ELPM_INC */  false,
            /* EOR */       true,
            /* FMUL */      false,
            /* FMULS */     false,
            /* FMULSU */    false,
            /* ICALL */     true,
            /* IJMP */      true,
            /* IN */        true,
            /* INC */       true,
            /* JMP */       false,
            /* LAC */       false,
            /* LAS */       false,
            /* LAT */       false,
            /* LD_X */      true,
            /* LD_X_INC */  true,
            /* LD_X_DEC */  true,
            /* LD_Y */      true,
            /* LD_Y_INC */  true,
            /* LD_Y_DEC */  true,
            /* LDD_Y */     false,
            /* LD_Z */      true,
            /* LD_Z_INC */  true,
            /* LD_Z_DEC */  true,
            /* LDD_Z */     false,
            /* LDI */       true,
            /* LDS */       false,
            /* LDS_SHORT */ true,
            /* LPM_R0 */    false,
            /* LPM */       false,
            /* LPM_INC */   false,
            /* LSR */       true,
            /* MOV */       true,
            /* MOVW */      false,
            /* MUL */       false,
            /* MULS */      false,
            /* MULSU */     false,
            /* NEG */       true,
            /* NOP */       true,
            /* OR */        true,
            /* ORI */       true,
            /* OUT */       true,
            /* POP */       true,
            /* PUSH */      true,
            /* RCALL */     true,
            /* RET */       true,
            /* RETI */      true,
            /* RJMP */      true,
            /* ROR */       true,
            /* SBC */       true,
            /* SBCI */      true,
            /* SBI */       true,
            /* SBIC */      true,
            /* SBIS */      true,
            /* SBIW */      false,
            /* SBRC */      true,
            /* SBRS */      true,
            /* SLEEP */     true,
            /* SPM */       false,
            /* SPM_2 */     false,
            /* ST_X */      true,
            /* ST_X_INC */  true,
            /* ST_X_DEC */  true,
            /* ST_Y */      true,
            /* ST_Y_INC */  true,
            /* ST_Y_DEC */  true,
            /* STD_Y */     false,
            /* ST_Z */      true,
            /* ST_Z_INC */  true,
            /* ST_Z_DEC */  true,
            /* STD_Z */     false,
            /* STS */       false,
            /* STS_SHORT */ true,
            /* SUB */       true,
            /* SUBI */      true,
            /* SWAP */      true,
            /* WDR */       true,
            /* XCH */       false,
        },
    })
}

var ATtiny4 = tiny4_5_9_10(4)
var ATtiny5 = tiny4_5_9_10(5)
var ATtiny9 = tiny4_5_9_10(9)
var ATtiny10 = tiny4_5_9_10(10)
