package main

import (
	"strings"
)

func ProcessStatus(n NagiosState) (int, string) {
	var rc int = NAGIOS_UNKNOWN
	var msg string
	var perfdata string

	u_str := strings.Join(n.Unknown, ", ")
	c_str := strings.Join(n.Critical, ", ")
	w_str := strings.Join(n.Warning, ", ")
	o_str := strings.Join(n.Ok, ", ")
	if len(n.PerfData) > 0 {
		perfdata = " | " + strings.Join(n.PerfData, " ")
	}

	if len(n.Unknown) > 0 {
		rc = NAGIOS_UNKNOWN
		msg = u_str
		if len(c_str) > 0 {
			msg += "; " + c_str
		}

		if len(w_str) > 0 {
			msg += "; " + w_str
		}

		if len(o_str) > 0 {
			msg += "; " + o_str
		}

		if len(perfdata) > 0 {
			msg += perfdata
		}
	} else if len(n.Critical) > 0 {
		rc = NAGIOS_CRITICAL
		msg = c_str

		if len(w_str) > 0 {
			msg += "; " + w_str
		}

		if len(o_str) > 0 {
			msg += "; " + o_str
		}

		if len(perfdata) > 0 {
			msg += perfdata
		}
	} else if len(n.Warning) > 0 {
		rc = NAGIOS_WARNING
		msg = w_str

		if len(o_str) > 0 {
			msg += "; " + o_str
		}
	} else if len(n.Ok) > 0 {
		rc = NAGIOS_OK
		msg = o_str

		if len(perfdata) > 0 {
			msg += perfdata
		}
	} else {
		// shouldn't happen
		rc = NAGIOS_UNKNOWN
		msg = "No results at all found"
	}

	return rc, msg
}
