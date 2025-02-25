package api

import (
	"fmt"

	"github.com/wakatime/wakatime-cli/pkg/exitcode"
	"github.com/wakatime/wakatime-cli/pkg/wakaerror"
)

// Err represents a general api error.
type Err struct {
	Err error
}

var _ wakaerror.Error = Err{}

// Error method to implement error interface.
func (e Err) Error() string {
	return string(e.Err.Error())
}

// ExitCode method to implement wakaerror.Error interface.
func (Err) ExitCode() int {
	return exitcode.ErrAPI
}

// Message method to implement wakaerror.Error interface.
func (e Err) Message() string {
	return fmt.Sprintf("api error: %s", e.Err)
}

// ErrAuth represents an authentication error.
type ErrAuth struct {
	Err error
}

var _ wakaerror.Error = ErrAuth{}

// Error method to implement error interface.
func (e ErrAuth) Error() string {
	return string(e.Err.Error())
}

// ExitCode method to implement wakaerror.Error interface.
func (ErrAuth) ExitCode() int {
	return exitcode.ErrAuth
}

// Message method to implement wakaerror.Error interface.
func (e ErrAuth) Message() string {
	return fmt.Sprintf("invalid api key... find yours at wakatime.com/api-key. %s", e.Err)
}

// ErrBadRequest represents a 400 response from the API.
type ErrBadRequest struct {
	Err error
}

var _ wakaerror.Error = ErrBadRequest{}

// Error method to implement error interface.
func (e ErrBadRequest) Error() string {
	return string(e.Err.Error())
}

// ExitCode method to implement wakaerror.Error interface.
func (ErrBadRequest) ExitCode() int {
	return exitcode.ErrGeneric
}

// Message method to implement wakaerror.Error interface.
func (e ErrBadRequest) Message() string {
	return fmt.Sprintf("bad request: %s", e.Err)
}

// ErrBackoff means we send later because currently rate limited.
type ErrBackoff struct {
	Err error
}

var _ wakaerror.Error = ErrBackoff{}

// Error method to implement error interface.
func (e ErrBackoff) Error() string {
	return string(e.Err.Error())
}

// ExitCode method to implement wakaerror.Error interface.
func (ErrBackoff) ExitCode() int {
	return exitcode.ErrBackoff
}

// Message method to implement wakaerror.Error interface.
func (e ErrBackoff) Message() string {
	return fmt.Sprintf("rate limited: %s", e.Err)
}
