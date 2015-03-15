package timer

// Implementation of TCCRxA port
type tccra struct {
    t *Timer
}

func (p tccra) Read() uint8 {
    return p.t.controlA
}

func (p tccra) Write(x uint8) {
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
    if x & 0x80 != 0 {
        p.t.forceOutputCompare(0)
    }
    if x & 0x40 != 0 {
        p.t.forceOutputCompare(1)
    }
    
    p.t.controlB = x & 0x3F
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
    
    // Inhibit a compare match on the next clock
    p.t.inhibitCompareMatch = true
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
    // Bits in TIFRx are cleared by writing a one to them.
    p.t.interruptFlags &= ^x
}
