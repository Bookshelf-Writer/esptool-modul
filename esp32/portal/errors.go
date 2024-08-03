package portal

import "errors"

//###########################################################//

var (
	ErrMismatchBytes  = errors.New("number of bytes received mismatch")
	ErrTimeoutRead    = errors.New("read timeout expired")
	ErrUnexpectedChar = errors.New("unexpected char after escape character")

	ErrResponseLength       = errors.New("invalid response length")
	ErrResponseStatusLength = errors.New("invalid response status length")
)
