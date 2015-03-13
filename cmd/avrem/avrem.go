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
    "runtime/pprof"
)

var cpuProfile = flag.String("cpuprofile", "", "filename to write profiling data to")
var throttleFreq = flag.Float64("freq", 0, "clock frequency to throttle emulation to (in MHz, 0 to run unthrottled)")
var mcu = flag.String("mcu", "mega168", "select specific MCU to use (use -mcus to list available MCU names)")
var mcus = flag.Bool("mcus", false, "list MCU names")

var mcuMap = map[string]*spec.MCUSpec{
    "tiny4": spec.ATtiny4,
    "tiny5": spec.ATtiny5,
    "tiny9": spec.ATtiny9,
    "tiny10": spec.ATtiny10,
    "mega48": spec.ATmega48,
    "mega88": spec.ATmega88,
    "mega168": spec.ATmega168,
}

func main() {
    flag.Parse()
    
    if flag.NArg() < 1 {
        fmt.Fprintf(os.Stderr, "usage: %s <program.hex>\n", os.Args[0])
        os.Exit(2)
    }
    
    if *mcus {
        fmt.Printf("MCUs available for use with -mcu:\n")
        for name := range mcuMap {
            fmt.Printf("  %s\n", name)
        }
        return
    }

    if *cpuProfile != "" {
        f, err := os.Create(*cpuProfile)
        if err != nil {
            fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
            os.Exit(1)
        }
        pprof.StartCPUProfile(f)
        defer func() {
            pprof.StopCPUProfile()
            f.Close()
        }()
    }
    
    runEmulator()
}

func runEmulator() {
    spec, ok := mcuMap[*mcu]
    if !ok {
        fmt.Fprintf(os.Stderr, "error: invalid value for -mcu (try -mcus for a list)\n")
        os.Exit(2)
    }
    log.Printf("[avr/cmd/avrem] using MCU spec: %s", spec.Label)

    clk := clock.New()
    
    em := emulator.NewEmulator(spec)
    em.SetLogging(true)
    clk.Add(em)
    
    loadProgram(em)
    setupIO(em, clk)
    
    throttleFreq_ := *throttleFreq
    
    for i := 0; i < 100; i++ {
        clk.LogFrequency()
        
        for i := 0; i < 1e5; i++ {
            clk.Run(20)
        }
        
        if throttleFreq_ != 0 {
            clk.Throttle(throttleFreq_ * 1e6)
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
            log.Printf("[avr/cmd/avrem] %s changed to high", a.Label)
        } else {
            log.Printf("[avr/cmd/avrem] %s changed to low", a.Label)
        }
    }
}
