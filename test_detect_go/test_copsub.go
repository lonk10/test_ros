package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	cop "test/test_detect_go/copter"
	sub "test/test_detect_go/sub"
	"time"
)

func testSub() error {
	subV, err := sub.NewSubVehicle("sub1", "/aburos")
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
	time.Sleep(1 * time.Second)
	fmt.Printf("subV.LocalAgent.GetPosition(): %v\n", subV.LocalAgent.GetPosition())
	//submarine works!!
	return nil
}

func testCopter() error {
	copter, err := cop.NewCopterVehicle("copter", "/aburos")
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
	copter.TakeOff(100.0)
	time.Sleep(2 * time.Second)
	//submarine works!!
	return nil
}

func main() {
	go testSub()
	go testCopter()
	sigChan := make(chan os.Signal, 1)

	// Notify the channel on receiving Interrupt (Ctrl+C) or SIGTERM signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Running program. Press Ctrl+C to exit...")

	// Block until a signal is received
	sig := <-sigChan
	fmt.Println()
	fmt.Printf("Received signal: %s, exiting...\n", sig)

	// Exit the program
	os.Exit(0)
	/*
		if err := testCopter(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}*/
}
