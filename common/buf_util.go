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

func SliceToUint32(bufByte []byte) (uint32, error)  {
	buffer := bytes.NewBuffer(bufByte)
	var u uint32
	err := binary.Read(buffer, binary.BigEndian, &u)
	return u, err
}

func Int8ToSlice(i int8) ([]byte, error)  {
	s := make([]byte, 0)
	buf := bytes.NewBuffer(s)
	err := binary.Write(buf, binary.BigEndian, i)
	bufByte := buf.Bytes()
	return bufByte, err
}

func Int16ToSlice(i int16) ([]byte, error) {
	s := make([]byte, 0)
	buf := bytes.NewBuffer(s)
	err := binary.Write(buf, binary.BigEndian, i)
	bufByte := buf.Bytes()
	return bufByte, err
}

func Int32ToSlice(i int32) ([]byte, error)  {
	s := make([]byte, 0)
	buffer := bytes.NewBuffer(s)
	err := binary.Write(buffer, binary.BigEndian, i)
	buf := buffer.Bytes()
	return buf, err
}

func Int64ToSlice(i int64) ([]byte, error)  {
	s := make([]byte, 0)
	buffer := bytes.NewBuffer(s)
	err := binary.Write(buffer, binary.BigEndian, i)
	buf := buffer.Bytes()
	return buf, err
}

func Uint32ToSlice(u uint32) ([]byte, error)  {
	s := make([]byte, 0)
	buffer := bytes.NewBuffer(s)
	err := binary.Write(buffer, binary.BigEndian, u)
	buf := buffer.Bytes()
	return buf, err
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

func StringToSlice(str string) []byte  {
	return []byte(str)
}

func SliceToString(s []byte) string  {
	return string(s)
}