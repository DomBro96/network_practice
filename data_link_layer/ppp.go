package data_link_layer

import (
	"errors"
	"network_practice/common"
)

type (

	Sender int64	// Sender and Receiver are mac address, 48 digit
	Receiver int64

	PPP struct {
		sender   Sender
		receiver Receiver
		frame *PPPFrame
	}
)

func NewPPP(s Sender, r Receiver) *PPP  {
	return &PPP{
		sender: s,
		receiver: r,
	}
}

func (ppp *PPP)EncIntoFrame(pt int16, data []byte) (*PPPFrame, error) {
	paddingData, err := ppp.DataPadding(data)
	if err != nil {
		return nil, err
	}
	if len(paddingData) > 1500 || len(paddingData) < 46{ // MTU range in 46 to 1500 byte
		return nil, errors.New("data length illegal. ")
	}
	fcs, err := ppp.CreateFcs(paddingData)	// create data fcs check num, add to the tail of frame
	if err != nil {
		return nil, err
	}
	header := PPPFrameHeader{
		Flag: F,
		A: A,
		C: C,
		P: pt,
	}
	tail := PPPFrameTail{
		Fcs:  fcs,
		Flag: F,
	}
	frame := &PPPFrame{
		Header: header,
		Data:   paddingData,
		Tail:   tail,
	}
	ppp.frame = frame
	return frame, nil
}

// receiver resolution frame
func (ppp *PPP) ResFromFrame (frameBuf []byte) ([]byte, error) {
	data := frameBuf[5 : len(frameBuf) - 3]
	data, err := ppp.DataParsing(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// src data shift left 16 bits, because CRC_16 is 17 bits
// residual data32 and CRC_16
func (ppp *PPP)CreateFcs(data []byte) (int16, error) {
	data32, err := common.SliceToInt32(data)
	if err != nil {
		return 0, err
	}
	fcs :=  (data32 << 16) % CRC16
	return int16(fcs), nil
}

// Using Asynchronous transfer frame, sender need padding data in frame.
// Replace the same as 'flag' data in PPP data to esc.
func (ppp *PPP) DataPadding(data []byte) ([]byte, error) {
	i := 0
	for {
		length := len(data)
		if length - i < 1 {
			break
		}
		s := data[i:i+1]
		num, err := common.SliceToInt8(s)
		if err != nil {
			return nil, err
		}
		if num == F {
			esc, err := common.Int16ToSlice(ESC)
			if err != nil {
				return nil, err
			}
			if i != 0 {
				data = common.AppendSlice(data[:i], esc, data[i+1:])
			}else {
				data = common.AppendSlice(esc, data[i+1:])
			}
			i += 2
		}else if num == 0x7D {
			esc7d, err := common.Int16ToSlice(ESC7D)
			if err != nil {
				return nil, err
			}
			if i != 0 {
				data = common.AppendSlice(data[:i], esc7d, data[i+1:])
			}else {
				data = common.AppendSlice(esc7d, data[i+1:])
			}
			data = common.AppendSlice(data[:i], esc7d, data[i+1:])
			i += 2
		}else {
			i += 1
		}
	}
	return data, nil
}

// receiver receive frame, need parsing data in frame.
// replace ESC or ESC7D to src frame data.
func (ppp *PPP) DataParsing(data []byte) ([]byte, error) {
	i := 0
	for {
		length := len(data)
		if length - i < 2 {
			break
		}
		s := data[i:i+2]
		num, err := common.SliceToInt16(s)
		if err != nil {
			return nil, err
		}
		if num == ESC {
			f, err := common.Int8ToSlice(F)
			if err != nil {
				return nil, err
			}
			if i != 0 {
				data = common.AppendSlice(data[:i], f, data[i+2:])
			}else {
				data = common.AppendSlice(f, data[i+2:])
			}
			i += 2
		}else if num == ESC7D {
			esc, err := common.Int8ToSlice(0x7D)
			if err != nil {
				return nil, err
			}
			if i != 0 {
				data = common.AppendSlice(data[:i], esc, data[i+2:])
			}else {
				data = common.AppendSlice(esc, data[i+2:])
			}
			data = common.AppendSlice(esc, data[i+2:])
			i += 2
		}else {
			i += 1
		}
	}
	return data, nil
}


