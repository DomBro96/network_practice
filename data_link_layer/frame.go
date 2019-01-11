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
		
	}

	EthernetFrameTail struct {

	}

	EthernetFrame struct {
		Header EthernetFrameHeader
		Data   []byte
		Tail   EthernetFrameTail
	}
)



