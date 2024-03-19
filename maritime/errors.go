package maritime

import (
	"errors"
	"fmt"
)

//go:generate stringer -type=ValidationError -output=error_string.go

const (
	ErrUnknown ValidationError = iota + 50000
	ErrVesselIDOrShipIDRequired
	ErrShipNameOrMMSIOrShipIDOrIMORequired
)

// ValidationError is an error implementation that transports as a const number
type ValidationError int

func (e ValidationError) Error() string { return e.String() }

type ServiceError struct {
	Errors []struct {
		Code   string `json:"code"`
		Detail string `json:"detail"`
	} `json:"errors"`
}

func (e ServiceError) Error() string {
	if len(e.Errors) == 0 {
		return ""
	}

	var err error
	for _, ee := range e.Errors {
		err = errors.Join(err, fmt.Errorf("%s: %s", ee.Code, ee.Detail))
	}

	return err.Error()
}
