package FTServer

import "reflect"

type FileTransfer struct {
	Files map[string] [] byte
}
func (ft FileTransfer) Send(Name string, File [] byte) string {
	ft.Files[Name]=File
	return "YOU SENT A FILE"
}
func (ft FileTransfer) Download(nome string) [] byte{
	return ft.Files[nome]
}
func (ft FileTransfer) List() []string{
	keys := reflect.ValueOf(ft.Files).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
