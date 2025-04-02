package exception

type NotFountError struct {
	Error string
}

func NewNotFoundError(error string) NotFountError {
	return NotFountError{Error: error}
}
