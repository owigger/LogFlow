package main

import (
	"fmt"
	"github.com/trustmaster/goflow"
	"terreactive.ch/LogFlow/components"
)

// the test nework
type testNet struct {
	flow.Graph
}

// construct the entire network
func NewTestNet() *testNet {
	n := &testNet{}
	n.InitGraphState()
	n.Add(&components.Readfile{}, "Readfile")
	n.Add(&components.Print{}, "Print")
	n.Connect("Readfile", "Line", "Print", "In")
	n.Connect("Readfile", "Error", "Print", "In")
	n.MapInPort("In", "Readfile", "Filename")
	return n
}

func main() {
	fmt.Println("creating network")
	net := NewTestNet()
	in := make(chan string)
	fmt.Println("attaching input")
	net.SetInPort("In", in)
	flow.RunNet(net)

	fmt.Println("executing")
	in <- "t/data/20190412.base"
	in <- "no_such_file"
	in <- "t/data/textfile"

	close(in)
	<-net.Wait()
	fmt.Println("bye")
}
