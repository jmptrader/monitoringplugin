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
	// GetStatus returns the output and the exit code of the plugin
	GetStatus() (exitCode int, message string)
	// GetLongOutput returns the long output (if there is any)
	GetLongOutput() (output string)
	// GetPerformanceData returns a perfdata name to perfdata value map
	GetPerformanceData() map[string]Unit
}

// Check is an interface for the check you want to run.
type Check interface {
	// Run is a function that runs the actual check.
	Run() (checkResult CheckResult)
}
