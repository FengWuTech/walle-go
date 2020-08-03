package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestWatcher(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("getwd failed: %s", err.Error())
	}
	tmpDir := filepath.Join(wd, "_gdp_test_watcher")

	t.Log("test TempDir:", tmpDir)

	defer func() {
		if Exists(tmpDir) {
			err := os.RemoveAll(tmpDir)
			t.Logf("RemoveAll test_tmp_dir=%s, err=%v", tmpDir, err)
		}
	}()

	if !Exists(tmpDir) {
		err := os.MkdirAll(tmpDir, 0777)
		if err != nil {
			t.Log("create tmpDir failed:", tmpDir, err)
		}
	}

	w := NewWatcher(tmpDir, 100)

	ca := make(chan int)
	var onceA sync.Once

	w.RegisterFileWatcher("*", func(e *WatcherEvent) error {
		fmt.Println("event a:", e)
		onceA.Do(func() {
			ca <- 7
		})
		return nil
	})

	var onceB sync.Once
	w.RegisterFileWatcher("b*", func(e *WatcherEvent) error {
		fmt.Println("event b:", e)

		onceB.Do(func() {
			ca <- 5
		})

		return nil
	})
}
