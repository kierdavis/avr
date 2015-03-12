package clock

import (
    "time"
)

type Process interface {
    Run(ticks uint)
}

type Clock struct {
    procs []Process
    t time.Time
    ticks uint
}

func New() (c *Clock) {
    return &Clock{}
}

func (c *Clock) Add(p Process) {
    c.procs = append(c.procs, p)
}

func (c *Clock) Run(ticks uint) {
    for _, proc := range c.procs {
        proc.Run(ticks)
    }
    c.ticks += ticks
}

func (c *Clock) MonitorFrequency() (freq float64) {
    now := time.Now()
    if !c.t.IsZero() {
        dur := now.Sub(c.t)
        secs := float64(dur) / float64(time.Second)
        freq = float64(c.ticks) / secs
    }
    
    c.t = now
    c.ticks = 0
    return freq
}
