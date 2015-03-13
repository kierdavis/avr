package clock

import (
    "log"
    "time"
)

type Process interface {
    Run(ticks uint)
}

type Clock struct {
    procs []Process
    lastFreqCheck time.Time
    lastThrottle time.Time
    ticksSinceFreqCheck uint
    ticksSinceThrottle uint
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
    c.ticksSinceFreqCheck += ticks
    c.ticksSinceThrottle += ticks
}

func (c *Clock) MonitorFrequency() (freq float64) {
    now := time.Now()
    if !c.lastFreqCheck.IsZero() {
        dur := now.Sub(c.lastFreqCheck)
        secs := float64(dur) / float64(time.Second)
        freq = float64(c.ticksSinceFreqCheck) / secs
    }
    
    c.lastFreqCheck = now
    c.ticksSinceFreqCheck = 0
    return freq
}

func (c *Clock) LogFrequency() {
    freq := c.MonitorFrequency()
    log.Printf("[avr/clock] Running at: %.1f MHz (%.1f ns/tick)", freq / 1e6, 1e9 / freq)
}

func (c *Clock) Throttle(freq float64) {
    periodSecs := 1 / freq
    period := time.Duration(periodSecs * float64(time.Second))
    
    if !c.lastThrottle.IsZero() {
        targetTime := c.lastThrottle.Add(period * time.Duration(c.ticksSinceThrottle))
        sleepDur := targetTime.Sub(time.Now())
        if sleepDur > 0 {
            time.Sleep(sleepDur)
        }
    }
    
    c.lastThrottle = time.Now()
    c.ticksSinceThrottle = 0
}
