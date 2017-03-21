package monitoringplugin

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var (
	rangeRegex = regexp.MustCompile(`^(?P<invert>[@])?(?:(?P<start>(?:\-?[\d]+(?:\.\d+)?|\~))(?::))?(?::)*(?P<end>\-?[\d]+(?:\.\d+)?)?$`)
)

// Range represents a range specified by nagios plugin guidelines
// See https://nagios-plugins.org/doc/guidelines.html#THRESHOLDFORMAT for further informations
type Range struct {
	Noop   bool
	Invert bool
	Start  float64
	End    float64
}

// NewRange from a text representation of a nagios range value
// You may use this to parse range parameters for your plugin
func NewRange(option string) (r Range, err error) {
	r.Start = 0
	r.End = math.Inf(1)

	if option == "" {
		r.Noop = true
		return
	}

	parsedOption, err := r.parseRange(option)
	if err != nil {
		return
	}

	if parsedOption["invert"] == "@" {
		r.Invert = true
	}

	startOption := parsedOption["start"]
	if startOption != "" {
		if startOption == "~" {
			r.Start = math.Inf(-1)
		} else {
			if r.Start, err = strconv.ParseFloat(startOption, 64); err != nil {
				err = fmt.Errorf("Can't parse start float: %s", err)
				return
			}
		}
	}

	endOption := parsedOption["end"]
	if endOption != "" {
		if r.End, err = strconv.ParseFloat(endOption, 64); err != nil {
			err = fmt.Errorf("Can't parse end float: %s", err)
			return
		}
	}

	if r.End < r.Start {
		err = fmt.Errorf("End(%f) is smaller than start(%f)", r.End, r.Start)
		return
	}

	return
}

func (r Range) parseRange(option string) (result map[string]string, err error) {
	result = make(map[string]string)

	match := rangeRegex.FindStringSubmatch(option)
	if match == nil {
		err = fmt.Errorf("Invalid range format: '%s'", option)
		return
	}

	for index, groupName := range rangeRegex.SubexpNames() {
		if index != 0 {
			result[groupName] = match[index]
		}
	}

	return
}

// Check whether the value meets the requirements.
func (r Range) Check(value float64) (result bool) {
	result = false
	if r.Noop {
		return
	}

	if value < r.Start {
		result = true
	}

	if value > r.End {
		result = true
	}

	if r.Invert {
		result = !result
	}
	return
}

// ToString returns a perfdata compatible text representation of the range
func (r Range) ToString() string {
	if r.Noop {
		return ""
	}

	buffer := new(bytes.Buffer)
	if r.Invert {
		buffer.WriteByte('@')
	}
	if r.Start != 0.0 {
		if math.IsInf(r.Start, -1) {
			buffer.WriteString("~:")
		} else {
			buffer.WriteString(strconv.FormatFloat(r.Start, 'f', -1, 64))
			buffer.WriteByte(':')
		}
	}
	if !math.IsInf(r.End, 1) {
		buffer.WriteString(strconv.FormatFloat(r.End, 'f', -1, 64))
	}

	return buffer.String()
}
