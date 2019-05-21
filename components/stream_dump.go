package components

// prints a LogStream do stdout in full detail, for debugging

import (
	"fmt"
	"github.com/trustmaster/goflow"
	"terreactive.ch/LogFlow/flow_types"
)

type StreamDump struct {
	flow.Component
	In <-chan flow_types.LogStream
}

func (x *StreamDump) OnIn(logmsg flow_types.LogStream) {
	if *logmsg.Raw != "" {
		fmt.Println("Raw: ", *logmsg.Raw)
	}
}
