package http_utils

import (
	"encoding/json"
	"github.com/rotelando/bookstore_items-api/utils/errors"
	"net/http"
)

func SendJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func SendError(w http.ResponseWriter, err *errors.RestErr) {
	SendJson(w, err.Status, err)
}
