package gpio

// An InputPinAdapter encapsulates access to a digital input pin.
type InputPinAdapter interface {
    GetState() bool
    SetPullupEnabled(bool)
}

// An OutputPinAdapater encapsulates access to a digital output pin.
type OutputPinAdapter interface {
    SetState(bool)
}
