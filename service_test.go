package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

var service Service

func aService() error {
	service = &Greeter{}
	_, ok := service.(*Greeter)
	if ok {
		return nil
	}
	return fmt.Errorf("Service not available")
}

func iStartIt() error {
	greeter := service.(*Greeter)
	if greeter.msg != "" {
		return fmt.Errorf("Service already started!")
	}
	service.Start()
	return nil
}

func itShouldDisplayMessage(arg1 string) error {
	greeter := service.(*Greeter)
	if greeter.msg != "Hello, World!" {
		return fmt.Errorf("Incorrect Message")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a service$`, aService)
	ctx.Step(`^I start it$`, iStartIt)
	ctx.Step(`^it should display "([^"]*)" message$`, itShouldDisplayMessage)
}
