package ros_resources

import (
	"fmt"

	"github.com/abu-lang/goabu/memory"
	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
)

type Service struct {
	name    string
	client  *rclgo.Client
	rostype types.ServiceTypeSupport
	args    []string
}

func NewService(node *rclgo.Node, name string, rostype types.ServiceTypeSupport) (ROSdelegate, memory.Resources, error) {
	c, err := node.NewClient(name, rostype, nil)
	if err != nil {
		return nil, memory.MakeResources(), fmt.Errorf("Could not create service: %v", err)
	}

	return Service{name: name, client: c, rostype: rostype}, memory.MakeResources(), nil
}

func (s Service) Modified(name string, mem memory.Resources, errs chan<- error) *memory.Resources {
	return nil
}
