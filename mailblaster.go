package main

import (
	"fmt"
	"net/mail"

	"github.com/ppreeper/email"
	"github.com/ppreeper/stringpad"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(stringpad.RJustLen("Hello", 20))
	for i := 0; i < 1000; i++ {
		//create message
		m := email.Message{
			From:     mail.Address{Name: "Mailer", Address: "mailer@example.com"},
			To:       []mail.Address{{Name: "Standard User", Address: "standard.user@example.com"}},
			Subject:  "test subject",
			Body:     Message(),
			MimeType: "text/html",
		}

		fmt.Println(string(m.BuildMessage()))

		// Connect to the remote SMTP server.
		// user := email.User{Username: "mailer@example.com", Password: "pa55w0rd"}
		// server := email.SMTPServer{Host: "mail.example.com", Port: "25", STARTTLS: false}

		// err := server.Send(&user, &m)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
}

// Message
func Message() string {
	return Greeting() + Body() + Salutation()
}

// Greeting
func Greeting() string {
	return "Greeting"
}

// Body
func Body() string {
	return "body"
}

// Salutation
func Salutation() string {
	return "salutation"
}
