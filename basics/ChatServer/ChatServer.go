package ChatServer

import (
	"../FTProxy"
	"../NamingProxy"
	. "strings"
)
type Chat struct {
	Msgs [] string
}

func (c Chat) Send(Msg string, User string) string{
	if HasPrefix(Msg,":f ") {
		var naming_server = NamingProxy.NS{IP: "127.0.0.1", Port: 1024, OID: 0}
		ft := naming_server.Lookup("File Transfer")
		return FTProxy.FT(ft).Send(Msg[3:(LastIndex(Msg,"|"))],[]byte(Msg[(LastIndex(Msg,"|")):]))
	}else if HasPrefix(Msg,":d ") {
		var naming_server = NamingProxy.NS{IP: "127.0.0.1", Port: 1024, OID: 0}
		ft := naming_server.Lookup("File Transfer")
		return string(FTProxy.FT(ft).Download(Msg[3:]))
	}else if HasPrefix(Msg,":l ") {
		var naming_server = NamingProxy.NS{IP: "127.0.0.1", Port: 1024, OID: 0}
		ft := naming_server.Lookup("File Transfer")
		return Join(FTProxy.FT(ft).List(), "|") //transforma a lista de nomes de arquivos numa única String
	} else {
		c.Msgs[len(c.Msgs)] = Msg + ":" + User
		return "YOU SENT A MESSAGE"//status da operação com formato do tipo da função
	}
}
func (c Chat) Listen() [] string {
	return c.Msgs
}
