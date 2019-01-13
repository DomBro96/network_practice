package data_link_layer

type (

	PPPFrameHeader struct {
		Flag int8		// frame start flag
		A int8		// meaningless
		C int8			// meaningless
		P int16			// protocol
	}

	PPPFrameTail struct {
		Fcs  int16 // fcs
		Flag int8  // frame end flag
	}

	PPPFrame struct {
		Header PPPFrameHeader
		Data   []byte  // data message
		Tail   PPPFrameTail
	}


	EthernetFrameHeader struct {
		Preamble   int64		// 7 bytes preamble 1010....1010, 1 byte start delimiter	10101011
		DesAddress string		// 6 bytes des address
		SrcAddress string		// 6 bytes src address
		Type 	int16			// protocol type
	}

	EthernetFrameTail struct {
		FCS uint32			    // 32 bits fcs
	}

	EthernetFrame struct {
		Header EthernetFrameHeader
		Data   []byte
		Tail   EthernetFrameTail
	}
)



