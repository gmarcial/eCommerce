package mock

//Logger mocking of the interface of logging
type Logger struct {
	InfowMock func(msg string, keysAndValues ...interface{})
	ErrorwMock func(msg string, keysAndValues ...interface{})
}

//Infow with behavior mocked
func (logger *Logger) Infow(msg string, keysAndValues ...interface{}) {
	if keysAndValues != nil {
		logger.InfowMock(msg, keysAndValues)
	}
	logger.InfowMock(msg, nil)
}

//Errorw with behavior mocked
func (logger *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	if keysAndValues != nil {
		logger.ErrorwMock(msg, keysAndValues)
	}
	logger.ErrorwMock(msg, nil)
}
