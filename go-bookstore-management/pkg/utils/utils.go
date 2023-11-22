package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// Get request and marshal this using json
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if _, err := json.Marshal([]byte(body)); err != nil {
			return
		}
	}
}
