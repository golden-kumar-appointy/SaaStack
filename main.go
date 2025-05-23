package main

import (
	"fmt"
	"saastack/demo"
	"time"
)

const SLEEP = time.Second * 2

func main() {
	fmt.Println("Demo1")
	response := demo.SendEmailViaAWSSES()
	fmt.Println("SendEmailViaAWSSES:\n", response)

	time.Sleep(SLEEP)

	fmt.Println("\nDemo2")
	response = demo.SendEmailAndNotification()
	fmt.Println("SendEmailAndNotification:\n", response)

	time.Sleep(SLEEP)

	fmt.Println("\nDemo3")
	response = demo.UnimplementedInterfaceHandler()
	fmt.Println("UnimplementedInterfaceHandler:\n", response)

	time.Sleep(SLEEP)

	fmt.Println("\nDemo4")
	response = demo.UnimplementedInterfacePlugin()
	fmt.Println("UnimplementedInterfacePlugin:\n", response)
}
