package data_link_layer

import (
	"errors"
	"fmt"
	"network_practice/common"
)

type (
	// adapter send Mac frame on the Ethernet, can simply understand it as NIC
	Adapter struct {
		MacAdd   string		// 8 bits mac address
		Receiver chan int	// channel analog receiver
	}
)

func NewAdapter(macAdd string) *Adapter {
	return  &Adapter{
		MacAdd: macAdd,
		Receiver: make(chan int, 0),
	}
}

// adapter send frame to Ethernet, return bytes
func (a *Adapter) SendFrame(frame *EthernetFrame) ([]byte, error) {
	preBuf, err := common.Int64ToSlice(frame.Header.Preamble)
	if err != nil {
		return nil, err
	}
	dstBuf := common.StringToSlice(frame.Header.DesAddress)
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
	fmt.Printf("from address :" + a.MacAdd + "send frame to address: " + frame.Header.DesAddress)
	return frameBuf, nil
}

// adapter receive frame from Ethernet
func (a *Adapter) ReceiveFrame(frame []byte) ([]byte, error) {
	dataBuf := frame[(8 + 6 + 6 + 2):(len(frame) - 4)]
	if len(dataBuf) > 1500 || len(dataBuf) < 46 {
		return nil, errors.New("data length illegal. ")
	}
	fcsBuf := frame[8:(len(frame) - 4)]
	fcs, err := common.SliceToUint32(fcsBuf)
	if err != nil {
		return nil, err
	}
	frameFcs, err := common.SliceToUint32(frame[len(frame) - 4:])
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
			DesAddress: dstAdd,
			SrcAddress: a.MacAdd,
			Type:       pt,
		},
		Data: data,
		Tail: EthernetFrameTail{
			FCS: fcs,
		},
	}
}


