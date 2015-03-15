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
	FlagKind = flag.String("kind", "flat", "the kind of decoder to create (choices: 'flat')")
	FlagPkg = flag.String("pkg", "", "the package name to use for the generated file")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("avrmakedec: ")
	flag.Parse()
	
	var kind Kind
	switch *FlagKind {
	case "flat":
		kind = Flat
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
	fmt.Print(string(g.Format()))
}
