package monitoringplugin

import "fmt"

type DefaultCheckResultOpts struct {
	DefaultStatus  int
	DefaultMessage string
}

type DefaultCheckResult struct {
	currentStatus  int
	message        string
	defaultMessage string
	longOutput     string
	perfData       map[string]float64
}

func NewDefaultCheckResult(options *DefaultCheckResultOpts) (checkResult *DefaultCheckResult) {
	checkResult = new(DefaultCheckResult)
	if options == nil {
		options = &DefaultCheckResultOpts{
			DefaultStatus:  3,
			DefaultMessage: "No result provided!",
		}
	}
	checkResult.currentStatus = options.DefaultStatus
	checkResult.defaultMessage = options.DefaultMessage
	checkResult.perfData = make(map[string]float64)
	return
}

func (checkResult *DefaultCheckResult) GetStatus() (exitCode int, message string) {
	exitCode = checkResult.currentStatus
	message = checkResult.message
	if message == "" {
		message = checkResult.defaultMessage
	}
	return
}

func (checkResult *DefaultCheckResult) GetLongOutput() string {
	return checkResult.longOutput
}

func (checkResult *DefaultCheckResult) GetPerformanceData() map[string]float64 {
	return checkResult.perfData
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
		checkResult.message = fmt.Sprintf("%s; %s", checkResult.message, message)
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

func (checkResult *DefaultCheckResult) SetPerformanceData(label string, value float64) {
	checkResult.perfData[label] = value
}
