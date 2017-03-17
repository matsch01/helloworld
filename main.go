package main

import (
	"fmt"

	"github.com/levigross/grequests"
)

func sayAnotherThing() {
	fmt.Println("say another thing")
}

func requestSomething() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		fmt.Println("Unable to make request: ", err)
	}
	fmt.Println(resp.String())
}

func sayAThing() {
	fmt.Println("say a thing")
}

func sayHello() {
	fmt.Println("hello world")
}

func main() {
	sayAThing()
	sayHello()
	requestSomething()
}
