// Package gpio implements a general purpose digital I/O port.
// Works with:
//   ATmega48/88/168
package gpio

import (
    "fmt"
    "github.com/kierdavis/avr/emulator"
)

// TODO: pullup-disable bit in MCUCR (see http://www.atmel.com/Images/doc2545.pdf, page 87)

// Values in DDR register
const (
    Input = 0
    Output = 1
)

type GPIO struct {
    letter byte
    width uint
    dirs uint8
    outputs uint8
    pullups uint8
    inputAdapters [8]InputPinAdapter
    outputAdapters [8]OutputPinAdapter
}

func New(portLetter byte, width uint) (g *GPIO) {
    return &GPIO{
        letter: portLetter,
        width: width,
    }
}

func (g *GPIO) SetInputAdapter(pinNumber uint, adapter InputPinAdapter) {
    g.inputAdapters[pinNumber] = adapter
}

func (g *GPIO) SetOutputAdapter(pinNumber uint, adapter OutputPinAdapter) {
    g.outputAdapters[pinNumber] = adapter
}

func (g *GPIO) AddTo(em *emulator.Emulator) {
    em.RegisterPortByName(fmt.Sprintf("PORT%c", g.letter), port{g})
    em.RegisterPortByName(fmt.Sprintf("DDR%c", g.letter), ddr{g})
    em.RegisterPortByName(fmt.Sprintf("PIN%c", g.letter), pin{g})
}

// called when a PIN port is read
func (g *GPIO) getInputs() (x uint8) {
    if g.getInput(7) {x |= 0x80}
    if g.getInput(6) {x |= 0x40}
    if g.getInput(5) {x |= 0x20}
    if g.getInput(4) {x |= 0x10}
    if g.getInput(3) {x |= 0x08}
    if g.getInput(2) {x |= 0x04}
    if g.getInput(1) {x |= 0x02}
    if g.getInput(0) {x |= 0x01}
    return x
}

func (g *GPIO) getInput(pinNumber uint) bool {
    adapter := g.inputAdapters[pinNumber]
    if adapter != nil {
        return adapter.GetState()
    } else {
        return false
    }
}

// called when a PORT/DDR/PIN port is written
func (g *GPIO) updateOutputs(changed uint8) {
    if changed & 0x80 != 0 {g.updateOutput(7)}
    if changed & 0x40 != 0 {g.updateOutput(6)}
    if changed & 0x20 != 0 {g.updateOutput(5)}
    if changed & 0x10 != 0 {g.updateOutput(4)}
    if changed & 0x08 != 0 {g.updateOutput(3)}
    if changed & 0x04 != 0 {g.updateOutput(2)}
    if changed & 0x02 != 0 {g.updateOutput(1)}
    if changed & 0x01 != 0 {g.updateOutput(0)}
}

func (g *GPIO) updateOutput(pinNumber uint) {
    if pinNumber < g.width { // ignore out-of-range pins
        if (g.dirs >> pinNumber) & 1 == Input {
            adapter := g.inputAdapters[pinNumber]
            if adapter != nil {
                adapter.SetPullupEnabled((g.pullups >> pinNumber) & 1 != 0)
            }
        
        } else {
            adapter := g.outputAdapters[pinNumber]
            if adapter != nil {
                adapter.SetState((g.outputs >> pinNumber) & 1 != 0)
            }
        }
    }
}
