package monitoringplugin

import "fmt"

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

type EasyResult struct {
	status  int
	message string
}

func NewEasyResult(status int, message string) (easyResult *EasyResult) {
	easyResult = new(EasyResult)
	easyResult.status = status
	easyResult.message = message
	return
}

func (easyResult *EasyResult) GetStatus() (int, string) {
	return easyResult.status, easyResult.message
}

func (easyResult *EasyResult) GetLongOutput() (output string) {
	output = ""
	return
}

func (easyResult *EasyResult) GetPerformanceData() (perfData map[string]float64) {
	perfData = make(map[string]float64)
	return
}

type Check interface {
	Run() (checkResult CheckResult)
}
