package components

// the receiving end
// interfaces a line-b<-line log source to the LogFlow system

import (
	"github.com/trustmaster/goflow"
	"terreactive.ch/LogFlow/flow_types/flow_types"
)

type LineToStream struct {
	flow.Component
	In  <-chan string
	Out chan<- LogStream
}

func (x *LineToStram) OnIn(logline string) {
	var rawline string
	rawline <- logline
	x.Raw = *rawline
}
