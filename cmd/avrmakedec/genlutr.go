package main

func (g *Generator) GenerateLutr() {
    g.GenerateFlatFunc("decodeFlatNonRC", false)
    g.GenerateFlatFunc("decodeFlatRC", true)
    g.GenerateLutrFunc("NonRC", false)
    g.GenerateLutrFunc("RC", true)
}

func (g *Generator) GenerateLutrFunc(suffix string, rc bool) {
    g.Printf("func Decode%s(word uint16) avr.Instruction {\n", suffix)
    g.Printf("  return decodeLut%s[word]\n", suffix)
    g.Printf("}\n")
    g.Printf("var decodeLut%s [65536]avr.Instruction\n", suffix)
    g.Printf("func init() {\n")
    g.Printf("  for i := 0; i < 65536; i++ {\n")
    g.Printf("    decodeLut%s[i] = decodeFlat%s(uint16(i))\n", suffix, suffix)
    g.Printf("  }\n")
    g.Printf("}\n")
}
