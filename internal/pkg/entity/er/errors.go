package er

// APPError declare custom error
type APPError struct {
	Status int    `json:"-"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Err    error  `json:"-"`
}

// WithError append error into APPError
func (e *APPError) WithError(err error) *APPError {
	return &APPError{
		Status: e.Status,
		Code:   e.Code,
		Msg:    e.Msg,
		Err:    err,
	}
}

func (e *APPError) Error() string {
	return e.Msg
}

func newAPPError(status int, code int, msg string) *APPError {
	return &APPError{Status: status, Code: code, Msg: msg}
}
