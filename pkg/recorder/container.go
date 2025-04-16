package recorder

import (
	genContainer "dockgen/gen-container"
)

type Container struct {
	co   *genContainer.Container
	path string
}

func NewContainer(co *genContainer.Container, path string) *Container {
	return &Container{
		co:   co,
		path: path,
	}

}

func (c Container) Record() ([]string, error) {
	err := c.co.Closed()
	if err != nil {
		return nil, err
	}
	history := NewHistoryFileRecorder(c.path)
	return history.Record()
}
