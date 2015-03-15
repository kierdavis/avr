package main

func (g *Generator) GenerateFlat() {
    g.Printf("  switch {\n")
    for _, instDef := range InstDefs {
        var format string
        switch instDef.RCMode {
        case RC:
            format = "    case word & 0x%04X == 0x%04X && reducedCore:\n"
        case NotRC:
            format = "    case word & 0x%04X == 0x%04X && !reducedCore:\n"
        case Either:
            format = "    case word & 0x%04X == 0x%04X:\n"
        }
        g.Printf(format, instDef.Mask, instDef.Match)
        g.Printf("      return avr.%s\n", instDef.Inst.String())
    }
    g.Printf("  }\n")
    g.Printf("  return -1\n")
}
