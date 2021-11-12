package bindings

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Content
type Content interface {
	GenerateContent() (io.Reader, error)
	Name() string
}

type Binding struct {
	name        string
	bindingType string
	output      string
	content     Content
}

func NewBinding(name, bindingType, output string, content Content) *Binding {
	return &Binding{
		name:        name,
		bindingType: bindingType,
		output:      output,
		content:     content,
	}
}
func (b *Binding) CreateBinding() error {
	var err error
	var typeFile, contentFile *os.File
	if _, err = os.Stat(b.output); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(b.output, os.ModePerm); err != nil {
			return err
		}
	}

	//make binding name folder
	bindingDir := fmt.Sprintf("%s/%s", b.output, b.name)
	if _, err = os.Stat(bindingDir); errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(bindingDir, os.ModePerm); err != nil {
			return err
		}
	}

	//Write binding type
	if typeFile, err = os.OpenFile(fmt.Sprintf("%s/%s", bindingDir, "type"), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	}
	defer typeFile.Close()
	if _, err = typeFile.WriteString(b.bindingType); err != nil {
		return err
	}

	//Write binding content

	if contentFile, err = os.OpenFile(fmt.Sprintf("%s/%s", bindingDir, b.content.Name()), os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return err
	}
	defer contentFile.Close()

	var reader io.Reader
	var bindingContent []byte
	if reader, err = b.content.GenerateContent(); err != nil {
		return err
	}
	if bindingContent, err = ioutil.ReadAll(reader); err != nil {
		return err
	}
	if _, err = contentFile.Write(bindingContent); err != nil {
		return err
	}
	return nil

}
