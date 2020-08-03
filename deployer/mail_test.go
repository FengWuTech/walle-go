package deployer

import "testing"

func TestSendMail(t *testing.T) {
	d := New(105)
	title := "上线部署成功"
	body := d.toMailMessage(title)
	SendMail(title, body, []string{"yangyang@aiforward.com",})
}