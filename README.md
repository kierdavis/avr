avr
===

`avr` is an implementation of Atmel's AVR microprocessor specification in the Go
programming language.

Note: this package is a work-in-progress. More to come shortly.

## Features

* Implements entire AVR instruction set, with the exception of `BREAK`, `DES`, `SLEEP`, `SPM` and `WDR`.
* Emulates various hardware modules:
    * digital GPIO pins
    * timer/counter
* Accurately supports individual MCUs:
    * ATtiny4/5/9/10
    * ATmega48/88/168
    * more to come soon!

## Installation

`avr` is go-gettable:

    # go get github.com/kierdavis/avr/cmd/avrem

## Usage

The `avrem` command (in the "cmd" subdirectory) is a command-line interface to
the emulator. To try it out, run it with an Intel HEX file like so:

    # avrem program.hex

Flags include `-mcu` to specify the name of the MCU spec to use, `-mcus` to list
the names of all available MCU specs, and `-freq` to specify the execution
frequency.

The "programs" subdirectory contains example programs. Many of these are
[Arduino][arduino] programs, and have precompiled IHEX program files for a
number of MCUs present in the same directory.

Running the "blink" example for a 16 MHz ATmega168:

    # avrem -mcu mega168 -freq 16 programs/blink/blink-atmega168.hex

## Performance

The maximum unthrottled clock rate approaches 35 MHz on my 2.3 GHz Intel i7
processor. Real AVR microprocessors range from 1 to 20 MHz, so this software can
theoretically simulate an mid-range AVR processor twice as quickly as the
physical implementation.

## Packages

* `github.com/kierdavis/avr` - miscellaneous shared code
* `github.com/kierdavis/avr/clock` - manages synchronisation between concurrent processes of emulator
* `github.com/kierdavis/avr/emulator` - implementation of CPU emulator
* `github.com/kierdavis/avr/hardware/gpio` - implementation of digital GPIO pins
* `github.com/kierdavis/avr/hardware/timer` - implementation of timer/counter module
* `github.com/kierdavis/avr/loader/ihexloader` - links Intel HEX file parser with loading programs into emulators
* `github.com/kierdavis/avr/spec` - specifications of the many different models of AVR processor (MCUs)
