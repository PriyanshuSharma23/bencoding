package bencoding

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: w,
	}
}

func (e Encoder) encodeStr(s string) error {
	_, err := e.w.Write([]byte(strconv.Itoa(len(s)) + ":" + s))
	return err
}

func (e Encoder) encodeNumber(i int) error {
	_, err := e.w.Write([]byte("i" + strconv.Itoa(i) + "e"))
	return err
}

func (e Encoder) encodeDict(dict map[string]any) error {
	_, err := e.w.Write([]byte("d"))
	if err != nil {
		return err
	}

	for k, v := range dict {
		err := e.encodeStr(k)
		if err != nil {
			return err
		}

		err = e.Encode(v)
		if err != nil {
			return err
		}
	}

	_, err = e.w.Write([]byte("e"))
	if err != nil {
		return err
	}

	return nil
}

func (e Encoder) encodeList(list []any) error {
	_, err := e.w.Write([]byte("l"))
	if err != nil {
		return err
	}

	for _, v := range list {
		err := e.Encode(v)
		if err != nil {
			return err
		}
	}

	_, err = e.w.Write([]byte("e"))
	if err != nil {
		return err
	}

	return nil
}

func (e Encoder) Encode(val interface{}) error {
	var err error = nil

	switch s := reflect.TypeOf(val).Kind(); s {
	case reflect.Int:
		err = e.encodeNumber(val.(int))
	case reflect.String:
		err = e.encodeStr(val.(string))
	case reflect.Map:
		v, ok := val.(map[string]any)
		if !ok {
			return fmt.Errorf("unsupported type, only accept map[string]any")
		}
		err = e.encodeDict(v)
	case reflect.Slice:
		err = e.encodeList(val.([]any))
	default:
		return fmt.Errorf("unsupported type: %s", s)
	}

	if err != nil {
		return err
	}

	return nil
}
