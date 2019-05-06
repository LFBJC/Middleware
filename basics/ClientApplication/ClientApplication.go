package main

import(
	"../ChatProxy"
	"../NamingProxy"
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Bem vindo!");
	fmt.Println("Digite o seu nome");
	user_name, err := reader.ReadString('\n');
	if err!=nil{
		fmt.Println("erro ao tentar ler a entrada");
	}
	fmt.Println("este é um chat com transferência de arquivos");
	fmt.Println("para enviar mensagens basta digitá-las e apertar enter");
	fmt.Println("escreva \":f \" + nome do arquivo para enviar e \":d \" + nome do arquivo para baixar e em seguida aperte enter");
	fmt.Println("escreva \":l \" e aperte enter para ver os arquivos disponiveis");
	nameServer := NamingProxy.NS{"127.0.0.1:",1025,0};
	chat := ChatProxy.Chat(nameServer.Lookup("Chat"));
	for {
		//é preciso definir que tipo é esse chat e colocar como tipo que chama nas funcões Listen e Send
		msgs := chat.Listen();
		for i := 0; i < len(msgs); i++ {
			fmt.Println(msgs[i]);
		}
		msg, err := reader.ReadString('\n');
		if err != nil {
			fmt.Println("erro ao tentar ler a entrada");
		}
		resp := chat.Send(msg, user_name);
		if (msg[:3] == ":d ") { //download de arquivo
			b := []byte(resp)
			f, err := os.Create("Download"+msg[(strings.LastIndex(msg,"/")):]);
			if err != nil {
				fmt.Print("Você pediu um Download: erro ao criar o arquivo")
			}
			_, err2 := f.Write(b);
			if err2!= nil {
				fmt.Print("Você pediu um Download: erro ao tentar escrever no arquivo")
			}
		} else {
			fmt.Println(resp);
		}
	}
}