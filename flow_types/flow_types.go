package flow_types

import "time"

// The LogStream is the basic data type that flows through LogFlow pipes.
// It's merely a container that contains links to various other types
// available for log messages in various phases of processing.
// Some basic fields (like timestamp) are always available.
type LogStream struct {
	Timestamp time.Time  // our idea of when this message was generated
	Raw       *string    // raw logline string
	Syslog    *LogSyslog // RFC3164 syslog-parsed logline structure
	Taclog    *LogTaclog // tacLOG parsed logline header+message
	Kvp       *LogKvp    // generic key-value pair structure
}

// RFC 3164
type LogSyslog struct {
	Severity int8
	Facility int8
	Header   string
	Message  string
}

// canonical tacLOG message
type LogTaclog struct {
	Stime        time.Time // syslog's idea of when the messages was received
	Rtime        time.Time // the senders idea of when the messages was sent
	Platform     string    // the tacLOG platform
	Host         string    // host name of log source
	Msgid        string    // 16 character tacLOG unique messasge ID
	Program      string    // the senders program name
	Pid          int       // the senders process ID
	Message      string    // the payload
	AlevId       string    // Event_ID or Alert_ID
	AlevCategory string    // log, event, or alert
	AlevText     string    // Patrick, what is this? should use Message?
}

// Key-Value pairs, parsed from message
type LogKvp map[string]string
