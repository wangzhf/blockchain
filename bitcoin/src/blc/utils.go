package blc

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 将一个int64的整数转为二进制后，每8bit一个byte
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	// 转为[]byte并返回
	return buff.Bytes()
}
