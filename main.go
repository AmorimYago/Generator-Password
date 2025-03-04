package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	var passwordLength int

	fmt.Println("Gerador de Senhas")
	fmt.Println("---------------------")
	fmt.Println("Escolha o tamanho da senha:")
	fmt.Scanln(&passwordLength)

	password := generatePassword(passwordLength)
	fmt.Println(password)
}


func generatePassword(length int) string {
	// Caracteres da senha
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()+?><:{}[]"
	allChars := lowerCase + upperCase + numbers + special

	// Add os caractéres obrigatórios (Letra maiuscula, numeros e caracter especial)
	mandatory := []byte{
		upperCase[rand.Intn(len(upperCase))],
		numbers[rand.Intn(len(numbers))],
		special[rand.Intn(len(special))],
	}

	// Cria a senha
	password := make([]byte, length-len(mandatory))

	// Insere os caracteres obrigatórios
	for i := range password {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	// Adiciona os caracteres obrigatórios
	password = append(password, mandatory...)

	// Embaralha a senha
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})


	return string(password)
}
