package main

import (
	"encoding/json"
)

func mustMarshalErrToJSON(err error) []byte {
	var e = struct {
		Reason string `json:"reason"`
	}{
		Reason: err.Error(),
	}
	b, marshalErr := json.Marshal(e)
	if marshalErr != nil {
		panic(marshalErr)
	}
	return b
}
