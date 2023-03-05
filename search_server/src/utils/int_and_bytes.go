package utils

import (
	"bytes"
	"encoding/binary"
)

// IntToBytes int转二进制数组
func IntToBytes(n int) ([]byte, error) {
	data := int64(n)
	byteBuf := bytes.NewBuffer([]byte{})
	if err := binary.Write(byteBuf, binary.BigEndian, data); err != nil {
		return nil, err
	}
	return byteBuf.Bytes(), nil
}

// BytesToInt 二进制数组转int
func BytesToInt(bys []byte) (int64, error) {
	byteBuff := bytes.NewBuffer(bys)
	var data int64
	if err := binary.Read(byteBuff, binary.BigEndian, &data); err != nil {
		return 0, err
	}
	return data, nil
}
