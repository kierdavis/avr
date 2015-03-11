package main

import (
    "flag"
    "fmt"
    "github.com/kierdavis/avr/clock"
    "github.com/kierdavis/avr/emulator"
    "github.com/kierdavis/avr/hardware/gpio"
    "github.com/kierdavis/avr/hardware/timer"
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
    clk := clock.New()
    
    em := emulator.NewEmulator(spec.ATmega168)
    em.SetLogging(true)
    loadProgram(em)
    setupIO(em, clk)
    
    go em.Run(clk.Spawn(1))
    
    for i := 0; i < 1e8; i++ {
        clk.Tick(10)
    }
    
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

func setupIO(em *emulator.Emulator, clk *clock.Master) {
    gpioB := gpio.New('B', 8)
    gpioB.SetOutputAdapter(5, &PrintingOutputPinAdapter{Label: "LED"})
    gpioB.AddTo(em)
    
    t0 := timer.New(0)
    t0.SetLogging(true)
    t0.AddTo(em)
    go t0.Run(clk.Spawn(1))
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
