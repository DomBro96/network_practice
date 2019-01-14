package data_link_layer

import (
	"errors"
	"fmt"
	"network_practice/common"
)

type (
	// adapter send Mac frame on the Ethernet, can simply understand it as NIC
	Adapter struct {
		MacAdd   string			// 8 bits mac address
		Receiver chan []byte	// channel analog receiver
	}
)

func NewAdapter(macAdd string) *Adapter {
	return  &Adapter{
		MacAdd: macAdd,
		Receiver: make(chan []byte, 1),
	}
}

// adapter send frame to Ethernet, return bytes
func (a *Adapter) SendFrame(frame *EthernetFrame) ([]byte, error) {
	preBuf, err := common.Int64ToSlice(frame.Header.Preamble)
	if err != nil {
		return nil, err
	}
	dstBuf := common.StringToSlice(frame.Header.DstAddress)
	srcBuf := common.StringToSlice(frame.Header.SrcAddress)
	ptBuf, err := common.Int16ToSlice(frame.Header.Type)
	if err != nil {
		return nil, err
	}
	fcsBuf, err := common.Uint32ToSlice(frame.Tail.FCS)
	if err != nil {
		return nil, err
	}
	frameBuf := common.AppendSlice(preBuf, dstBuf, srcBuf, ptBuf, frame.Data, fcsBuf)
	fmt.Printf("from address :" + a.MacAdd + "send frame to address: " + frame.Header.DstAddress)
	return frameBuf, nil
}

// adapter receive frame from Ethernet
func (a *Adapter) ReceiveFrame() ([]byte, error) {
	frameBuf := <- a.Receiver	// listening
	dstBuf := frameBuf[8:(8 + 6)]
	dst := common.SliceToString(dstBuf)
	if dst != a.MacAdd {
		return nil, errors.New("dst address not current host. ")
	}
	dataBuf := frameBuf[(8 + 6 + 6 + 2):(len(frameBuf) - 4)]
	if len(dataBuf) > 1500 || len(dataBuf) < 46 {
		return nil, errors.New("data length illegal. ")
	}
	fcsBuf := frameBuf[8:(len(frameBuf) - 4)]
	fcs, err := common.SliceToUint32(fcsBuf)
	if err != nil {
		return nil, err
	}
	frameFcs, err := common.SliceToUint32(frameBuf[len(frameBuf) - 4:])
	if err != nil {
		return nil, err
	}
	if fcs != frameFcs {
		return nil, errors.New("data crc check error. ")
	}
	return dataBuf, nil
}

func (a *Adapter) EncIntoFrame(dstAdd string, data []byte, pt int16) *EthernetFrame  {
	dstBuf := common.StringToSlice(dstAdd)
	srcBuf := common.StringToSlice(a.MacAdd)
	ptBuf, _ := common.Int16ToSlice(pt)
	crcBuf := common.AppendSlice(dstBuf, srcBuf, ptBuf, data)
	fcs := common.CreateCrc32(crcBuf)
	return &EthernetFrame{
		Header: EthernetFrameHeader{
			Preamble:   PREAMBLE,
			DstAddress: dstAdd,
			SrcAddress: a.MacAdd,
			Type:       pt,
		},
		Data: data,
		Tail: EthernetFrameTail{
			FCS: fcs,
		},
	}
}


// strictly, uniast, multicast broadcast are adapter features, but if there is no ethernet
// there is no practical meaning.
func (a *Adapter) Uniast(dstAdd string, data []byte, pt int16) error  {
	frame := a.EncIntoFrame(dstAdd, data, pt)
	frameBuf, err :=  a.SendFrame(frame)
	if err != nil {
		return err
	}
	da := &Adapter{
		MacAdd: dstAdd,
		Receiver: make(chan []byte, 1),
	}
	da.Receiver <- frameBuf
	go func() {
		da.ReceiveFrame()								// dst adapter receive frame
	}()
	go func() {
		a.ReceiveFrame()
	}()
	return nil
}

func (a *Adapter) Broadcast(data []byte, pt int16) ([]byte,  error) {		// broadcast dst address all bits is "1"
	frame := a.EncIntoFrame("OxEEEEEEEEEEEE", data, pt)
	frameBuf, err :=  a.SendFrame(frame)
	if err != nil {
		return nil, err
	}
	return frameBuf, err
}