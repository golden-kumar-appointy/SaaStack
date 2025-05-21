package email

import "fmt"

type service1 interface {
	service1func(int) error
}

type Service1Client struct{}

type UnImplementedService1Client struct{}

var UserPreference = ""

func NewService1Client(userPreference string) service1 {
	switch userPreference {
	case "service1":
		return &Service1Client{}
	default:
		return &UnImplementedService1Client{}
	}
}

func (s Service1Client) service1func(integrt int) error {
	fmt.Println("Default: Not implemented")
	return nil
}

func (s UnImplementedService1Client) service1func(integrt int) error {
	fmt.Println("Default: Not implemented")
	return nil
}

// write a struct, interface,
// how can we declare methods
