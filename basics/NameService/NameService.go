package NameService

import (
	"../Shared"
	"reflect"
)

type Naming struct {
	Lookup_table map[string] Shared.AOR
}
func (o Naming) Bind(service_name string, aor Shared.AOR) string {
	o.Lookup_table[service_name] = aor;
	return "NAME REGISTERED"
}
func (o Naming) Lookup(service_name string) Shared.AOR{
	return o.Lookup_table[service_name]
}
func (o Naming) List() []string {
	keys := reflect.ValueOf(o.Lookup_table).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}
