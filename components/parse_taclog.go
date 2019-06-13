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
	"terreactive.ch/LogFlow/flow_types"
)

type ParseTaclog struct {
	flow.Component
	In <-chan flow_types.LogStream
}

func (x *ParseTaclog) OnIn(logmsg flow_types.LogStream) {
	if *logmsg.Raw != "" {
	}
}
