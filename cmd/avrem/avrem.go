package main

import (
    "flag"
    "fmt"
    "github.com/kierdavis/avr/emulator"
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
    
    em.Run(1000)
    
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
