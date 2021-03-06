package emulator

import (
    "github.com/kierdavis/avr"
    "github.com/kierdavis/avr/spec"
    "log"
)

// An Emulator encapsulates the state of a processor.
type Emulator struct {
    Spec        *spec.MCUSpec
    regions     []Region
    ports       [][]Port
    prog        []uint16
    ram         []uint8
    pc          uint32
    pcmask      uint32
    sp          uint16
    rampx       uint8
    rampy       uint8
    rampz       uint8
    rampd       uint8
    eind        uint8
    regs        [32]uint8
    flags       [8]uint8
    logging     bool
    excessTicks uint
}

// NewEmulator creates and returns an initialised Emulator for the given MCUSpec.
func NewEmulator(mcuSpec *spec.MCUSpec) (em *Emulator) {
    em = &Emulator{
        Spec:    mcuSpec,
        regions: make([]Region, len(mcuSpec.Regions)),
        ports:   make([][]Port, len(mcuSpec.IOBankSizes)),
        prog:    make([]uint16, 1<<mcuSpec.LogProgMemSize),
        ram:     make([]uint8, 1<<mcuSpec.LogRAMSize),
        pc:      0,
        pcmask:  (1 << mcuSpec.LogProgMemSize) - 1,
    }

    for i := range em.ports {
        em.ports[i] = make([]Port, mcuSpec.IOBankSizes[i])
    }

    // register standard ports
    em.RegisterPortByName("SPL", SplPort{em})
    em.RegisterPortByName("SPH", SphPort{em})
    em.RegisterPortByName("SREG", SregPort{em})

    // create memory regions
    for i, regionSpec_ := range mcuSpec.Regions {
        switch regionSpec := regionSpec_.(type) {
        case spec.RegsRegionSpec:
            em.regions[i] = RegsRegion{em, regionSpec}
        case spec.IORegionSpec:
            em.regions[i] = IORegion{em, regionSpec}
        case spec.RAMRegionSpec:
            em.regions[i] = RAMRegion{em, regionSpec}
        }
    }

    return em
}

func (em *Emulator) SetLogging(enabled bool) {
    em.logging = enabled
}

func (em *Emulator) RegisterPort(pref avr.PortRef, port Port) {
    em.ports[pref.BankNum][pref.Index] = port
}

func (em *Emulator) RegisterPortByName(name string, port Port) (ok bool) {
    pref, ok := em.Spec.Ports[name]
    if !ok {
        return false
    }
    em.RegisterPort(pref, port)
    return true
}

func (em *Emulator) UnregisterPort(pref avr.PortRef) {
    em.ports[pref.BankNum][pref.Index] = nil
}

func (em *Emulator) UnregisterPortByName(name string) (ok bool) {
    pref, ok := em.Spec.Ports[name]
    if !ok {
        return false
    }
    em.UnregisterPort(pref)
    return true
}

func (em *Emulator) InterruptsEnabled() bool {
    return em.flags[avr.FlagI] != 0
}

func (em *Emulator) Interrupt(num uint) {
    if em.InterruptsEnabled() {
        em.flags[avr.FlagI] = 0
        em.pushPC()
        em.pc = uint32(num * em.Spec.InterruptVectorSize)
    }
}

func (em *Emulator) InterruptByName(name string) (ok bool) {
    num, ok := em.Spec.Interrupts[name]
    if !ok {
        return false
    }
    em.Interrupt(num)
    return true
}

func (em *Emulator) Run(ticks uint) {
    // subtract ticks that were executed on the last call to Run
    ticksExecuted := em.excessTicks

    var decodeFunc func(uint16) avr.Instruction
    if em.Spec.Family == spec.ReducedCore {
        decodeFunc = DecodeRC
    } else {
        decodeFunc = DecodeNonRC
    }

    for ticksExecuted < ticks {
        word := em.fetchProgWord()
        inst := decodeFunc(word)
        if inst < 0 {
            em.warn(InvalidInstructionWarning{em.pc - 1, word})
            ticksExecuted++
            continue
        }

        if !em.Spec.Available[inst] {
            em.warn(UnavailableInstructionWarning{em.pc - 1, inst, em.Spec})
            ticksExecuted++
            continue
        }

        handler := handlers[inst]
        ticksExecuted += handler(em, word)
    }

    em.excessTicks = ticksExecuted - ticks
}

// Copy program words from buf into program memory starting at the given address.
// The method panics if the address is out of range at any point (the size of the
// program memory is equal to 1 << em.Spec.LogProgMemSize).
func (em *Emulator) WriteProg(address uint16, buf []uint16) {
    for _, word := range buf {
        em.prog[address] = word
        address++
    }
}

func (em *Emulator) fetchProgWord() (word uint16) {
    word = em.prog[em.pc]
    em.pc = (em.pc + 1) & em.pcmask
    return word
}

func (em *Emulator) demap(addr uint16) (r Region) {
    // TODO: optimise
    for _, r := range em.regions {
        if r.Contains(addr) {
            return r
        }
    }

    return nil
}

func (em *Emulator) loadDataByte(addr uint16) uint8 {
    r := em.demap(addr)
    if r != nil {
        return r.Load(addr)
    } else {
        em.warn(UnmappedAddressWarning{em.pc - 1, addr})
        return 0
    }
}

func (em *Emulator) storeDataByte(addr uint16, val uint8) {
    r := em.demap(addr)
    if r != nil {
        r.Store(addr, val)
    } else {
        em.warn(UnmappedAddressWarning{em.pc - 1, addr})
    }
}

func (em *Emulator) push(val uint8) {
    em.storeDataByte(em.sp, val)
    em.sp--
}

func (em *Emulator) pop() uint8 {
    em.sp++
    return em.loadDataByte(em.sp)
}

func (em *Emulator) pushPC() {
    if em.Spec.LogProgMemSize > 16 { // pc is 3 bytes
        em.push(uint8(em.pc >> 16))
    }

    em.push(uint8(em.pc >> 8))
    em.push(uint8(em.pc))
}

func (em *Emulator) popPC() {
    em.pc = uint32(em.pop())
    em.pc |= uint32(em.pop()) << 8

    if em.Spec.LogProgMemSize > 16 { // pc is 3 bytes
        em.pc |= uint32(em.pop()) << 16
    }
}

func (em *Emulator) readPort(bankNum uint, index uint16) uint8 {
    port := em.ports[bankNum][index]
    if port == nil {
        em.warn(UnmappedPortWarning{em.pc - 1, bankNum, index})
        return 0
    }

    return port.Read()
}

func (em *Emulator) writePort(bankNum uint, index uint16, val uint8) {
    port := em.ports[bankNum][index]
    if port == nil {
        em.warn(UnmappedPortWarning{em.pc - 1, bankNum, index})
        return
    }

    port.Write(val)
}

// Skip the next instruction. Returns 1 if one word was skipped or 2 if two
// words were skipped.
func (em *Emulator) skip() (cycles uint) {
    word := em.fetchProgWord()
    var inst avr.Instruction

    if em.Spec.Family == spec.ReducedCore {
        inst = DecodeRC(word)
    } else {
        inst = DecodeNonRC(word)
    }

    if inst.IsTwoWord() {
        em.fetchProgWord()
        return 2
    }
    return 1
}

// Log a warning, if warning logging is enabled. Warnings include events such as
// * invalid instructions
// * instructions not available on this particular MCU
// * accesses to an unmapped data memory address
// which are ignored by a real MCU but are often indicative of software errors.
func (em *Emulator) warn(w Warning) {
    if em.logging {
        log.Printf("[avr/emulator:(*Emulator).warn] %s\n", w.String())
    }
}
