package main

import "fmt"
import "terreactive.ch/LogFlow/flow_types"

func Print_log_item(i flow_types.Log_item) {
	if i.Raw != nil {
		fmt.Println("Raw: ", *i.Raw)
	} else {
		fmt.Println("Raw logline missing")
	}

	if i.Syslog != nil {
		fmt.Println("Syslog: ", *i.Syslog)
	} else {
		fmt.Println("Syslog block is missing")
	}

	if i.Kvp != nil {
		fmt.Println("Key-Value pairs: ", *i.Kvp)
	} else {
		fmt.Println("Key-Value-Pair block is missing")
	}
}

func main() {
	var x flow_types.Log_item
	var logline flow_types.Log_raw
	logline = "Apr 12 16:33:03 proxy owi104: Apr 12 16:33:03 @MA1Wbpw--1uqtk-- user.info upstart: Connection from private client"
	x.Raw = &logline

	Print_log_item(x)
}
