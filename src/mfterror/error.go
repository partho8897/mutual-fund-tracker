package mfterror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MFTError struct {
	ErrType MFTErrorTypes `json:"errType"`
	ErrMsg  string        `json:"errMsg"`
	Details string        `json:"details"`
}

func (e MFTError) Error() string {
	errBytes, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("ErrorType:%v ErrorEnum:%v ErrorDetails:%v", e.ErrType, e.ErrMsg, e.Details)
	} else {
		return string(errBytes)
	}
}

func (e MFTError) WithDetails(details string) *MFTError {
	e.Details = details
	return &e
}

func (e MFTError) GetHTTPErrorCode() int {
	if e.ErrType == ErrorTypeInvalidArgument || e.ErrType == ErrorTypeNotSupported {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
