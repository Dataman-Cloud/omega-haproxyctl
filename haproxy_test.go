package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBackupConfigFile(t *testing.T) {
	const (
		source = "./test.txt"
		backup = "./test.txt.bak"
	)

	os.Create(source)

	_, err := backupConfigFile(source, backup)
	assert.Nil(t, err)

	_, err = os.Stat(backup)
	assert.Nil(t, err)
}

func TestContentNotChanged(t *testing.T) {
	const (
                source = "./test.txt"
                backup = "./test.txt.bak"
        )
	
	notchanged := contentNotChanged(source, backup)
	assert.True(t, notchanged, "notchanged should be true")

	os.Remove(source)
	os.Remove(backup)
}
