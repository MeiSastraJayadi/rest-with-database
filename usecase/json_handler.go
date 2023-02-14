package usecase

import (
	"encoding/json"
	"io"
)

type WithJSONEncode interface {
	Length() int
}

type UseJSONDecode interface {
	GetName() string
}

func ToJSON(w io.Writer, data WithJSONEncode) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func FromJSON(r io.Reader, data UseJSONDecode) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
