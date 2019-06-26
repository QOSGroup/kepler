package mail

import (
	"github.com/QOSGroup/kepler/server/config"
	"net/smtp"
	"strings"
)

func Send(address string, sub string, msg string) error {
	conf := config.DefaultMailConfig()
	auth := smtp.PlainAuth(
		conf.Identity,
		conf.Sender,
		conf.Password,
		conf.Host,
	)
	content := strings.Replace("From: "+conf.Sender+"~To: "+address+"~Subject: "+sub+"~~", "~", "\r\n", -1) + msg
	err := smtp.SendMail(
		conf.Host+":"+conf.Port,
		auth,
		conf.Sender,
		[]string{address},
		[]byte(content),
	)
	return err
}
