package Invoker

import (
	"../ChatServer"
	"../FTServer"
	"../NameService"
	"../RequestHandlers"
	"../Shared"
	"encoding/json"
)

type Invoker struct {
	IP string
	Port int
}
func (i Invoker)Invoke(){
	for true {
		srh := RequestHandlers.SRH{i.IP,i.Port}
		b, c:= srh.Recieve()
		var message Shared.RequestFormat;
		json.Unmarshal(b,&message)
		var result interface{}; //guarda o resultado da invocação qualquer que seja
		switch message.OID {// id do objeto definido a priori no nosso caso [0 = Servidor de nomes, 1 = Chat, 2 = transferencia de arquivos]
		case 0:
			naming := NameService.Naming{}
			switch message.Function {
			case "Bind":
				args := message.Args.(Shared.BindMessage)
				result = naming.Bind(args.Service_name, args.AOR)
				break
			case "Lookup":
				result = naming.Lookup(message.Args.(string))
			default:
				result = naming.List()
				break
			}
			break
		case 1:
			chat := ChatServer.Chat{make([]string,0)}
			switch message.Function {
			case "Send":
				args := message.Args.(Shared.ChatMsg)
				result = chat.Send(args.Msg,args.User)
				break
			default:
				result = chat.Listen()
				break
			}
			break
		default:
			ft := FTServer.FileTransfer{map[string][]byte{}}
			switch message.Function {
			case "Send":
				args := message.Args.(Shared.FTMsg)
				result = ft.Send(args.File_name,args.File)
				break
			case "Download":
				result = ft.Download(message.Args.(string))
				break
			default:
				result = ft.List()
				break
			}
			break
		}
		srh.Send(result.([]byte), c)//envia resposta ao cliente
	}
}