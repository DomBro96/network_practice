package data_link_layer

const (
	F     = 0x7E	  // frame delimiter : start of header / end of header
	A     = 0xFF
	C     = 0x03
	ESC   = 0x7D5E
	ESC7D = 0x7D5D
	CRC16 = 98309 	 // X^16 + X^15 + X^2 + 1
	IP = 0x0021 	 // IP  protocol
	LCP = 0xC021	 // LCP protocol

)
