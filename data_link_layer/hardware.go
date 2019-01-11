package data_link_layer

type (
	// adapter send Mac frame on the Ethernet, can simply understand it as NIC
	Adapter struct {
		MacAdd [6]byte
	}
)


// adapter send frame to Ethernet
func (a *Adapter) SendFrame(dstAdd [6]byte, )  {

}

// adapter receive frame from Ethernet
func (a *Adapter) ReceiveFrame()  {

}


