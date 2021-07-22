package register

import "fmt"

type cat struct {
	name        string
	age         int
	color       string
	description string
}

func (c *cat) run() string {
	fmt.Println("run")
	return "run"
}

func (c *cat) sleep() string {
	fmt.Println("sleep")
	return "sleep"
}

func (c *cat) smile() string {
	fmt.Println("smile")
	return "smile"
}
