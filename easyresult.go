package monitoringplugin

// EasyResult is an implementation of CheckResult.
// Use this if you just need a simple check output.
type EasyResult struct {
	status  int
	message string
}

// NewEasyResult is a constructor for EasyResult.
func NewEasyResult(status int, message string) (easyResult *EasyResult) {
	easyResult = new(EasyResult)
	easyResult.status = status
	easyResult.message = message
	return
}

// GetStatus implements CheckResult.GetStatus
func (easyResult *EasyResult) GetStatus() (int, string) {
	return easyResult.status, easyResult.message
}

// GetLongOutput implements CheckResult.GetLongOutput
func (easyResult *EasyResult) GetLongOutput() (output string) {
	output = ""
	return
}

// GetPerformanceData implements CheckResult.GetPerformanceData
func (easyResult *EasyResult) GetPerformanceData() (perfData map[string]Unit) {
	perfData = make(map[string]Unit)
	return
}
