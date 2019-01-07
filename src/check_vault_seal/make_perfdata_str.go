package main

import (
	"errors"
	"fmt"
	"strings"
)

func MakePerfDataString(label string, val string, uom *string, warn *string, crit *string, min *string, max *string) (string, error) {
	var perfdata string
	var _lbl string
	var _val string
	var _uom string
	var _warn string
	var _crit string
	var _min string
	var _max string

	/*
	   This is the expected format:

	   'label'=value[UOM];[warn];[crit];[min];[max]

	   Notes:

	       * space separated list of label/value pairs
	       * label can contain any characters except the equals sign or single quote (')
	       * the single quotes for the label are optional. Required if spaces are in the label
	       * label length is arbitrary, but ideally the first 19 characters are unique (due to a limitation in RRD). Be aware of a limitation in the amount of data that NRPE returns to Nagios
	         to specify a quote character, use two single quotes
	       * warn, crit, min or max may be null (for example, if the threshold is not defined or min and max do not apply). Trailing unfilled semicolons can be dropped
	       * min and max are not required if UOM=%
	       * value, min and max in class [-0-9.]. Must all be the same UOM. value may be a literal "U" instead, this would indicate that the actual value couldn't be determined
	       * warn and crit are in the range format (see the Section called Threshold and Ranges). Must be the same UOM
	       * UOM (unit of measurement) is one of:
	           * no unit specified - assume a number (int or float) of things (eg, users, processes, load averages)
	           * s - seconds (also us, ms)
	           * % - percentage
	           * B - bytes (also KB, MB, TB)
	           * c - a continous counter (such as bytes transmitted on an interface)
	       * It is up to third party programs to convert the Nagios Plugins performance data into graphs.

	   Source: https://nagios-plugins.org/doc/guidelines.html#AEN200
	*/

	_lbl = strings.TrimSpace(label)
	if _lbl == "" {
		return perfdata, errors.New("Label can't be empty")
	}

	_val = strings.TrimSpace(val)
	if _val == "" {
		return perfdata, errors.New("Value can't be empty")
	}

	// check if unit of measurement is valid
	if uom != nil {
		if *uom != "" && *uom != "s" && *uom != "ms" && *uom != "us" && *uom != "ns" && *uom != "%" && *uom != "B" && *uom != "KB" && *uom != "MB" && *uom != "GB" && *uom != "TB" && *uom != "c" {
			return perfdata, errors.New(fmt.Sprintf("Unit of measurement %s is invalid", *uom))
		}
		_uom = strings.TrimSpace(*uom)
	}

	if strings.Contains(label, "'") || strings.Contains(label, "=") {
		return perfdata, errors.New(fmt.Sprintf("Label must not contain single quote (') or equal sign (=) character"))
	}

	if warn != nil {
		_warn = strings.TrimSpace(*warn)
	}

	if crit != nil {
		_crit = strings.TrimSpace(*crit)
	}

	if min != nil {
		_min = strings.TrimSpace(*min)
	}

	if max != nil {
		_max = strings.TrimSpace(*max)
	}

	perfdata = fmt.Sprintf("'%s'=%s%s;%s;%s;%s;%s", _lbl, _val, _uom, _warn, _crit, _min, _max)

	// remove trailing semicolon(s)
	perfdata = strings.TrimRight(perfdata, ";")
	return perfdata, nil
}
