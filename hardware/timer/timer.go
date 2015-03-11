// Package gpio implements a timer/counter unit.
// Tested compatibility:
//   ATmega48/88/168
// Untested compatability:
//   ATtiny4/5/9/10
package timer

import (
    "fmt"
    "github.com/kierdavis/avr/clock"
    "github.com/kierdavis/avr/emulator"
    "github.com/kierdavis/avr/hardware/gpio"
    "log"
)

// TODO: for a OSCx output, require corresponding DDR bit to be set to output

// TODO: writing to TCNT prevents a compare match on the next clock

// TODO: thread-safety!!

type Timer struct {
    digit uint
    controlA uint8
    controlB uint8
    count uint8
    compareValA uint8
    compareValB uint8
    interruptMask uint8
    interruptFlags uint8
    downwards bool // count direction
    ocPinStates [2]bool
    ocPinCallbacks [2]func(bool)
    logging bool
}

func New(digit uint) (t *Timer) {
    return &Timer{
        digit: digit,
    }
}

func (t *Timer) SetLogging(enabled bool) {
    t.logging = enabled
}

func (t *Timer) AddTo(em *emulator.Emulator) {
    em.RegisterPortByName(fmt.Sprintf("TCCR%dA", t.digit), tccra{t})
    em.RegisterPortByName(fmt.Sprintf("TCCR%dB", t.digit), tccrb{t})
    em.RegisterPortByName(fmt.Sprintf("TCNT%d", t.digit), tcnt{t})
    em.RegisterPortByName(fmt.Sprintf("OCR%dA", t.digit), ocra{t})
    em.RegisterPortByName(fmt.Sprintf("OCR%dB", t.digit), ocrb{t})
    em.RegisterPortByName(fmt.Sprintf("TIMSK%d", t.digit), timsk{t})
    em.RegisterPortByName(fmt.Sprintf("TIFR%d", t.digit), tifr{t})
}

// Connect an output-compare pin to a GPIO port by calling the GPIO's
// OverrideOutput method.
func (t *Timer) OverrideOCPin(ocPinNum uint, gpioPinNum uint, g *gpio.GPIO) {
    t.ocPinCallbacks[ocPinNum] = g.OverrideOutput(gpioPinNum)
}

// Run the timer off the specified clock.
func (t *Timer) Run(clk clock.Clock) {
    // TODO: implement clock source selector again
    
    for {
        now := clk.Now()
        t.Tick()
        clk.Await(now + 1)
    }
}

// Note: in PWM modes, OCRA/OCRB do not exhibit a newly written value until the count overflows

// TODO: in PC-PWM mode, an OCRx bit may transition without a compare match for two reasons (see datasheet page 98)

// Tick the timer.
func (t *Timer) Tick() {
    // Get the waveform generation mode bits
    wgmA := t.controlA & 0x03 // bits 1 and 0
    wgmB := (t.controlB & 0x08) >> 1 // bit 3 (shifted to bit 2)
    wgm := wgmB | wgmA
    
    // Prepare to tick counter
    switch wgm {
    case 0: // Normal
        t.tickNormalMode()
    case 1: // Phase-correct PWM (TOP = 0xFF)
        t.tickPCPWMMode(0xFF)
    case 2: // Clear timer on compare
        t.tickCTCMode()
    case 3: // Fast PWM (TOP = 0xFF)
        t.tickFastPWMMode(0xFF)
    case 5: // Phase-correct PWM (TOP = OCRA)
        t.tickPCPWMMode(t.compareValA)
    case 7: // Fast PWM (TOP = OCRA)
        t.tickPCPWMMode(t.compareValA)
    }
    
    // Actually tick the counter
    if t.downwards {
        t.count--
    } else {
        t.count++
    }
    
    /*
    if t.logging {
        log.Printf("[avr/hardware/timer:(*Timer).tick] ticked, count is now $%02X", t.count)
    }
    */
}

// Tick the timer in Normal mode.
func (t *Timer) tickNormalMode() {
    t.downwards = false
    if t.count == 0xFF { // Overflow
        t.interruptFlags |= 0x01 // set TOVx bit
    }
    
    t.checkOCPinNormalMode(0, t.compareValA)
    t.checkOCPinNormalMode(1, t.compareValB)
}

// Check for output-compare in normal mode.
func (t *Timer) checkOCPinNormalMode(ocPinNum uint, compareVal uint8) {
    if t.count == compareVal {
        // Get COMxy bits
        shiftAmt := 6 - 2*ocPinNum // 0 => 6, 1 => 4
        com := (t.controlA >> shiftAmt) & 0x03
        switch com {
        case 0: // OCy disabled
            // do nothing
        case 1: // toggle OCy
            t.toggleOCPin(ocPinNum)
        case 2: // clear OCy
            t.clearOCPin(ocPinNum)
        case 3: // set OCy
            t.setOCPin(ocPinNum)
        }
    }
}

// Tick the timer in phase-corrected PWM mode.
func (t *Timer) tickPCPWMMode(top uint8) {
    if t.downwards {
        if t.count - 1 == 0x00 { // Reached BOTTOM
            t.downwards = false // Begin counting upwards
        }
    } else {
        if t.count + 1 == top { // Reached TOP
            t.downwards = true // Begin counting downwards
        }
    }
    
    t.checkOCPinPCPWMMode(0, t.compareValA)
    t.checkOCPinPCPWMMode(1, t.compareValB)
}

// TODO: there is a special case in the OC pin checking for both PWM modes when OCRy == TOP and COMxy1 is set (see datasheet page 101)

// Check for output-compare in phase-corrected PWM mode.
func (t *Timer) checkOCPinPCPWMMode(ocPinNum uint, compareVal uint8) {
    if t.count == compareVal {
        // Get COMxy bits
        shiftAmt := 6 - 2*ocPinNum // 0 => 6, 1 => 4
        com := (t.controlA >> shiftAmt) & 0x03
        switch com {
        case 0: // OCy disabled
            // do nothing
        case 1: // Toggle OCy (only on OC pin 0 with WGM2 bit set)
            if ocPinNum == 0 && (t.controlB & 0x80) != 0 {
                t.toggleOCPin(ocPinNum)
            }
        case 2: // Clear OCy if counting upwards or set OCy if counting downwards
            if t.downwards {
                t.setOCPin(ocPinNum)
            } else {
                t.clearOCPin(ocPinNum)
            }
        case 3: // Set OCy if counting upwards or clear OCy if counting downwards
            if t.downwards {
                t.clearOCPin(ocPinNum)
            } else {
                t.setOCPin(ocPinNum)
            }
        }
    }
}

// Tick the timer in clear-timer-on-compare mode.
func (t *Timer) tickCTCMode() {
    t.downwards = false
    if t.count == t.compareValA {
        // this tick should set counter to 0
        t.count = 0xFF
    }
}

// Tick the timer in fast PWM mode.
func (t *Timer) tickFastPWMMode(top uint8) {
    t.downwards = false
    if t.count == top {
        // this tick should set counter to 0
        t.count = 0xFF
    }
    
    t.checkOCPinFastPWMMode(0, t.compareValA)
    t.checkOCPinFastPWMMode(1, t.compareValB)
}

// Check for output-compare in fast PWM mode.
func (t *Timer) checkOCPinFastPWMMode(ocPinNum uint, compareVal uint8) {
    // BOTTOM
    if t.count == 0x00 {
        // Get COMxy bits
        shiftAmt := 6 - 2*ocPinNum // 0 => 6, 1 => 4
        com := (t.controlA >> shiftAmt) & 0x03
        switch com {
        case 0: // OCy disabled
            // do nothing
        case 1: // Toggle OCy on compare match (only on OC pin 0 with WGM2 bit set)
            // do nothing
        case 2: // Clear OCy on compare match, set OCy at BOTTOM
            t.setOCPin(ocPinNum)
        case 3: // Set OCy on compare match, clear OCy at BOTTOM
            t.clearOCPin(ocPinNum)
        }
    }
    
    // Compare match
    if t.count == compareVal {
        // Get COMxy bits
        shiftAmt := 6 - 2*ocPinNum // 0 => 6, 1 => 4
        com := (t.controlA >> shiftAmt) & 0x03
        switch com {
        case 0: // OCy disabled
            // do nothing
        case 1: // Toggle OCy on compare match (only on OC pin 0 with WGM2 bit set)
            if ocPinNum == 0 && (t.controlB & 0x80) != 0 {
                t.toggleOCPin(ocPinNum)
            }
        case 2: // Clear OCy on compare match, set OCy at BOTTOM
            t.clearOCPin(ocPinNum)
        case 3: // Set OCy on compare match, clear OCy at BOTTOM
            t.setOCPin(ocPinNum)
        }
    }
}

// Toggle an output-compare pin.
func (t *Timer) toggleOCPin(ocPinNum uint) {
    t.ocPinStates[ocPinNum] = !t.ocPinStates[ocPinNum]
    t.updateOCPin(ocPinNum)
}

// Set an output-compare pin to low.
func (t *Timer) clearOCPin(ocPinNum uint) {
    if t.ocPinStates[ocPinNum] {
        t.ocPinStates[ocPinNum] = false
        t.updateOCPin(ocPinNum)
    }
}

// Set an output-compare pin to high.
func (t *Timer) setOCPin(ocPinNum uint) {
    if !t.ocPinStates[ocPinNum] {
        t.ocPinStates[ocPinNum] = true
        t.updateOCPin(ocPinNum)
    }
}

// Push the new status of an output-compare pin to the GPIO layer.
func (t *Timer) updateOCPin(ocPinNum uint) {
    callback := t.ocPinCallbacks[ocPinNum]
    if callback != nil {
        callback(t.ocPinStates[ocPinNum])
    }
}

// Implementation of TCCRxA port
type tccra struct {
    t *Timer
}

func (p tccra) Read() uint8 {
    return p.t.controlA
}

func (p tccra) Write(x uint8) {
    if p.t.logging {
        log.Printf("[avr/hardware/timer:tccrb.Write] $%02X (%08b) written to TCCRxA", x, x)
    }
    p.t.controlA = x
}

// Implementation of TCCRxB port
type tccrb struct {
    t *Timer
}

func (p tccrb) Read() uint8 {
    return p.t.controlB
}

func (p tccrb) Write(x uint8) {
    if p.t.logging {
        log.Printf("[avr/hardware/timer:tccrb.Write] $%02X (%08b) written to TCCRxB", x, x)
    }
    p.t.controlB = x
}

// Implementation of TCNTx port
type tcnt struct {
    t *Timer
}

func (p tcnt) Read() uint8 {
    return p.t.count
}

func (p tcnt) Write(x uint8) {
    p.t.count = x
}

// Implementation of OCRxA port
type ocra struct {
    t *Timer
}

func (p ocra) Read() uint8 {
    return p.t.compareValA
}

func (p ocra) Write(x uint8) {
    p.t.compareValA = x
}

// Implementation of OCRxb port
type ocrb struct {
    t *Timer
}

func (p ocrb) Read() uint8 {
    return p.t.compareValB
}

func (p ocrb) Write(x uint8) {
    p.t.compareValB = x
}

// Implementation of TIMSKx port
type timsk struct {
    t *Timer
}

func (p timsk) Read() uint8 {
    return p.t.interruptMask
}

func (p timsk) Write(x uint8) {
    p.t.interruptMask = x
}

// Implementation of TIFRx port
type tifr struct {
    t *Timer
}

func (p tifr) Read() uint8 {
    return p.t.interruptFlags
}

func (p tifr) Write(x uint8) {
    p.t.interruptFlags = x
}
