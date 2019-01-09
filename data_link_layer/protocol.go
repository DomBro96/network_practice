package data_link_layer

import (
	"errors"
	"network_practice/common"
)

type (
	PPP struct {
		Sender   int64
		Receiver int64
	}
)



func (ppp *PPP)PPPIntoFrame(pt int16, data []byte) (*PPPFrame, error) {
	if len(data) > 1500 || len(data) < 46{		// MTU range in 46 to 1500 byte 
		return nil, errors.New("data length not illegal. ")
	}
	d, err := ppp.BytePadding(data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (ppp *PPP)BytePadding(data []byte) ([]byte, error) {	// replace the same as 'flag' data in PPP data
	i := 0
	for {
		length := len(data)
		if length - i < 8 {
			break
		}
		s := data[i:i+8]
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
				data = common.AppendSlice(data[:i], esc, data[i+8:])
			}else {
				data = common.AppendSlice(esc, data[i+8:])
			}
			i += 16
		}else if num == 0x7D {
			esc7d, err := common.Int16ToSlice(ESC7D)
			if err != nil {
				return nil, err
			}
			if i != 0 {
				data = common.AppendSlice(data[:i], esc7d, data[i+8:])
			}else {
				data = common.AppendSlice(esc7d, data[i+8:])
			}
			data = common.AppendSlice(data[:i], esc7d, data[i+8:])
			i += 16
		}else {
			i += 8
		}
	}
	return data, nil
}



