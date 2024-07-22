package types

type ErrNotFound struct {
	msg string
}

func NewErrNotFound(msg string) error {
	return &ErrNotFound{msg: msg}
}

func (e *ErrNotFound) Error() string {
	return e.msg
}
