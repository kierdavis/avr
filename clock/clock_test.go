package clock

import (
    "fmt"
    "time"
)

func ExampleClock() {
    // Create a new master clock.
    clk := New()
    
    // Spawn a worker goroutine.
    go func() {
        for {
            now := clk.Now()
            
            // Do some processing that takes 6 clock ticks to complete.
            fmt.Println("Worker 1 tick")
            
            // Wait for other workers to catch up.
            clk.Await(now + 6)
        }
    }()
    
    // Spawn a worker goroutine.
    go func() {
        for {
            now := clk.Now()
            
            // Do some processing that takes 2 clock ticks to complete.
            fmt.Println("Worker 2 tick")
            
            // Wait for other workers to catch up.
            clk.Await(now + 2)
        }
    }()
    
    // Create a signal that can be used to stop the clock.
    stopSig := NewSignal()
    go func() {
        time.Sleep(time.Second * 5)
        stopSig.Trigger()
    }()
    
    // Run the clock at a frequency of 2 Hz (a period of half a second).
    clk.Run(time.Second / 2, stopSig)
}
