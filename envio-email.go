package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func main() {

	var mensagem string
	var destinatario string
	var tipoEnvio int
	var assunto string

	fmt.Println("Informe o assunto do e-mail")
	fmt.Scanln(&assunto)

	fmt.Println("Digite a mensagem do e-mail")
	fmt.Scanln(&mensagem)

	fmt.Println("Informe o destinatário")
	fmt.Scanln(&destinatario)

	fmt.Println("Escolha o tipo de envio (1) Síncrono (2) Assíncrono")
	fmt.Scanln(&tipoEnvio)

	switch tipoEnvio {
	case 1:
		enviaEmail(assunto, destinatario, mensagem)
	case 2:
		go enviaEmail(assunto, destinatario, mensagem)
		fmt.Println("Envio de email agendado")
		fmt.Scanln(&mensagem)
	default:
		fmt.Println("Tipo de envio desconhecido")
		os.Exit(-1)
	}

}

func enviaEmail(assunto string, destinatario string, mensagem string) {

	remetente := ".....@gmail.com"

	auth := smtp.PlainAuth("",
		".....@gmail.com",
		"##############",
		"smtp.gmail.com",
	)

	msgBody := "From: " + remetente + "\n" +
		"To: " + destinatario + "\n" +
		"Subject:" + assunto + "\n\n" +
		mensagem + " - Enviado via GoLang"

	err := smtp.SendMail("smtp.gmail.com:587",
		auth,
		remetente, []string{destinatario}, []byte(msgBody))

	if err != nil {
		log.Printf("Erro ao enviar email: %s", err)
		return
	}

	fmt.Printf("email enviado com sucesso")

}
