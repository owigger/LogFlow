package flow_types

import "time"

// The Log_item is the basic data type that flows through LogFlow pipes.
// It's merely a container that contains links to various other types
// available for log messages in various phases of processing.
// Some basic fields (like timestamp) are always available.
type Log_item struct {
	Timestamp time.Time   // some idea of when this message was generated
	Raw       *Log_raw    // link to raw logline
	Syslog    *Log_syslog // link to RFC3164 syslog-parsed logline
	Kvp       *Log_kvp    // link to the generic key-value pair map
}

// The unparsed, raw log message as received
type Log_raw string

// RFC 3164
type Log_syslog struct {
	Severity int8
	Facility int8
	Header   string
	Message  string
}

// canonical tacLOG message
type Log_taclog struct {
	S_time        time.Time
	R_time        time.Time
	Platform      string
	Host          string
	Msgid         string
	Program       string
	Pid           int
	Message       string
	Alev_id       string
	Alev_category string
	Alev_text     string
}

// Key-Value pairs, parsed from message
type Log_kvp map[string]string
