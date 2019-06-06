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
	in := make(chan string)
	net := NewTestNet()
	net.SetInPort("In", in)
	flow.RunNet(net)
	in <- "t/data/20190412.base"
	close(in)
	<-net.Wait()
	fmt.Println("bye")
}
