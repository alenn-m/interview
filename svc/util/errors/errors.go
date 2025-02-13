package errors

type ErrValidation struct {
	Message string
}

func (e ErrValidation) Error() string {
	return e.Message
}
