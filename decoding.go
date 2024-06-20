package bencoding

import (
	"bufio"
	"errors"
)

type Decoder struct {
	r bufio.Reader
}

func (d Decoder) decodeString() (string, error) {
	return "", errors.New("unimplemented")
}

func (d Decoder) decodeNumber() (int, error) {
	return 0, errors.New("unimplemented")
}

func (d Decoder) decodeList() ([]any, error) {
	return []any{}, errors.New("unimplemented")
}

func (d Decoder) decodeMap() (map[string]any, error) {
	return map[string]any{}, errors.New("unimplemented")
}

func (d Decoder) Decode() (any, error) {
	return nil, errors.New("unimplemented")
}
