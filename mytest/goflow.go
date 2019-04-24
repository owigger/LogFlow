package main

import (
	"fmt"
	"github.com/goflow"
	"regexp"
	"strings"
)

type count struct {
	tag   string
	count int
}

type splitter struct {
	goflow.Component

	In         <-chan string
	Out1, Out2 chan<- string
}

func (t *splitter) OnIn(s string) {
	t.Out1 <- s
	t.Out2 <- s
}

type wordCounter struct {
	goflow.Component
	Sentence <-chan string
	Count    chan<- *count
}

func (wc *wordCounter) OnSentence(sentence string) {
	wc.Count <- &count{"Words", len(strings.Split(sentence, " "))}
}

type letterCounter struct {
	goflow.Component
	Sentence <-chan string
	Count    chan<- *count
	re       *regexp.Regexp
}

func (lc *letterCounter) OnSentence(sentence string) {
	lc.Count <- &count{"Letters", len(lc.re.FindAllString(sentence, -1))}
}

func (lc *letterCounter) Init() {
	lc.re = regexp.MustCompile("[a-zA-Z]")
}

type printer struct {
	goflow.Component
	Line <-chan *count //inport
}

func (p *printer) OnLine(c *count) {
	fmt.Println(c.tag+":", c.count)
}

type counterNet struct {
	goflow.Graph
}

// Assembling the network
func NewCounterNet() *counterNet {
	n := &counterNet{}
	n.InitGraphState()
	n.Add(&splitter{}, "splitter")
	n.Add(&wordCounter{}, "wordCounter")
	n.Add(&letterCounter{}, "letterCounter")
	n.Add(&printer{}, "printer")
	n.Connect("splitter", "Out1", "wordCounter", "Sentence")
	n.Connect("splitter", "Out2", "letterCounter", "Sentence")
	n.Connect("wordCounter", "Count", "printer", "Line")
	n.Connect("letterCounter", "Count", "printer", "Line")
	n.MapInPort("In", "splitter", "In")
	return n
}

//Launching the network
func main() {
	net := NewCounterNet()
	in := make(chan string)
	net.SetInPort("In", in)
	goflow.RunNet(net)
	in <- "I never put off till tomorrow what I can do the day after."
	in <- "Fashion is a form of ugliness so intolerable that we hae to alter it every six months."
	in <- "Life is too important to be taken seriously."
	close(in)
	<-net.Wait()
}
