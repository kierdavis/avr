// Package clock implements a global synchronisation mechanism based on this one
// by nwidger:
// http://nwidger.github.io/blog/post/writing-an-nes-emulator-in-go-part-2/
package clock

import (
    "time"
)

// A Signal allows a single receiver to Wait for a single sender to Trigger an
// event.
type Signal struct {
    C chan struct{}
}

func NewSignal() (s Signal) {
    return Signal{make(chan struct{}, 1)}
}

func (s Signal) Trigger() {
    s.C <- struct{}{}
}

func (s Signal) Wait() {
    <-s.C
}

// A Clock encapsulates a synchronisation mechanism. A user of a Clock should
// call Now to obtain the current tick number, do some processing, and then
// call Await with what the tick number should be after the processing has
// completed.
type Clock interface {
    Now() uint64
    Await(uint64)
}

type Master struct {
    now uint64
    waiting map[uint64][]Signal
}

func New() (m *Master) {
    return &Master{
        now: 0,
        waiting: make(map[uint64][]Signal),
    }
}

func (m *Master) Now() uint64 {
    return m.now
}

func (m *Master) Await(targetTick uint64) {
    if m.now < targetTick {
        sig := NewSignal()
        m.waiting[targetTick] = append(m.waiting[targetTick], sig)
        sig.Wait()
    }
}

func (m *Master) Tick() {
    m.now++
    waiting := m.waiting[m.now]
    if len(waiting) > 0 {
        for _, sig := range waiting {
            sig.Trigger()
        }
        delete(m.waiting, m.now)
    }
}

func (m *Master) TickN(n uint) {
    for n > 0 {
        m.Tick()
        n--
    }
}

func (m *Master) Run(period time.Duration, stop Signal) {
    batchSize := 100
    
    batchStartTime := time.Now()
    for {
        //k := 10000
        //t1 := time.Now()
        //for i := 0; i < k; i++ {
            batchStartTime = batchStartTime.Add(period * time.Duration(batchSize))
            sleepTime := batchStartTime.Sub(time.Now())
            if sleepTime > 0 {
                time.Sleep(sleepTime)
            }
            
            for j := 0; j < batchSize; j++ {
                m.Tick()
            }
            
            select {
            case <-stop.C:
                return
            default:
            }
        //}
        //t2 := time.Now()
        //
        //secs := float64(t2.Sub(t1)) / float64(time.Second)
        //fmt.Printf("%d ticks in %f msecs\t(effective freq %f MHz)\n", k*batchSize, secs * 1e3, (float64(k*batchSize) / secs) / 1e6)
    }
}

type Divider struct {
    parent Clock
    divisor uint64
}

// A Divider takes a parent Clock and produces a new Clock that ticks once for
// every divisor ticks of the parent. For example, a divisor of 3 means that
// the Divider ticks a third as often as the parent.
func NewDivider(parent Clock, divisor uint64) (d *Divider) {
    return &Divider{parent: parent, divisor: divisor}
}

func (d *Divider) Now() uint64 {
    return d.parent.Now() / d.divisor
}

func (d *Divider) Await(targetTick uint64) {
    d.parent.Await(targetTick * d.divisor)
}
