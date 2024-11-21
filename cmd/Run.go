package main

import (
	"banco/cmd/actions/bank"
	"banco/cmd/actions/user"
	"fmt"
)

func RunApp() {

	var escolha int = 0

	for {
		fmt.Println("<---------------->")

		fmt.Println("Seja bem vindo ao nosso banco!")
		fmt.Println("Digite o seu tipo de usuario: [1]-usuário [2]-Bank")
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			user.UserMenu()

		case 2:
			bank.LoginBank()

		}

		fmt.Println("<---------------->")

	}

}
