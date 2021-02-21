package label

type SourceLabels map[string]string

type Source interface {
	Get() (SourceLabels, error)
}
