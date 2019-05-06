package Shared


// Deve se ter um sentido de interfaces iguais utilizadas em server e client para que os dados recebidos e enviados sejam congruentes
/* A função invoke tanto em server(invoker) quanto em client(requestor) devem ter os mesmos tipos que chamam funções para que esses tipos mantenham os
estados(mensagem, usuario, porta, etc). Deve se ter do inicio ao fim o mesmo tipo que chama as funções afim de salvar esses estados
*/
/*Pode ser necessario o retorno do servidor de names ja retornar a conexão e entao, criar uma estrutura tipo que possua essa conexão dentro e passar adiante,
do client proxy até o request handler client, passar para o lado servidor o que for necessario.
*/

/*O ip do servidor deve ser posto na aplicação servidor e ser passado para o cadastro do servername, afim de que o client receba quando fizer a chamada.
O servidor não pode ter o seu ip setado dentro da parte do middleware(so relembrando)
*/

type AOR struct {
	IP string
	Port int
	OID int
}

type ChatMsg struct{
	Msg string
	User string
}

type FTMsg struct{
	File []byte
	File_name string
}
 type BindMessage struct {
 	Service_name string
 	AOR AOR
 }

type RequestFormat struct {
	Args interface{}
	OID int
	Function string
} 