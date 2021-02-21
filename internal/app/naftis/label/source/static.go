package source

import "gitlab.com/naftis/app/naftis/internal/app/naftis/label"

type Static struct {
	labels label.SourceLabels
}

func NewStatic(labels label.SourceLabels) *Static {
	return &Static{labels: labels}
}

func (s *Static) Get() (label.SourceLabels, error) {
	return s.labels, nil
}
