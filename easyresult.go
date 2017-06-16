package monitoringplugin

// EasyResult is an implementation of CheckResult.
// Use this if you just need a simple check output.
type EasyResult struct {
	Status  int
	Message string
}

// GetStatus implements CheckResult.GetStatus
func (easyResult *EasyResult) GetStatus() (int, string) {
	return easyResult.Status, easyResult.Message
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

func (easyResult *EasyResult) GetDynamicPerformanceDataSpec() (spec []PerformanceDataSpec) {
	spec = make([]PerformanceDataSpec, 0)
	return
}
