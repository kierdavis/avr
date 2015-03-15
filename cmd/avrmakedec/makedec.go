// avr-mkdec is a tool to emit an efficient AVR instruction decoder from a
// specification of the binary syntax of the instruction set. It is used when
// preprocessing the AVR package and it designed to be used with the
// 'go generate' command.
package main

import (
	"fmt"
)

func main() {
	var g Generator
	g.Generate("emulator", Flat)
	fmt.Print(string(g.Format()))
}
