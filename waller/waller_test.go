package waller

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"walle-go/logger"
)

func TestMain(m *testing.M) {
	logger.Setup()
	exitVal := m.Run()
	os.Exit(exitVal)
}


func TestWaller_Run(t *testing.T) {
	pwd, _  := os.Getwd()
	w := Waller{
		Env: []string{"SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.vW5C5k8U5d/Listeners",},
		Dir: pwd,
	}
	err := w.Run(map[string]string{}, "ssh-add -l")
	assert.Equal(t, err, nil)
}

func TestWaller_Put(t *testing.T) {
	pwd, _  := os.Getwd()
	w := Waller{
		Env:  []string{"SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.vW5C5k8U5d/Listeners",},
		Dir:  pwd,
		Host: "192.168.238.178",
		User: "root",
		Port: 22,
	}
	err := w.Put(map[string]string{}, "logs/walle-go.log", "/tmp")
	assert.Equal(t, err, nil)
}

func TestWaller_RemoteRun(t *testing.T) {
	pwd, _  := os.Getwd()
	w := Waller{
		Env:  []string{"SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.vW5C5k8U5d/Listeners",},
		Dir:  pwd,
		Host: "192.168.238.178",
		User: "root",
		Port: 22,
	}
	_, err := w.RemoteRun(map[string]string{}, "uptime")
	_, err = w.RemoteRun(map[string]string{}, "mkdir /tmp/aaa")
	assert.Equal(t, err, nil)
}