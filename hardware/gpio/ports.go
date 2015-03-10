package gpio

// Implementation of PORTx I/O port
type port struct {
    g *GPIO
}

func (p port) Read() uint8 {
    return p.g.outputs
}

func (p port) Write(x uint8) {
    diff := p.g.outputs ^ x
    p.g.outputs = x
    p.g.updateOutputs(diff)
}

// Implementation of DDRx I/O port
type ddr struct {
    g *GPIO
}

func (p ddr) Read() uint8 {
    return p.g.dirs
}

func (p ddr) Write(x uint8) {
    diff := p.g.dirs ^ x
    p.g.dirs = x
    p.g.updateOutputs(diff)
}

// Implementation of PINx I/O port
type pin struct {
    g *GPIO
}

func (p pin) Read() uint8 {
    return p.g.getInputs()
}

func (p pin) Write(x uint8) {
    diff := p.g.pullups ^ x
    p.g.pullups = x
    p.g.updateOutputs(diff)
}

// TODO: PINx does not have read-modify-write capabilities.
// Possible solution: add more methods on Port to enable changing individual bits
// Alternatively, add one method to return the value that Write would modify
