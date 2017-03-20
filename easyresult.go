package monitoringplugin

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
