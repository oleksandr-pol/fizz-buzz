package game

type ReaderError struct {
	msg string
}

func (e *ReaderError) Error() string {
	return e.msg
}
