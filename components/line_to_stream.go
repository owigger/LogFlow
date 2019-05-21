package components

// the receiving end
// interfaces a line-b<-line log source to the LogFlow system

import (
	"github.com/trustmaster/goflow"
	"terreactive.ch/LogFlow/flow_types"
)

type LineToStream struct {
	flow.Component
	In  <-chan string
	Out chan<- flow_types.LogStream
}

func (x *LineToStream) OnIn(logline string) {
	var newlog flow_types.LogStream
	newlog.Raw = &logline
}
