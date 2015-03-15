package emulator

// Generate a decoder using avrmakedec
//go:generate go install github.com/kierdavis/avr/cmd/avrmakedec
//go:generate avrmakedec -output decode_gen.go -pkg emulator -kind lutr
