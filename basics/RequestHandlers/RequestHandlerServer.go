package RequestHandlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type SRH struct {
	IP string
	Port int
}
func (h SRH) Recieve() ([] byte, net.Conn) {
	l, err := net.Listen("tcp",h.IP+strconv.Itoa(h.Port))
	if err !=nil {
		fmt.Print("erro ao tentar escutar a porta "+strconv.Itoa(h.Port)+", "+string(err.Error()))
	}
	conn, err2 := l.Accept()
	if err2 !=nil {
		fmt.Print("erro ao tentar ao receber conexão do cliente")
	}
	len_bytes := make([]byte, 8)
	conn.Read(len_bytes)
	var length int;
	reader := bytes.NewReader(len_bytes)
	err3 := binary.Read(reader, binary.LittleEndian, &length)
	if err3 != nil {
		fmt.Println("binary.Read failed:", err3)
	}
	buffer := make([]byte,length)
	conn.Read(buffer);
	return buffer, conn;
}
func (h SRH) Send(msg []byte, conn net.Conn) {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(len(msg)))
	_,errW1 := conn.Write(bs)
	for errW1 != nil{
		_,errW1 = conn.Write(msg)
	}
	_,errW2 := conn.Write(msg);
	for errW2 != nil{
		_,errW2 = conn.Write(msg)
	}
}

