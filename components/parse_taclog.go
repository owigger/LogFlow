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
	var length, colon int
	var msgid string
	if *logmsg.Raw != "" {
		length = len(raw)
		if length < 35 {
			return // too short to contain a tacLOG line
		}
		colon = strings.Index(raw[16:], ":")
		if colon < 0 {
			return // no colon after position 16
		}
		if raw[colon+16+18] != '@' {
			return // no msgid where expected
		}
		msgid = raw[colon+16+19 : colon+16+35]
		var taclogStruct flow_types.LogTaclog
		logmsg.Taclog = &taclogStruct
		taclogStruct.Msgid = msgid
	}
	x.Out <- logmsg
}
