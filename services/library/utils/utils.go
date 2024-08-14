package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	//Check if body is empty
	if r.Body == nil {
		return fmt.Errorf("MISSING REQUEST BODY")
	}
	return json.NewDecoder(r.Body).Decode(payload) //Create new decoder which will read data from r.Body, then decode reads data from json and decode it to payload
}
