package recorder

import (
	genContainer "dockgen/gen-container"
	"dockgen/pkg/dockCopy"
	"fmt"
	"strings"
)

type Container struct {
	co   *genContainer.Container
	path string

	cmdLen int
}

func (c *Container) BeforeRecord(cmdList []string) {
	for _, cmd := range cmdList {
		cmd = strings.TrimSpace(cmd)
		cmd = fmt.Sprintf("%s\n", cmd)
		dockCopy.Copy(c.co.AttachContainer.Conn, strings.NewReader(cmd))
	}

	c.cmdLen = len(cmdList)
}

func (c *Container) AfterRecord(cmd []string) {
	return
}

func NewContainer(co *genContainer.Container, path string) *Container {
	return &Container{
		co:   co,
		path: path,
	}

}

func (c *Container) Record() ([]string, error) {
	err := c.co.Closed()
	if err != nil {
		return nil, err
	}
	history := NewHistoryFileRecorder(c.path)

	list, err := history.Record()

	if err != nil {
		return nil, err
	}

	return list[c.cmdLen:], nil
}
