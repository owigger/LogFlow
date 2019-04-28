// an terminal node
// prints it input

package drain_print

import (
	"fmt"
)

type Drain_print struct {
	flow.Component
	In <-chan string
}

// Incoming line
func (c *Drain_print) OnPortName(argName string) {
	fmt.Println(string)
}
