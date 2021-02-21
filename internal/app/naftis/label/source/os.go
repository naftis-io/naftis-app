package source

import (
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label"
	"runtime"
)

type Os struct {
}

func NewOs() *Os {
	return &Os{}
}

func (o *Os) Get() (label.SourceLabels, error) {
	result := make(label.SourceLabels, 0)

	result[label.Architecture] = runtime.GOARCH
	result[label.OperatingSystem] = runtime.GOOS

	return result, nil
}
