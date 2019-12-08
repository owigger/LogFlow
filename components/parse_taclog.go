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
	"strconv"
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
	var length, colon1, colon2, pltf, space2, bracket1 int
	var program string
	if *logmsg.Raw != "" {
		length = len(raw)
		if length < 35 {
			return // too short to contain a tacLOG line
		}
		colon1 = strings.Index(raw[16:], ":")
		if colon1 < 0 {
			return // no colon after position 16
		}
		colon1 += 16

		if raw[colon1+18] != '@' {
			return // no msgid where expected
		}
		var taclogStruct flow_types.LogTaclog
		logmsg.Taclog = &taclogStruct

		taclogStruct.Sent = raw[0:15]
		pltf = strings.Index(raw[16:], " ")
		taclogStruct.Platform = raw[16 : pltf+16]
		taclogStruct.Host = raw[pltf+17 : colon1]
		taclogStruct.Received = raw[colon1+2 : colon1+17]
		taclogStruct.Msgid = raw[colon1+19 : colon1+35]

		colon2 = strings.Index(raw[colon1+36:], ":")
		if colon2 < 0 {
			// no colon after MsgID, everything is log message
			taclogStruct.Message = raw[colon1+36:]
			return
		}
		colon2 += colon1 + 36

		space2 = strings.Index(raw[colon1+36:colon2], " ")
		if space2 < 0 {
			// illegal. assume everything is log message
			taclogStruct.Message = raw[colon1+36:]
			return
		}
		space2 += colon1 + 36

		taclogStruct.AlevId = raw[colon1+36 : space2]

		program = raw[space2+1 : colon2]
		bracket1 = strings.Index(program, "[")
		if bracket1 < 0 {
			taclogStruct.Program = program
		} else {
			// we have a PID within brackets
			taclogStruct.Program = program[0 : bracket1]
			taclogStruct.Pid, _ = strconv.Atoi(program[bracket1:len(program)])
		}

		taclogStruct.Message = raw[colon2+2:]

		/*
			Not yet parsed:
			taclogStruct.AlevCategory string    // log, event, or alert
			taclogStruct.AlevText     string    // Patrick, what is this? should use Message?
		*/
	}
	x.Out <- logmsg
}
