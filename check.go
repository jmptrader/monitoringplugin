package monitoringplugin

// Nagios Status Types
const (
	OK       = 0
	WARNING  = 1
	CRTIICAL = 2
	UNKNOWN  = 3
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
	Warning           Range
	Critical          Range
	Minimum           float64
	Maximum           float64
}

type CheckResult interface {
	GetStatus() (exitCode int, message string)
	GetLongOutput() (output string)
	GetPerformanceData() map[string]float64
}

type Check interface {
	Run() (checkResult CheckResult)
}

func UnitToString(unit int) string {
	switch unit {
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
