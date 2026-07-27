package main

import (
	"fmt"
	"net/smtp"
	"time"
)

func EmailWorker(id int, ch chan Receipent) {

	smtphost := "localhost"
	smtpport := "1025"

	for r := range ch {
		// formattedStr := fmt.Sprintf(
		// 	"To: %s\r\nSubject: Testing\r\n\r\n%s\r\n", r.Email, "Hello")

		// msg:= []byte(formattedStr)

		msg, err := ExecuteTemplate(r)

		if err != nil {
			continue
		}

		err = smtp.SendMail(smtphost+":"+smtpport, nil, "khan@gmail.com", []string{r.Email}, []byte(msg))
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(50 * time.Millisecond)
	}
}
