package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	var command string
	if runtime.GOOS == "windows" {
		command = "cmd.exe"
	} else {
		log.Fatalf("O sistema operacional %s não é suportado", runtime.GOOS)
	}

	cmd1 := exec.Command(command, "/C", "wmic nicconfig where TcpipNetbiosOptions=0 call SetTcpipNetbios 2")
	err := cmd1.Run()
	if err != nil {
		log.Fatalf("Erro ao executar o comando wmic: %s", err)
	} else {
		fmt.Print("Comando Efetuado com Sucesso")
	}
	fmt.Print("\n")
	fmt.Println("Script Finalizado")

	if runtime.GOOS == "windows" {
		fmt.Print("Deseja reiniciar o computador agora? (s/n): ")
		var input string
		fmt.Scanln(&input)
		if input == "s" {
			cmd4 := exec.Command(command, "/C", "shutdown", "/r", "/t", "0")
			err = cmd4.Run()
			if err != nil {
				log.Fatalf("Erro ao reiniciar o computador: %s", err)
			}
		}
	}
}
