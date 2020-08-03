package deployer

import (
	"github.com/go-gomail/gomail"
	"walle-go/conf"
	"walle-go/logger"
)

func SendMail(title string, body string, to []string) {
	var (
		c conf.Conf
	)
	c.GetConf()
	m := gomail.NewMessage()
	m.SetHeader("From", c.Mail.User)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(c.Mail.Host, c.Mail.Port, c.Mail.User, c.Mail.Password)
	logger.Infof("h: %v p: %v u: %v p: %v", c.Mail.Host, c.Mail.Port, c.Mail.User, c.Mail.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
