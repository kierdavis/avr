package main

func (g *Generator) GenerateFlat() {
    g.GenerateFlatFunc("DecodeNonRC", false)
    g.GenerateFlatFunc("DecodeRC", true)
}

func (g *Generator) GenerateFlatFunc(name string, rc bool) {
    g.Printf("func %s(word uint16) avr.Instruction {\n", name)
    g.Printf("  switch {\n")
    for _, instDef := range InstDefs {
        if rc {
            if instDef.RCMode == NotRC {
                continue
            }
        } else {
            if instDef.RCMode == RC {
                continue
            }
        }
        g.Printf("    case word & 0x%04X == 0x%04X:\n", instDef.Mask, instDef.Match)
        g.Printf("      return avr.%s\n", instDef.Inst.String())
    }
    g.Printf("  }\n")
    g.Printf("  return -1\n")
    g.Printf("}\n")
}
