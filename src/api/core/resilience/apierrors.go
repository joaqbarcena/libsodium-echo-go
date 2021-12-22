package resilience

type ApiError struct {
	status  int
	message string
	cause   error
}

func (err *ApiError) Error() string {
	return err.message
}

func (err *ApiError) Status() int {
	return err.status
}

func NewApiError(status int, message string) *ApiError {
	return &ApiError{
		status:  status,
		message: message,
		cause:   nil,
	}
}
