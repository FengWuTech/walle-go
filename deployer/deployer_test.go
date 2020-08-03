package deployer

import (
	"os"
	"path"
	"runtime"
	"testing"
	"walle-go/logger"
	"walle-go/models"
)

func TestMain(m *testing.M) {
	logger.Setup()
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic("setup chdir error")
	}
	models.Setup(true)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestDeployer_WalleDeploy(t *testing.T) {
	d := New(122)
	d.WalleDeploy()
}