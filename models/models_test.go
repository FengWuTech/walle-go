package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"path"
	"runtime"
	"testing"
	"walle-go/logger"
)

func TestSetup(t *testing.T) {
	logger.Setup()
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic("setup chdir error")
	}
	Setup(true)
	var project Projects
	db.First(&project)
	logger.Debug(*project.Id)
}