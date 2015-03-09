To-do list
==========

## Short term

* Finish implementing instruction set.
* Finalise `ATmega48/88/168` and `ATtiny4/5/9/10` specs by adding a `Supports`
  field.

## Long term

* Write an Intel Hex format parsing package.
* Write an emulator frontend (does not need to do any I/O yet).
* Manually test basic functionality of all instructions.
* Begin adding I/O packages to the `io` subdirectory. I/O packages should be
  enabled/disabled in the frontend with command-line flags.
    * UART (command-line interface)
    * SPI, as a middleware that can be used to connect other devices e.g.
      Ethernet controller
    * Pins, as a middleware that can be used to connect other devices e.g.
      simple LEDs & switches
* Optimise, optimise, optimise

## Thoughts

* In hindsight, it would have been better practice to write the frontend early
  on, and then implement and test instructions incrementally (instead of doing
  all the implementing and then all the testing sequentially).
