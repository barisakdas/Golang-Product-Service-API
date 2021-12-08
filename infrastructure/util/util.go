package util

import (
	"encoding/json"
	"net/http"
)

func HttpNotFound(w http.ResponseWriter) {
	resp := make(map[string]string)
	resp["message"] = "Resource not found."
	respMessage, _ := json.Marshal(resp)
	w.Write(respMessage)
}
