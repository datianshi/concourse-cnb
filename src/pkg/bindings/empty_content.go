package bindings

import "io"

type EmptyContent struct {
}

func (e *EmptyContent) GenerateContent() (io.Reader, error) {
	return nil, nil
}

func (e *EmptyContent) Name() string {
	return ""
}
