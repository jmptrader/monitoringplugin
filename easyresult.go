package monitoringplugin

// EasyResult is an implementation of CheckResult.
// Use this if you just need a simple check output.
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

func (easyResult *EasyResult) GetPerformanceData() (perfData map[string]Unit) {
	perfData = make(map[string]Unit)
	return
}
