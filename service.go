package main

import (
	"fmt"
)

type Service interface {
	Start()
}

type Greeter struct {
	msg string
}

func (g *Greeter) Start() {
	g.msg = "Hello, World!"
}

func main() {
	var service Service = &Greeter{}

	service.Start()

	greeter := service.(*Greeter)
	fmt.Println(greeter.msg)
}
