# **Generator Password**

Este é um projeto simples escrito em Go (Golang) que gera senhas aleatórias com base no tamanho especificado pelo usuário que fiz enquanto aprendo sobre a linguagem. O código garante que a senha gerada contenha pelo menos uma letra maiúscula, um número e um caractere especial, além de outros caracteres aleatórios para completar o comprimento desejado.


## **O que é?**
O projeto é um Gerador de Senhas que cria senhas seguras e aleatórias. Ele foi desenvolvido em Go, uma linguagem de programação conhecida por sua simplicidade e eficiência. O programa solicita ao usuário o comprimento da senha desejada e, em seguida, gera uma senha que inclui:

Letras minúsculas (a-z)

Letras maiúsculas (A-Z)

Números (0-9)

Caracteres especiais (!@#$%^&*()+?><:{}[])

A senha gerada é embaralhada para garantir que seja aleatória e segura.

## **Como funciona?**

O código é dividido em duas funções principais:

1. `main:` Esta função é o ponto de entrada do programa. Ela solicita ao usuário o comprimento da senha e chama a função `generatePassword` para gerar a senha. Em seguida, exibe a senha gerada.

2. `generatePassword:` Esta função é responsável por gerar a senha. Ela cria uma senha que inclui pelo menos uma letra maiúscula, um número e um caractere especial. O restante da senha é preenchido com caracteres aleatórios escolhidos de um conjunto que inclui letras minúsculas, letras maiúsculas, números e caracteres especiais. Por fim, a senha é embaralhada para garantir a aleatoriedade.

## **Como Rodar?**
Digitar `go run main.go` no seu terminal

## **Como usar?**
1. **Executar o programa:** Execute o programa conforme descrito na seção "Como rodar?".

2. **Inserir o comprimento da senha:** Quando solicitado, insira o número de caracteres que você deseja que a senha tenha. Por exemplo, se você quiser uma senha com 12 caracteres, digite 12 e pressione Enter.

3. **Obter a senha:** O programa gerará uma senha aleatória e a exibirá no terminal.

## **Possiveis Melhorias:**
- Permitir que o usuário escolha quais tipos de caracteres deseja incluir na senha (por exemplo, apenas letras e números).

- Gerar múltiplas senhas de uma vez.

- Salvar as senhas geradas em um arquivo.
