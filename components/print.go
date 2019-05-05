package components

import (
	"fmt"
	"github.com/trustmaster/goflow"
)

type Print struct {
	flow.Component
	In <-chan string
}

// Incoming line
func (c *Print) OnIn(a string) {
	fmt.Println(a)
}
