package errors

import "fmt"

const (
	NotFoundUser = iota
)

var errMessage = map[int64]string{
	NotFoundUser: "Not Found User",
}

func ErrOf(code int64, args ...interface{}) error {
	if message, ok := errMessage[code]; ok {
		return fmt.Errorf("%s : %v", message, args)
	} else {
		return fmt.Errorf("not found err code")
	}
}
