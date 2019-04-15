package flow_types

import "time"

// The Log_item is the basic data type that flows through LogFlow pipes.
// It's merely a container that contains links to various other types
// available for log messages in various phases of processing.
// Some basic fields (like timestamp) are always available.
type Log_item struct {
	Timestamp time.Time   // some idea of when this message was generated
	raw       *Log_raw    // link to raw logline
	syslog    *Log_syslog // link to RFC3164 syslog-parsed logline
	kvp       *Log_kvp    // link to the generic key-value pair map
}

// The unparsed, raw log message as received
type Log_raw string

// RFC 3164
type Log_syslog struct {
	severity int8
	facility int8
	header   string
	message  string
}

// canonical tacLOG message
type Log_taclog struct {
	s_time        time.Time
	r_time        time.Time
	platform      string
	host          string
	msgid         string
	program       string
	pid           int
	message       string
	alev_id       string
	alev_category string
	alev_text     string
}

// Key-Value pairs, parsed from message
type Log_kvp map[string]string
