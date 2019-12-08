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
	n.Add(&components.LineToStream{}, "LineToStream")
	n.Add(&components.ParseTaclog{}, "ParseTaclog")
	n.Add(&components.StreamDump{}, "StreamDump")
	n.Add(&components.Print{}, "Print")
	n.Connect("Readfile", "Line", "LineToStream", "In")
	n.Connect("LineToStream", "Out", "ParseTaclog", "In")
	n.Connect("ParseTaclog", "Out", "StreamDump", "In")
	n.Connect("Readfile", "Error", "Print", "In")
	n.MapInPort("In", "Readfile", "Filename")
	return n
}

func main() {
	fmt.Println("creating network")
	in := make(chan string)
	net := NewTestNet()
	net.SetInPort("In", in)
	flow.RunNet(net)

	in <- "t/data/20190412.base"
	in <- "t/data/20190610.app"

	close(in)
	<-net.Wait()
	fmt.Println("bye")
}
