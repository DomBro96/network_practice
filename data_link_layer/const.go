package data_link_layer

const (
	F = int8(0x7E)	// frame delimiter : start of header / end of header
	A = int16(0xFF)
	C = int8(0x03)
	ESC = int16(0x7D5E)
	ESC7D = int16(0x7D5D)
)
