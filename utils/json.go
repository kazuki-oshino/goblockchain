package utils

import "encoding/json"

// JSONStatus is return json message.
func JSONStatus(message string) []byte {
	m, _ := json.Marshal(struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
	return m
}
