package components

import (
	"fmt"
	"github.com/trustmaster/goflow"
	"io/ioutil"
)

type Readfile struct {
	flow.Component
	filename <-chan string // filename inport
	Lines    chan<- string // send line-by-line outport
}

func (c *Readfile) OnIn(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	c.Lines <- string(b)
}
