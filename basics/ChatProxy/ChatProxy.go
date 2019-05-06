package ChatProxy

import (
	"../Requestor"
	"../Shared"
)

type Chat struct {
	IP string
	Port int
	OID int
}

func (c Chat) Send(Msg string, User string) string{
	return Requestor.Invoke(Shared.AOR(c), "Chat", "Send", Shared.ChatMsg{Msg, User}).(string)
}
func (c Chat) Listen() [] string{
	return Requestor.Invoke(Shared.AOR(c), "Chat", "Listen", nil).([]string) //converter para o tipo desejado na saída
}
