package common

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
)

func SliceToInt8(bufByte []byte) (int8, error) {
	buf := bytes.NewBuffer(bufByte)
	var i int8
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
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

func CreateCrc32(buf []byte) uint32 {
	crcValue := crc32.ChecksumIEEE(buf)
	return crcValue
}