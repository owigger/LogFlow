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
	n.MapInPort("filename", "Readfile", "filename")
	return n
}

func main() {
	fmt.Println("creating netowrk")
	net := NewTestNet()
	in := make(chan string)
	fmt.Println("attaching input")
	net.SetInPort("In", in)
	flow.RunNet(net)

	fmt.Println("executing")
	in <- "t/data/20190412.base"

	close(in)
	// <-net.Wait
	fmt.Println("bye")
}
