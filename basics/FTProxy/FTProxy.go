package FTProxy

import (
	"../Requestor"
	"../Shared"
)

type FT struct {
	IP string
	Port int
	OID int
}

func (ft FT) Send(Name string, File [] byte) string {
	return (Requestor.Invoke(Shared.AOR(ft), "FileTransfer", "Send", Shared.FTMsg{File,Name})).(string) //retorna string
}
func (ft FT) Download(nome string) []byte {
	return (Requestor.Invoke(Shared.AOR(ft), "FileTransfer", "Download", nome)).([] byte) //converter para o tipo de retorno
}
func (ft FT) List() [] string {
	return Requestor.Invoke(Shared.AOR(ft), "FileTransfer", "List", nil).([] string) //converter para o tipo de retorno
}
