package utils

import (
	"encoding/json"
	"github.com/mutual-fund-tracker/src/mfterror"
	"io"
	"log"
	"net/http"
)

func DecodeResponseBody(body *http.Response, target interface{}) *mfterror.MFTError {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v\n", err)
		}
	}(body.Body)

	if decodingErr := json.NewDecoder(body.Body).Decode(target); decodingErr != nil {
		log.Printf("Error decoding response body: %v\n", decodingErr)
		return mfterror.ERR_UNKNOWN.WithDetails("Error decoding response body")
	}

	return nil
}
