package infra

import (
	"io"
	"io/ioutil"
	"os"
)

type (
	// FileSystem implements domain.FileSystem .
	FileSystem struct{}
)

// Exist implements domain.FileSystem#Exist .
func (fsys FileSystem) Exist(dst string) bool {
	_, err := os.Stat(dst)
	return err == nil
}

// MkdirAll implements domain.FileSystem#MkdirAll .
func (fsys FileSystem) MkdirAll(dst string) error {
	return os.MkdirAll(dst, 0755)
}

// Write implements domain.FileSystem#Write .
func (fsys FileSystem) Write(dst string, data []byte) error {
	return ioutil.WriteFile(dst, data, 0644)
}

// GetWriteCloser implements domain.FileSystem#GetWriteCloser .
func (fsys FileSystem) GetWriteCloser(dst string) (io.WriteCloser, error) {
	return os.Create(dst)
}

// Getwd implements domain.FileSystem#Getwd .
func (fsys FileSystem) Getwd() (string, error) {
	return os.Getwd()
}
