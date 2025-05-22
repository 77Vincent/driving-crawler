package main

import (
	"log"
	"net/smtp"
)

const (
	src   = "wojiaotianqi@gmail.com"
	dest  = "wojiaotianqi@gmail.com"
	dest2 = "tianqi.wen@fastretailing.com"
	dest3 = "yenan.rong@fastretailing.com"
	dest4 = "yenanrong0426@gmail.com"
)

func notify() {
	// Email details
	from := src                    // Replace with your email
	password := "ayvhkjtazrfxuebd" // app password
	to := []string{dest, dest2, dest3, dest4}

	// SMTP server configuration (Gmail example)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	message := []byte("To: " + from + "\r\n" +
		"Subject: 赶紧预约，已找到空缺！\r\n" +
		"\r\n" +
		"https://www.keishicho-gto.metro.tokyo.lg.jp/keishicho-u/reserve/offerList_detail?tempSeq=363&accessFrom=offerList\r\n")

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send Email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
