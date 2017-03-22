package monitoringplugin

import (
	"fmt"
	"math"
	"strconv"
)

// Nagios Performance Data Units Of Measurement
const (
	NUMBER     = 0
	SECONDS    = 1
	PERCENTAGE = 2
	BYTES      = 3
	KILOBYTES  = 4
	MEGABYTES  = 5
	GIGABYTES  = 6
	TERABYTES  = 7
	PETABYTES  = 8
	EXABYTES   = 9
	ZETTABYTES = 10
	COUNTER    = 11
)

type PerformanceDataSpec struct {
	Label             string
	UnitOfMeasurement int
	Warning           *Range
	Critical          *Range
	Minimum           float64
	Maximum           float64
}

func (perfDataSpec PerformanceDataSpec) unitString() string {
	switch perfDataSpec.UnitOfMeasurement {
	case SECONDS:
		return "s"
	case PERCENTAGE:
		return "%"
	case BYTES:
		return "B"
	case KILOBYTES:
		return "KB"
	case MEGABYTES:
		return "MB"
	case GIGABYTES:
		return "GB"
	case TERABYTES:
		return "TB"
	case PETABYTES:
		return "PB"
	case EXABYTES:
		return "EB"
	case ZETTABYTES:
		return "ZB"
	case COUNTER:
		return "c"
	default:
		return ""
	}
}

func (perfDataSpec PerformanceDataSpec) FormatPerfDataFromMap(perfData map[string]float64) string {
	var perfValue string
	value, hasValue := perfData[perfDataSpec.Label]
	if hasValue {
		perfValue = strconv.FormatFloat(value, 'f', -1, 64)
	} else {
		perfValue = "U"
	}
	return perfDataSpec.FormatPerfData(perfValue)
}

func (perfDataSpec PerformanceDataSpec) FormatPerfData(perfValue string) string {
	var (
		minValue  string
		maxValue  string
		warnValue string
		critValue string
	)

	if math.IsInf(perfDataSpec.Minimum, -1) {
		minValue = ""
	} else {
		minValue = strconv.FormatFloat(perfDataSpec.Minimum, 'f', -1, 64)
	}

	if math.IsInf(perfDataSpec.Maximum, 1) {
		maxValue = ""
	} else {
		maxValue = strconv.FormatFloat(perfDataSpec.Maximum, 'f', -1, 64)
	}

	if perfDataSpec.Warning != nil {
		warnValue = perfDataSpec.Warning.ToString()
	}

	if perfDataSpec.Critical != nil {
		critValue = perfDataSpec.Critical.ToString()
	}

	return fmt.Sprintf("'%s'=%s%s;%s;%s;%s;%s",
		perfDataSpec.Label, perfValue, perfDataSpec.unitString(),
		warnValue, critValue,
		minValue, maxValue)
}
