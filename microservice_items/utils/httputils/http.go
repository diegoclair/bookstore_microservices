package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// RespondJSON - to handle json response
func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// RespondError - to handle json response
func RespondError(w http.ResponseWriter, err resterrors.RestErr) {
	RespondJSON(w, err.StatusCode(), err)
}
