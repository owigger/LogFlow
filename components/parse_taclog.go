package components

// parses raw field, fills taclog fields

// definition of a tacLOG header from msgid.c:
/*
 *  msgid_pos()
 *
 *  Quick logline parser. Searches only and finds the position of the MsgID.
 *
 *  Nov 23 05:10:53 platform host: Nov 23 05:10:53 @L2JLnMD2---1e--- logtext...
 *                  ^            ^                 ^
 *  <--------------> <~~~~~~~~~~> <--------------->
 *      16 bytes        variable       17 bytes
 *
 *  Rule: the '@' of the MsgID is always located 18 bytes after the first colon
 *  after position 16.
 *
 */

import (
	"github.com/trustmaster/goflow"
	"strings"
	"terreactive.ch/LogFlow/flow_types"
)

type ParseTaclog struct {
	flow.Component
	In  <-chan flow_types.LogStream
	Out chan<- flow_types.LogStream
}

func (x *ParseTaclog) OnIn(logmsg flow_types.LogStream) {
	var raw = *logmsg.Raw
	var length, colon, pltf int
	if *logmsg.Raw != "" {
		length = len(raw)
		if length < 35 {
			return // too short to contain a tacLOG line
		}
		colon = strings.Index(raw[16:], ":")
		if colon < 0 {
			return // no colon after position 16
		}
		colon += 16
		if raw[colon+18] != '@' {
			return // no msgid where expected
		}
		var taclogStruct flow_types.LogTaclog
		logmsg.Taclog = &taclogStruct

		taclogStruct.Sent = raw[0:15]
		pltf = strings.Index(raw[16:], " ")
		taclogStruct.Platform = raw[16 : pltf+16]
		taclogStruct.Host = raw[pltf+17 : colon]
		taclogStruct.Received = raw[colon+2 : colon+17]
		taclogStruct.Msgid = raw[colon+19 : colon+35]
		/*
			taclogStruct.Program      string    // the senders program name
			taclogStruct.Pid          int       // the senders process ID
			taclogStruct.Message      string    // the payload
			taclogStruct.AlevId       string    // Event_ID or Alert_ID
			taclogStruct.AlevCategory string    // log, event, or alert
			taclogStruct.AlevText     string    // Patrick, what is this? should use Message?
		*/
	}
	x.Out <- logmsg
}
