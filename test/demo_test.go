package test

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	buf := []byte{'0', '1', '2'}
	data := make([]byte, 10)
	copy(data, buf)
	println(string(data))
	//io.ReadFull()
	//binary.BigEndian.PutUint32()
}

func TestDemo2(t *testing.T) {
	var data []byte = []byte("你好")
	buf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data))) //将len（data）的结果放进buf前四个，比如6, 00000000 00000000 00000000 00000110   分成四个字节放进buf去。
	copy(buf[4:], data)
	fmt.Println(len(data))
	fmt.Printf("%v,%s    ", buf[:4], buf[4:])
}
