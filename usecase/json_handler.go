package usecase

import (
	"encoding/json"
	"io"
)

type WithJSON interface {
	Length() int
}

func ToJSON(w io.Writer, data WithJSON) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func FromJSON(r io.Reader, data *WithJSON) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
