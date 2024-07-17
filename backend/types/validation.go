package types

type ValidationError struct {
	Key    string `json:"key"`
	Reason string `json:"reason"`
}
