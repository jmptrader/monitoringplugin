package monitoringplugin

import "fmt"

type DefaultCheckResultOpts struct {
	DefaultStatus  int
	DefaultMessage string
}

type DefaultCheckResult struct {
	currentStatus       int
	message             string
	defaultMessage      string
	longOutput          string
	perfData            map[string]Unit
	dynamicPerfDataSpec []PerformanceDataSpec
}

func NewDefaultCheckResult(options *DefaultCheckResultOpts) (checkResult *DefaultCheckResult) {
	checkResult = new(DefaultCheckResult)
	if options == nil {
		options = &DefaultCheckResultOpts{
			DefaultStatus:  3,
			DefaultMessage: "Internal error: No result from check",
		}
	}
	checkResult.currentStatus = options.DefaultStatus
	checkResult.defaultMessage = options.DefaultMessage
	checkResult.perfData = make(map[string]Unit)
	checkResult.dynamicPerfDataSpec = make([]PerformanceDataSpec, 0)
	return
}

// GetStatus implements CheckResult.GetStatus
func (checkResult *DefaultCheckResult) GetStatus() (exitCode int, message string) {
	exitCode = checkResult.currentStatus
	message = checkResult.message
	if message == "" {
		message = checkResult.defaultMessage
	}
	return
}

// GetLongOutput implements CheckResult.GetLongOutput
func (checkResult *DefaultCheckResult) GetLongOutput() string {
	return checkResult.longOutput
}

// GetPerformanceData implements CheckResult.GetPerformanceData
func (checkResult *DefaultCheckResult) GetPerformanceData() map[string]Unit {
	return checkResult.perfData
}

func (CheckResult *DefaultCheckResult) GetDynamicPerformanceDataSpec() []PerformanceDataSpec {
	return CheckResult.dynamicPerfDataSpec
}

func (checkResult *DefaultCheckResult) compareStatus(newStatus int) bool {
	if checkResult.currentStatus >= 0 && checkResult.currentStatus < 3 && newStatus <= checkResult.currentStatus {
		return false
	}
	return true
}

func (checkResult *DefaultCheckResult) AddResult(status int, message string) {
	if checkResult.compareStatus(status) {
		checkResult.currentStatus = status
	}
	if message != "" {
		if checkResult.message == "" {
			checkResult.message = message
		} else {
			checkResult.message = fmt.Sprintf("%s; %s", checkResult.message, message)
		}
	}
}

func (checkResult *DefaultCheckResult) SetResult(status int, message string) {
	checkResult.currentStatus = status
	checkResult.message = message
}

func (checkResult *DefaultCheckResult) AppendToLongOutput(message string) {
	checkResult.longOutput += message
}

func (checkResult *DefaultCheckResult) SetLongOutput(message string) {
	checkResult.longOutput = message
}

func (checkResult *DefaultCheckResult) SetPerformanceData(label string, value Unit) {
	checkResult.perfData[label] = value
}

func (checkResult *DefaultCheckResult) SetPerformanceDataSpec(spec PerformanceDataSpec) {
	checkResult.dynamicPerfDataSpec = append(checkResult.dynamicPerfDataSpec, spec)
}
