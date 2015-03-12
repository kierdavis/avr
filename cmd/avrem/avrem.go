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
    "log"
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
    clk.Add(em)
    
    loadProgram(em)
    setupIO(em, clk)
    
    /*
    // main loop
    m := 20
    n := 10000000 / m
    for {
        t1 := time.Now()
        for i := 0; i < n; i++ {
            em.Run(uint(m))
            t0.Run(uint(m))
            totalTicks += uint64(m)
        }
        t2 := time.Now()
        
        secs := float64(t2.Sub(t1)) / float64(time.Second)
        fmt.Printf("Running at: %f MHz (%f ns/tick)\n", float64(n*m) / (secs * 1e6), (secs * 1e9) / float64(n*m))
    }
    */
    
    for {
        freq := clk.MonitorFrequency()
        log.Printf("[avr/cmd/avrem] Running at: %f MHz (%f ns/tick)", freq / 1e6, 1e9 / freq)
        
        for i := 0; i < 1e5; i++ {
            clk.Run(20)
        }
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

func setupIO(em *emulator.Emulator, clk *clock.Clock) {
    gpioB := gpio.New('B', 8)
    gpioB.SetOutputAdapter(5, &PrintingOutputPinAdapter{Label: "LED"})
    gpioB.AddTo(em)
    
    t0 := timer.New(0)
    t0.SetLogging(true)
    t0.AddTo(em)
    clk.Add(t0)
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
