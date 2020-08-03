package wsio

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"testing"
	"walle-go/deployer"
	"walle-go/logger"
	"walle-go/models"
)

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic("setup chdir error")
	}
	logger.Setup()
	models.Setup(false)
	exitVal := m.Run()
	os.Exit(exitVal)
}


func TestRepo_OnOpen(t *testing.T) {
	namspace := "fengwu"

	so := WebSocketIO{
		Namespace: namspace,
	}

	so.onOpen(nil, map[string]interface{}{"project_id":3})
}

func TestWebSocketIO_onBranches(t *testing.T) {
	d := deployer.NewProject(10)
	d.InitRepo()

	branches := d.Repo.Branches()
	fmt.Printf("%v",branches)
}