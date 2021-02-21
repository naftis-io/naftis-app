package label

import "gitlab.com/naftis/app/naftis/pkg/protocol/entity"

type Container struct {
	labels  map[string]entity.NodeLabel
	sources []Source
}

func NewContainer() *Container {
	return &Container{
		labels:  make(map[string]entity.NodeLabel, 0),
		sources: make([]Source, 0),
	}
}

func (c *Container) Refresh() error {
	for _, source := range c.sources {
		labels, err := source.Get()
		if err != nil {
			return err
		}

		for key, value := range labels {
			c.labels[key] = entity.NodeLabel{
				Key:   key,
				Value: value,
			}
		}
	}

	return nil
}

func (c *Container) AttachSource(source Source) {
	c.sources = append(c.sources, source)
}

func (c *Container) Map() map[string]entity.NodeLabel {
	return c.labels
}

func (c *Container) List() []entity.NodeLabel {
	result := make([]entity.NodeLabel, 0)

	for _, label := range c.labels {
		result = append(result, label)
	}

	return result
}
