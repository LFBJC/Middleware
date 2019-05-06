
package RequestHandlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

type CRH struct{
	AddressPort string
}
func (h CRH) Send(msg [] byte) net.Conn {
	l, err := net.Dial("tcp", h.AddressPort)
	if err != nil {
		fmt.Print("erro ao tentar conectar-se")
		log.Fatal(err)
	}
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(len(msg)))
	_,errW1 := l.Write(bs)
	for errW1 != nil{
		_,errW1 = l.Write(msg)
	}
	_,errW2 := l.Write(msg);
	for errW2 != nil{
		_,errW2 = l.Write(msg)
	}
	return l
}
func (h CRH) Recieve(conn net.Conn) [] byte{
	len_bytes := make([]byte, 8)
	conn.Read(len_bytes)
	var length int;
	reader := bytes.NewReader(len_bytes)
	err := binary.Read(reader, binary.LittleEndian, &length)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	buffer := make([]byte,length)
	conn.Read(buffer);
	return buffer
}
