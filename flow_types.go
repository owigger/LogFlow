package flow_types

import "time"

// a container for a single log message, in one or more forms
type Log_item struct {
    raw *Log_raw
    syslog *Log_syslog
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
    s_time time.Time
    r_time time.Time
    platform string
    host string
    msgid string
    program string
    pid int
    message string
    alev_id string
    alev_category string
    alev_text string
}

