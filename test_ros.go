package main

import (
	geom "test/msgs/geometry_msgs/msg"
	stdmsg "test/msgs/std_msgs/msg"
)

// geometry_msgs/PoseStamped equivalent in Go
type PoseStamped struct {
	Header stdmsg.Header `ros2:"std_msgs/Header"`
	Pose   geom.Pose     `ros2:"geometry_msgs/Pose"`
}

func main() {
}
