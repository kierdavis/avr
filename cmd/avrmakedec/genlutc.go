package main

func (g *Generator) GenerateLutc() {
    g.GenerateLutcFunc("DecodeNonRC", "decodeLutNonRC", false)
    g.GenerateLutcFunc("DecodeRC", "decodeLutRC", true)
}

func (g *Generator) GenerateLutcFunc(funcName string, lutName string, rc bool) {
    g.Printf("func %s(word uint16) avr.Instruction {\n", funcName)
    g.Printf("  return %s[word]\n", lutName)
    g.Printf("}\n")
    g.Printf("var %s = [65536]avr.Instruction{\n", lutName)
    for i := 0; i < 65536; i++ {
        g.Printf("  avr.%s,\n", Decode(uint16(i), rc))
    }
    g.Printf("}\n")
}
