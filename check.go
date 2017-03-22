package monitoringplugin

// Nagios Status Types
const (
	OK       = 0
	WARNING  = 1
	CRTIICAL = 2
	UNKNOWN  = 3
)

// CheckResult is an interface that provides data for the plugin output.
type CheckResult interface {
	GetStatus() (exitCode int, message string)
	GetLongOutput() (output string)
	GetPerformanceData() map[string]float64
}

// Check is an interface for the check you want to run.
type Check interface {
	// Run is a function that runs the actual check.
	Run() (checkResult CheckResult)
}
