package helper

import (
	"encoding/json"
	"errors"
	"io"
)

type Rules struct {
}

func NewRules() *Rules {
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("body is invalid")
	}

	return model, json.NewDecoder(data).Decode(model)
}
