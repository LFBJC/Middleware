package NamingProxy

import (
	"../Requestor"
	"../Shared"
)

type NS struct {
	IP string
	Port int
	OID int
}

func (o NS) Bind(service_name string, aor Shared.AOR) string {
	return Requestor.Invoke(Shared.AOR(o),  "Naming", "Bind", Shared.BindMessage{service_name,aor}).(string) // retorna o status da operação
}
func (o NS) Lookup(service_name string) Shared.AOR{
	return Requestor.Invoke(Shared.AOR(o),  "Naming", "Lookup", service_name).(Shared.AOR);
}
func (o NS) List() []string {
	return Requestor.Invoke(Shared.AOR(o), "Naming", "List", nil).([] string);
}
