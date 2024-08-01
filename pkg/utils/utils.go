package utils

import (
	"encoding/json"
	"net/http"
)

// ParseBody reads and parses the JSON body of an HTTP request.
func ParseBody(r *http.Request, x interface{}) error {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return err
	}
	return nil
}
