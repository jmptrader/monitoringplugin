package monitoringplugin

// Nagios Status Types
const (
	OK       = 0
	WARNING  = 1
	CRTIICAL = 2
	UNKNOWN  = 3
)

type CheckResult interface {
	GetStatus() (exitCode int, message string)
	GetLongOutput() (output string)
	GetPerformanceData() map[string]float64
}

type Check interface {
	Run() (checkResult CheckResult)
}
