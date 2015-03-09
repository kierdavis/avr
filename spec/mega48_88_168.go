package spec

import (
	"fmt"
	"github.com/kierdavis/avr"
)

func mega48_88_168(v int) *MCUSpec {
	ports := map[string]avr.PortRef{
		"PINB":   avr.PortRef{0, 0x03},
		"DDRB":   avr.PortRef{0, 0x04},
		"PORTB":  avr.PortRef{0, 0x05},
		"PINC":   avr.PortRef{0, 0x06},
		"DDRC":   avr.PortRef{0, 0x07},
		"PORTC":  avr.PortRef{0, 0x08},
		"PIND":   avr.PortRef{0, 0x09},
		"DDRD":   avr.PortRef{0, 0x0A},
		"PORTD":  avr.PortRef{0, 0x0B},
		"TIFR0":  avr.PortRef{0, 0x15},
		"TIFR1":  avr.PortRef{0, 0x16},
		"TIFR2":  avr.PortRef{0, 0x17},
		"PCIFR":  avr.PortRef{0, 0x1B},
		"EIFR":   avr.PortRef{0, 0x1C},
		"EIMSK":  avr.PortRef{0, 0x1D},
		"GPIOR0": avr.PortRef{0, 0x1E},
		"EECR":   avr.PortRef{0, 0x1F},
		"EEDR":   avr.PortRef{0, 0x20},
		"EEARL":  avr.PortRef{0, 0x21},
		"EEARH":  avr.PortRef{0, 0x22},
		"TCCR0A": avr.PortRef{0, 0x24},
		"TCCR0B": avr.PortRef{0, 0x25},
		"TCNT0":  avr.PortRef{0, 0x26},
		"OCR0A":  avr.PortRef{0, 0x27},
		"OCR0B":  avr.PortRef{0, 0x28},
		"GPIOR1": avr.PortRef{0, 0x2A},
		"GPIOR2": avr.PortRef{0, 0x2B},
		"SPCR":   avr.PortRef{0, 0x2C},
		"SPSR":   avr.PortRef{0, 0x2D},
		"SPDR":   avr.PortRef{0, 0x2E},
		"ACSR":   avr.PortRef{0, 0x30},
		"SMCR":   avr.PortRef{0, 0x33},
		"MCUSR":  avr.PortRef{0, 0x34},
		"MCUCR":  avr.PortRef{0, 0x35},
		"SPMCSR": avr.PortRef{0, 0x37},
		"SPL":    avr.PortRef{0, 0x3D},
		"SPH":    avr.PortRef{0, 0x3E},
		"SREG":   avr.PortRef{0, 0x3F},
		"WDTCSR": avr.PortRef{1, 0x00},
		"CLKPR":  avr.PortRef{1, 0x01},
		"PRR":    avr.PortRef{1, 0x04},
		"OSCCAL": avr.PortRef{1, 0x06},
		"PCICR":  avr.PortRef{1, 0x08},
		"EICRA":  avr.PortRef{1, 0x09},
		"PCMSK0": avr.PortRef{1, 0x0B},
		"PCMSK1": avr.PortRef{1, 0x0C},
		"PCMSK2": avr.PortRef{1, 0x0D},
		"TIMSK0": avr.PortRef{1, 0x0E},
		"TIMSK1": avr.PortRef{1, 0x0F},
		"TIMSK2": avr.PortRef{1, 0x10},
		"ADCL":   avr.PortRef{1, 0x18},
		"ADCH":   avr.PortRef{1, 0x19},
		"ADCSRA": avr.PortRef{1, 0x1A},
		"ADCSRB": avr.PortRef{1, 0x1B},
		"ADMUX":  avr.PortRef{1, 0x1C},
		"DIDR0":  avr.PortRef{1, 0x1E},
		"DIDR1":  avr.PortRef{1, 0x1F},
		"TCCR1A": avr.PortRef{1, 0x20},
		"TCCR1B": avr.PortRef{1, 0x21},
		"TCCR1C": avr.PortRef{1, 0x22},
		"TCNT1L": avr.PortRef{1, 0x24},
		"TCNT1H": avr.PortRef{1, 0x25},
		"ICR1L":  avr.PortRef{1, 0x26},
		"ICR1H":  avr.PortRef{1, 0x27},
		"OCR1AL": avr.PortRef{1, 0x28},
		"OCR1AH": avr.PortRef{1, 0x29},
		"OCR1BL": avr.PortRef{1, 0x2A},
		"OCR1BH": avr.PortRef{1, 0x2B},
		"TCCR2A": avr.PortRef{1, 0x50},
		"TCCR2B": avr.PortRef{1, 0x51},
		"TCNT2":  avr.PortRef{1, 0x52},
		"OCR2A":  avr.PortRef{1, 0x53},
		"OCR2B":  avr.PortRef{1, 0x54},
		"ASSR":   avr.PortRef{1, 0x56},
		"TWBR":   avr.PortRef{1, 0x58},
		"TWSR":   avr.PortRef{1, 0x59},
		"TWAR":   avr.PortRef{1, 0x5A},
		"TWDR":   avr.PortRef{1, 0x5B},
		"TWCR":   avr.PortRef{1, 0x5C},
		"TWAMR":  avr.PortRef{1, 0x5D},
		"UCSR0A": avr.PortRef{1, 0x60},
		"UCSR0B": avr.PortRef{1, 0x61},
		"UCSR0C": avr.PortRef{1, 0x62},
		"UBRR0L": avr.PortRef{1, 0x64},
		"UBRR0H": avr.PortRef{1, 0x65},
		"UDR0":   avr.PortRef{1, 0x66},
	}

	// GTCCR port only present on ATmega88/168
	if v != 48 {
		ports["GTCCR"] = avr.PortRef{0, 0x23}
	}

	var logProgMemSize, logDataSpaceSize, logRAMSize, logEEPROMSize uint
	switch v {
	case 48:
		logProgMemSize = 11 // 2 kW (4 kB)
		logDataSpaceSize = 10
		logRAMSize = 9      // 512 B
		logEEPROMSize = 8   // 256 B
	case 88:
		logProgMemSize = 12 // 4 kW (8 kB)
		logDataSpaceSize = 11
		logRAMSize = 10     // 1 kB
		logEEPROMSize = 9   // 512 B
	case 168:
		logProgMemSize = 13 // 8 kW (16 kB)
		logDataSpaceSize = 11
		logRAMSize = 10     // 1 kB
		logEEPROMSize = 9   // 512 B
	}

	return linkRegions(&MCUSpec{
		Label:            fmt.Sprintf("ATmega%d", v),
		Family:           EnhancedCore128K,
		NumRegs:          32,
		LogProgMemSize:   logProgMemSize,
		LogDataSpaceSize: logDataSpaceSize, // data memory address width
		LogRAMSize:       logRAMSize,
		LogEEPROMSize:    logEEPROMSize,
		IOBankSizes:      []uint{64, 160},
		Regions: []RegionSpec{
			RegsRegionSpec{start: 0x0000},
			IORegionSpec{start: 0x0020, bankNum: 0},
			IORegionSpec{start: 0x0060, bankNum: 1},
			RAMRegionSpec{start: 0x0100},
		},
		Ports: ports,
	})
}

var ATmega48 = mega48_88_168(48)
var ATmega88 = mega48_88_168(88)
var ATmega168 = mega48_88_168(168)
