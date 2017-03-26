package monitoringplugin

import (
	"fmt"
	"math"
	"strconv"
)

// PerformanceDataSpec is the specification of the performance data.
type PerformanceDataSpec struct {
	Label             string
	UnitOfMeasurement UnitSpecification
	Warning           *Range
	Critical          *Range
	Minimum           float64
	Maximum           float64
}

// FormatPerfDataFromMap calls FormatPerfData with the value from perfData.
// If perfData doesn't have a value with the label of this specification, the value is 'U'.
func (perfDataSpec PerformanceDataSpec) FormatPerfDataFromMap(perfData map[string]Unit) string {
	var perfValue string
	value, hasValue := perfData[perfDataSpec.Label]
	if hasValue {
		perfValue = strconv.FormatFloat(value.Value(), 'f', -1, 64)
	} else {
		perfValue = "U"
	}
	return perfDataSpec.FormatPerfData(perfValue)
}

// FormatPerfData outputs the performance value (float or 'U') with the specified UOM and Limits.
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
		perfDataSpec.Label, perfValue, perfDataSpec.UnitOfMeasurement.String(),
		warnValue, critValue,
		minValue, maxValue)
}
