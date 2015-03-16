// avrmakedec is a tool to emit an efficient AVR instruction decoder from a
// specification of the binary syntax of the instruction set. It is used when
// preprocessing the AVR package and it designed to be used with the
// 'go generate' command.
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
)

var (
    FlagKind   = flag.String("kind", "flat", "the kind of decoder to create (choices: 'flat' (single switch block), 'lutc' (lookup table initialised at compile time), 'lutr' (lookup table initialised at runtime))")
    FlagPkg    = flag.String("pkg", "", "the package name to use for the generated file")
    FlagOutput = flag.String("output", "decoder.go", "the file to write output to")
)

func main() {
    log.SetFlags(0)
    log.SetPrefix("avrmakedec: ")
    flag.Parse()

    var kind Kind
    switch *FlagKind {
    case "flat":
        kind = Flat
    case "lutc":
        kind = Lutc
    case "lutr":
        kind = Lutr
    default:
        log.Printf("error: bad value for -kind: %s (expected one of 'flat')", *FlagKind)
        os.Exit(2)
    }

    if *FlagPkg == "" {
        log.Printf("error: -pkg must be specified")
        os.Exit(2)
    }

    var g Generator
    g.Generate(*FlagPkg, kind)

    filename := *FlagOutput
    log.Printf("writing %s", filename)

    f, err := os.Create(filename)
    if err != nil {
        log.Printf("error: could not open %s for writing: %s", filename, err)
        os.Exit(1)
    }

    src := string(g.Format())
    fmt.Fprint(f, src)
    f.Close()
}
