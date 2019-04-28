// an input node
// reads a single text file

// 2019-04-20 owi - this is my first LogFlow component

package source_file

import "io/ioutil"

type source_file struct {
	flow.Component
	filename <-chan string // filename inport
	Line     chan<- string // send line-by-line outport
}

func (c *Source_file) OnIn(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(b))
}
