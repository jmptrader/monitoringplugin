package monitoringplugin

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var (
	rangeRegex = regexp.MustCompile(`^(?P<invert>[@])?(?:(?P<start>(?:[\d]+|\~))(?::))?(?::)*(?P<end>[\d]+)?$`)
)

type Range struct {
	Noop   bool
	Invert bool
	Start  float64
	End    float64
	option string
}

func NewRange(option string) (r Range, err error) {
	r.Start = 0
	r.End = math.Inf(1)
	r.option = option

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
				return
			}
		}
	}

	endOption := parsedOption["end"]
	if endOption != "" {
		if r.End, err = strconv.ParseFloat(endOption, 64); err != nil {
			return
		}
	}

	return
}

func (r Range) parseRange(option string) (result map[string]string, err error) {
	result = make(map[string]string)

	match := rangeRegex.FindStringSubmatch(option)
	if match == nil {
		err = fmt.Errorf("Could not parse range: '%s'", option)
		return
	}

	for index, groupName := range rangeRegex.SubexpNames() {
		if index != 0 {
			result[groupName] = match[index]
		}
	}

	return
}

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

func (r Range) ToString() string {
	return r.option
}
