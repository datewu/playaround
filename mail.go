package main

// Subject: <your-subject>
// Content-Type: multipart/mixed; boundary="-unique-str"
//
// ---unique-str
// Content-Type: text/html
// Content-Disposition: inline
//
// <html-body here>
// ---unique-str
// Content-Type: application; name=<attachment-mime>
// Content-Transfer-Encoding: base64
// Content-Disposition: attachment; filename=<attachment-name>
//
// <your base64-encoded attachment here>
// ---unique-str--

import (
	"log"
	"net/smtp"
)

func main() {
	to := []string{"wutuofu@qq.com"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: noreply new discount Gophers!\r\n" +
		"\r\n" +
		"nonono This is the email body again.\r\n")

	err := smtp.SendMail("localhost:25", nil, "noreply@wutuofu.com", to, msg)
	log.Println(err)
}

// [root@wutuofu ~]# cat l
// Subject: test again
// Content-Type: multipart/mixed; boundary="-unique-str"
//
// ---unique-str
// Content-Type: text/html
// Content-Disposition: inline
//
// <p>heelo world</p>
// <hr>
// <p>go world</p>
// ---unique-str
// Content-Type: application; name=<attachment-mime>
// Content-Transfer-Encoding: base64
// Content-Disposition: attachment; filename=<attachment-name>
//
// R0lGODlhHAAUAPMFACE5CDlaCHuMa+fn50J7ANaM//+MjP/WjIz///9r9/9rawAAAAAAAAAAAAAAAAAAACH/C05FVFNDQVBFMi4wAwEAAAAh+QQJBQAFACwAAAAAHAAUAAADhFi6PAIQCkZre0GQJ6b9hQMIR3kY2gBWIkAcXReMK8VxGQapddidpJxGQ+CtNKbSJqb0gETJQ2ZZQhkt0KSQJM0cM9EAYcwdfrkmzKtkBmkMUcNMK7hWMFF20voFwPNJbRYyhGh5NFgzhoBpAHYFQ4ySXhcjkowPV3iXjJQKGDGhoqOCCQAh+QQJBQAEACwAAAEAGgATAAAEeJDIOaqoM+ttgxfXsI3SIHhIeoIimZmAkM4IELQuAa+158U4EkgW+5xsgODoRFvxWC5YMwCiIk7KF2halV0vJJN3RmXOwCMxrXY8oznms20MFQLW7ObQju87szx9eH8ZQ0iCelQ4Jl+BiFdAEnGPfWUETkOZmpsfEQAh+QQJBQAGACwAAAIAFwASAAADcmi6074wwiZqbTIPEYLtwpBNhCcUqBWOxhagXqEG4sidMgB2gDC6KFnBdKNpOEGYKueLVADJ4VJWcUKjRJw0MnhFi1oPVxvklLRVSCUK62Ct3iQSW2txrmw2p7Yx6+J5OSJda089ZFghLh48OjoqMwADCQAh+QQJBQAEACwBAAQAFwAQAAAEcpDISeW4uGo6hAhA6A1b1QFCoRafQJbEmcz0HAiwnNw7KL6aT8H2uRVdmwFvNmwVbxpdzenDVVq1GTUVAMY+WZrxVgB4neEeKFEGKgEJQbocYncnwrg87NEiPQFDNiAqaShfAIJMKlhZh28ekZE+RpM4EQAh+QQFBQAEACwBAAEAGgATAAAEeZDISSsZeNjNb+4gFgijIGgghZFAoLABmg4sKSi4LXO0iwO3gLC12/RwSFutqCIhkybYiWd65koBA1DAcVpfWFPL2PqCxb/ABmYeAtIWr89ag6vm1W8ZDLjPFX9tUxIDe0hDgYeDNDdmL41WIxcsjpU/GoWImpucJxEAOw==
// ---unique-str--
