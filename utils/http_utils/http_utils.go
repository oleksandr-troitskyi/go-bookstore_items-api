package http_utils

import (
	"encoding/json"
	"github.com/troitskyA/go-bookstore_utils-lib/rest_errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err *rest_errors.RestErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}
