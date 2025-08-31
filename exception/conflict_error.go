package exception

type ConflictError struct {
	Message string
}

func (c ConflictError) Error() string {
	return c.Message
}
