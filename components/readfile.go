package components

import (
	"bufio"
	"github.com/trustmaster/goflow"
	"os"
)

type Readfile struct {
	flow.Component
	filename <-chan string // filename inport
	Line     chan<- string // send line-by-line outport
	Error    chan<- error  // send error messages
}

func (c *Readfile) OnIn(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		c.Error <- err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Line <- string(scanner.Text())
	}
}
