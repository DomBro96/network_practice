package common

import (
	"bytes"
	"encoding/binary"
)

func SliceToInt8(bufByte []byte) (int8, error) {
	buf := bytes.NewBuffer(bufByte)
	var i int8
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
}

func SliceToInt16(bufByte []byte) (int16, error)  {
	buf := bytes.NewBuffer(bufByte)
	var i int16
	err := binary.Read(buf, binary.BigEndian, &i)
	return i ,err
}

func SliceToInt32(bufByte []byte) (int32, error) {
	buf := bytes.NewBuffer(bufByte)
	var i int32
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
}

func Int8ToSlice(i int8) ([]byte, error)  {
	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)
	err := binary.Write(buf, binary.BigEndian, i)
	bufByte := buf.Bytes()
	return bufByte, err
}

func Int16ToSlice(i int16) ([]byte, error) {
	s1 := make([]byte, 0)
	buf := bytes.NewBuffer(s1)
	err := binary.Write(buf, binary.BigEndian, i)
	bufByte := buf.Bytes()
	return bufByte, err
}



func AppendSlice(bs ...[]byte) []byte {
	buf := make([]byte, 0)
	for _, b := range bs {
		buf = append(buf, b...)
	}
	return buf
}

