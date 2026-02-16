package goledger

import "encoding/json"

func decodeJSONLoose(in []byte) any {
	var v any
	if err := json.Unmarshal(in, &v); err != nil {
		return nil
	}
	return v
}
