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
	s := logmsg.Taclog
	if s != nil {
		fmt.Println("Taclog: ")
		fmt.Println("  msgid   : ", s.Msgid)
		fmt.Println("  Sent    : ", s.Sent)
		fmt.Println("  Received: ", s.Received)
		fmt.Println("  Platform: ", s.Platform)
		fmt.Println("  Host    : ", s.Host)
		fmt.Println("  AlevId  : ", s.AlevId)
		fmt.Println("  Program : ", s.Program)
		fmt.Println("  Pid     : ", s.Pid)
		fmt.Println("  Message : ", s.Message)
	}
}

