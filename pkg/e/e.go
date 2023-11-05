package e

type Error struct {
	Code int
	Msg  string
	Err  error
}

func New(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) WithMsg(msg string) *Error {
	e.Msg = msg
	return e
}

func (e *Error) WithErr(err error) *Error {
	e.Err = err
	return e
}
