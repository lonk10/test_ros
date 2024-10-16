package ros_resources

import (
	"aburos"

	"github.com/abu-lang/goabu/memory"
	"github.com/tiiuae/rclgo/pkg/rclgo"
)

type ioResourceMeta struct {
	resourceType string
	isInput      bool
	isOutput     bool
}

type resource struct {
	meta ioResourceMeta
	ROSdelegate
	managed []string
}

type ROSresources struct {
	memory.Resources
	inputs  chan string
	errors  chan error
	node    *rclgo.Node
	pubs    map[string]rclgo.Publisher
	subs    map[string]rclgo.Subscription
	clients map[string]rclgo.Client
	//delegates []*resource
	rosKeywords []string
	vecType     aburos.VehicleType
	//frames    map[string]frame
}

func (i *ROSresources) Modified(r string) {
	if !i.Has(r) {
		return
	}
	present := i.isKeyword(r)
	if !present {
		return
	}
	/*
		mods := resource.Modified(r, i.Resources.Extract(resource.managed), i.errors)
		if mods != nil {
			i.Resources.Enclose(mods.Extract(resource.managed))
		}*/
	switch r {
	case "start":
		i.vecType.Start()
	case "init":
		i.vecType.Init()
	case "mode":
		i.vecType.SetMode(i.Resources.Text["mode"])
	case "move":
		x := i.Resources.Float["move_x"]
		y := i.Resources.Float["move_y"]
		z := i.Resources.Float["move_z"]
		i.vecType.Move(x, y, z)
	case "position":
		lat := i.Resources.Float["position_lat"]
		lon := i.Resources.Float["position_lon"]
		alt := i.Resources.Float["position_alt"]
		i.vecType.SetPosition(lat, lon, alt)
	}
}

func (i *ROSresources) isKeyword(name string) bool {
	for _, r := range i.rosKeywords {
		if r == name {
			return true
		}
	}
	return false
}
