// Package ihexloader encapsulates loading of IHEX files into an Emulator.
package ihexloader

import (
    "github.com/kierdavis/avr/emulator"
    "github.com/kierdavis/ihex-go"
    "io"
)

// Load parses an IHEX file from r and loads the program data contained in
// it into em.
func Load(em *emulator.Emulator, r io.Reader) (err error) {
    dec := ihex.NewDecoder(r)
    buf := make([]uint16, 0, 8)
    
    for dec.Scan() {
        rec := dec.Record()
        if rec.Type == ihex.Data {
            buf = buf[:0]
            for i := 0; i+1 < len(rec.Data); i += 2 {
                lo := uint16(rec.Data[i])
                hi := uint16(rec.Data[i+1])
                buf = append(buf, (hi << 8) | lo)
            }
            
            em.WriteProg(rec.Address >> 1, buf)
        }
    }
    
    return dec.Err()
}
