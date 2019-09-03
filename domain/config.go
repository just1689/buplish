package domain

import "encoding/json"

type Config []Action

type Action struct {
	Action     string          `json:"action"`
	Parameters json.RawMessage `json:"parameters"`
}

type ParametersBuild struct {
	Tag        string `json:"tag"`
	Dockerfile string `json:"dockerfile"`
}

type ParametersPush struct {
	Tag string `json:"tag"`
}

type ParametersCall struct {
	Method string `json:"method"`
	URI    string `json:"uri"`
}
