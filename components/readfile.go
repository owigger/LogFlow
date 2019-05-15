package components

import (
	"bufio"
	"github.com/trustmaster/goflow"
	"os"
)

type Readfile struct {
	flow.Component
	Filename <-chan string // filename inport
	Line     chan<- string // send line-by-line outport
	Error    chan<- string  // send error messages
}

func (c *Readfile) OnFilename(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		errorstring := err.Error()
		c.Error <- errorstring
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Line <- string(scanner.Text())
	}
}
