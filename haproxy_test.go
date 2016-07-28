package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

const (
	Source = "./test.txt"
	Backup = "./test.txt.bak"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func setUp() {
	if _, err := os.Stat("/config"); os.IsNotExist(err) {
		os.Mkdir("/config", 0777)
	}

	CopyFile("/config/production.json", "config/production.json")

	os.Create(Source)
}

func tearDown() {
	os.Remove(Source)
	os.Remove(Backup)
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestBackupConfigFile(t *testing.T) {

	_, err := backupConfigFile(Source, Backup)
	assert.Nil(t, err)

	_, err = os.Stat(Backup)
	assert.Nil(t, err)
}

func TestContentChanged(t *testing.T) {

	changed := contentChanged(Source, Backup)
	assert.False(t, changed, "changed should be False")

}
