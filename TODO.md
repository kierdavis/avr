To-do list
==========

## Short term

* Implement UART I/O package.

## Long term

* Manually test basic functionality of all instructions.
* Begin adding I/O packages to the `io` subdirectory. I/O packages should be
  enabled/disabled in the frontend with command-line flags.
    * SPI, as a middleware that can be used to connect other devices e.g.
      Ethernet controller
    * Pins, as a middleware that can be used to connect other devices e.g.
      simple LEDs & switches
* Optimise, optimise, optimise.
* Take into account 16-register processors (since the 16 registers that are provided are actually regs 16-31).

## Thoughts

* In hindsight, it would have been better practice to write the frontend early
  on, and then implement and test instructions incrementally (instead of doing
  all the implementing and then all the testing sequentially).
