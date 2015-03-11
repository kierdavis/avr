package main

import (
    "flag"
    "fmt"
    "github.com/kierdavis/avr/clock"
    "github.com/kierdavis/avr/emulator"
    "github.com/kierdavis/avr/hardware/gpio"
    "github.com/kierdavis/avr/loader/ihexloader"
    "github.com/kierdavis/avr/spec"
    "os"
)

func main() {
    flag.Parse()
    
    if flag.NArg() < 1 {
        fmt.Fprintf(os.Stderr, "usage: %s <program.hex>\n", os.Args[0])
        os.Exit(2)
    }
    
    runEmulator()
}

func runEmulator() {
    em := emulator.NewEmulator(spec.ATmega168)
    em.LogWarnings(true)
    loadProgram(em)
    setupIO(em)
    
    clk := clock.New()
    go em.Run(clk.Spawn(1))
    
    clk.Tick(1e6)
    
    fmt.Println("OK.")
}

func loadProgram(em *emulator.Emulator) {
    f, err := os.Open(flag.Arg(0))
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
        os.Exit(1)
    }
    
    err = ihexloader.Load(em, f)
    f.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
        os.Exit(1)
    }
}

func setupIO(em *emulator.Emulator) {
    gpio := gpio.New('B', 8)
    gpio.SetOutputAdapter(5, &PrintingOutputPinAdapter{Label: "LED"})
    gpio.AddTo(em)
}

type PrintingOutputPinAdapter struct {
    Label string
    Prev bool
}

func (a *PrintingOutputPinAdapter) SetState(state bool) {
    if state != a.Prev {
        a.Prev = state
        if state {
            fmt.Printf("%s changed to high\n", a.Label)
        } else {
            fmt.Printf("%s changed to low\n", a.Label)
        }
    } else {
        fmt.Printf("%s remained the same\n", a.Label)
    }
}
