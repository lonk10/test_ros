package ros_resources

import (
	"fmt"

	"github.com/abu-lang/goabu/memory"
	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
)

type Topic struct {
	name    string
	client  *rclgo.Client
	rostype types.ServiceTypeSupport
	args    []string
}

func NewTopic(node *rclgo.Node, name string, rostype types.ServiceTypeSupport) (ROSdelegate, memory.Resources, error) {
	c, err := node.NewClient(name, rostype, nil)
	if err != nil {
		return nil, memory.MakeResources(), fmt.Errorf("Could not create service: %v", err)
	}

	return Topic{name: name, client: c, rostype: rostype}, memory.MakeResources(), nil
}

func (s Topic) Modified(name string, mem memory.Resources, errs chan<- error) *memory.Resources {
	return nil
}
