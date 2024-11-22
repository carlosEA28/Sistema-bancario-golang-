package user

import (
	"fmt"
	"log"
)

type User struct {
	Name     string
	Address  string
	Password string
	Balance  float64
}

func (u User) Depositar()             {}
func (u User) Sacar()                 {}
func (u User) VerSaldo()              {}
func (u User) Transferir(qtd float64) {}

var contas []User

func UserMenu() {
	for {
		fmt.Println("Digite a ação que deseja fazer: ")
		fmt.Println("[1]- Criar Conta")
		fmt.Println("[2]- Login")
		fmt.Println("[3]- Fechar")

		var escolha int
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			CriarConta()
		case 2:
			loginUser()
		case 3:
			fmt.Println("Obrigado pela preferência!")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

func LogedUserMenu() {
	fmt.Println("OI")
}

func CriarConta() {
	var nome, address, password string

	fmt.Println("<---------------->")
	fmt.Println("Criar a sua conta: ")

	fmt.Println("Digite seu nome:")
	fmt.Scan(&nome)

	fmt.Println("Digite seu endereço:")
	fmt.Scan(&address)

	fmt.Println("Digite sua senha:")
	fmt.Scan(&password)

	novoUser := User{
		Name:     nome,
		Address:  address,
		Password: password,
	}

	contas = append(contas, novoUser)

	fmt.Println("Conta criada com sucesso! Seu ID é:", len(contas)-1)
	fmt.Println("<---------------->")
}

func loginUser() {

	var idBusca int
	var nome, password string

	fmt.Println("Digite seu id: ")
	fmt.Scan(&idBusca)

	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Digite sua senha: ")
	fmt.Scan(&password)

	for id := range contas {

		if contas[id].Name == nome && contas[id].Password == password {
			fmt.Println("Login feito com sucesso")
			LogedUserMenu()
		} else {
			log.Fatal("Nome ou senha inválidos!")
		}

	}

}
