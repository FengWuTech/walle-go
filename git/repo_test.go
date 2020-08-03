package git

import (
	"fmt"
	"os"
	"testing"
	"walle-go/logger"
)

func TestMain(m *testing.M) {
	logger.Setup()
	exitVal := m.Run()
	os.Exit(exitVal)
}


func TestRepo_Init(t *testing.T) {
	dir := "/tmp/walle/codebase/test"//os.Getwd()
	r := Repo{Path: dir}
	r.Init("ssh://git@192.168.253.179:2222/golang/walle-go.git", "master")
}

func TestRepo_Commits(t *testing.T) {
	dir := "/tmp/walle/codebase/test"//os.Getwd()
	branch:="master"

	r := Repo{Path: dir,}
	branchList:=r.Commits(branch)
	fmt.Printf("%v",branchList)
}

func TestRepo_Branches(t *testing.T) {
	dir := "/tmp/walle/codebase/test"//os.Getwd()
	r := Repo{Path: dir,}
	branches := r.Branches()
	fmt.Printf("%v",branches)
}

func TestRepo_Tags(t *testing.T) {
	dir := "/tmp/walle/codebase/test"//os.Getwd()
	r := Repo{Path: dir,}
	branchList:=r.Tags()
	fmt.Printf("%v",branchList)
}

//切commit
func TestRepo_Checkout_2_Commit(t *testing.T) {
	dir := "/tmp/walle/codebase/test"//os.Getwd()
	r := Repo{Path: dir,}
	r.Checkout2Commit("master", "9df0bb689")
}
