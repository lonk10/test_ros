package ros_resources

import (
	"github.com/abu-lang/goabu/memory"
)

type ROSdelegate interface {
	Modified(string, memory.Resources, chan<- error) *memory.Resources
}
