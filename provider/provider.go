package provider

import (
	"io/ioutil"
)

type InputProvider interface {
	HasNext() bool
	GetNext() (string, error)
}

type FileInputProvider struct {
	files []string
	currentPtr int
}

func NewFileInputProvider(files []string) InputProvider {
	return &FileInputProvider {
		files: files,
		currentPtr: 0,
	}
}

func (p *FileInputProvider) HasNext() bool {
	return p.currentPtr < len(p.files)
}

func (p *FileInputProvider) GetNext() (string, error) {
	file := p.files[p.currentPtr]
	p.currentPtr += 1
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}