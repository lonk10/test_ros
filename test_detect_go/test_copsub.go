package main

import (
	"fmt"
	"os"
	cop "test/test_detect_go/copter"
	sub "test/test_detect_go/sub"
	"time"
)

func testSub() error {
	subV, err := sub.NewSubVehicle("sub1", "aburos")
	if err != nil {
		return fmt.Errorf("could not create sub: %v", err)
	}
	fmt.Println("Initializing...")
	err = subV.Init()
	if err != nil {
		return fmt.Errorf("could not initiliaze sub: %v", err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Arming motors...")
	subV.ArmMotors()
	time.Sleep(1 * time.Second)
	fmt.Println("Setting mode to guided...")
	subV.SetMode("guided")
	time.Sleep(1 * time.Second)
	fmt.Println("Setting position...")
	subV.SetPoint(-5.0, 0, -4.0)
	time.Sleep(10 * time.Second)
	fmt.Printf("subV.LocalAgent.GetPosition(): %v\n", subV.LocalAgent.GetPosition())
	//submarine works!!
	return nil
}

func testCopter() error {
	copter, err := cop.NewCopterVehicle("/copter1", "")
	if err != nil {
		return fmt.Errorf("could not create copter: %v", err)
	}
	fmt.Println("Initializing...")
	err = copter.Init()
	if err != nil {
		return fmt.Errorf("could not initiliaze copter: %v", err)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Arming motors...")
	copter.ArmMotors()
	time.Sleep(1 * time.Second)
	fmt.Println("Setting mode to guided...")
	copter.SetMode("guided")
	time.Sleep(1 * time.Second)
	fmt.Println("Taking off...")
	copter.TakeOff(10)
	time.Sleep(2 * time.Second)
	//submarine works!!
	return nil
}

func main() {
	if err := testSub(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	/*
		if err := testCopter(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}*/
}
